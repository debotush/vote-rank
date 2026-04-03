package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/debotush/vote-rank/internal/display"
	"github.com/debotush/vote-rank/internal/parser"
	"github.com/debotush/vote-rank/internal/ranking"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [input-file]",
	Short: "Run the voting ranking algorithm on an input file",
	Args:  cobra.ExactArgs(1),
	RunE:  runRanking,
}

func init() {
	rootCmd.AddCommand(runCmd)
}

func runRanking(_ *cobra.Command, args []string) error {
	filePath := args[0]
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("input file path: %s\n", filePath)
	data, err := parser.ParseFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to parse input file: %w", err)
	}
	fmt.Printf("Loaded %d candidates across %d teams, %d open positions\n\n",
		len(data.Candidates), data.Config.TeamCount, data.Config.OpenPositionCount)

	pressEnter(reader, "Press ENTER to start Phase 1...")
	phase1Result := ranking.Phase1(data.Candidates)
	fmt.Println(display.FormatPhase1State(
		phase1Result.Qualified,
		phase1Result.BelowTheLine,
		phase1Result.Draws,
	))

	pressEnter(reader, "Press ENTER to continue to Phase 2...")

	fmt.Println("\n=== PHASE 2: Resolving Draws ===")
	fmt.Println("----------------------------------------")

	var resolvedQualified []ranking.Phase1Result
	_ = resolvedQualified

	resolution := map[string][]string{}

	if len(phase1Result.Draws) == 0 {
		fmt.Println("No draws detected. Skipping Phase 2.")
	} else {
		for _, group := range phase1Result.Draws {
			fmt.Printf("\nDraw at %d votes between:\n", group[0].VoteCount)
			for i, c := range group {
				fmt.Printf("  %d. %s [%s]\n", i+1, c.Name, c.TeamName)
			}
			fmt.Println("\nEnter the names in draw-resolved order (one per line, empty line to finish):")

			var resolved []string
			for {
				fmt.Print("  > ")
				line, _ := reader.ReadString('\n')
				line = strings.TrimSpace(line)
				if line == "" {
					break
				}
				resolved = append(resolved, line)
			}

			if len(resolved) == len(group) {
				key := strconv.Itoa(group[0].VoteCount)
				resolution[key] = resolved
				fmt.Println("Draw resolved.")
			} else {
				fmt.Println("Mismatch in count — keeping original order for this group.")
			}
		}
	}

	finalQualified := ranking.Phase2WithResolution(
		phase1Result.Qualified,
		phase1Result.Draws,
		resolution,
	)

	fmt.Println("\nFinal order after Phase 2:")
	for i, c := range finalQualified {
		fmt.Printf("  %d. %-20s [%s] - %d votes\n", i+1, c.Name, c.TeamName, c.VoteCount)
	}

	pressEnter(reader, "\nPress ENTER to continue to Phase 3...")
	fmt.Println("\n=== PHASE 3: Seat Allocation ===")
	fmt.Println("----------------------------------------")

	phase3Result := ranking.Phase3(data.Config, finalQualified)
	fmt.Println(display.FormatElected(phase3Result))
	fmt.Println(display.FormatAlternates(phase3Result))

	pressEnter(reader, "Press ENTER to continue to Phase 4...")
	fmt.Println(display.FormatFullResults(phase3Result, phase1Result.BelowTheLine))
	fmt.Println("\nRanking complete!")

	return nil
}

func pressEnter(reader *bufio.Reader, prompt string) {
	fmt.Print(prompt)
	_, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("failed to instruction: %w", err)
	}
}

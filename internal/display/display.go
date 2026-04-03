package display

import (
	"fmt"
	"strings"

	"github.com/debotush/vote-rank/internal/model"
	"github.com/debotush/vote-rank/internal/ranking"
)

const separator = "----------------------------------------"

// FormatPhase1State renders the Phase 1 state for the user.
func FormatPhase1State(qualified []model.Candidate,
	belowTheLine []model.Candidate,
	draws [][]model.Candidate) string {

	var sb strings.Builder

	sb.WriteString("\n=== PHASE 1: Qualifying Candidates ===\n")
	sb.WriteString(separator + "\n")

	inDraw := make(map[string]bool)
	for _, group := range draws {
		for _, c := range group {
			inDraw[c.Name] = true
		}
	}

	sb.WriteString("Qualified Candidates (ranked by votes):\n")
	for i, c := range qualified {
		drawTag := ""
		if inDraw[c.Name] {
			drawTag = " [DRAW]"
		}
		sb.WriteString(fmt.Sprintf("  %d. %-20s [%s] - %d votes%s\n",
			i+1, c.Name, c.TeamName, c.VoteCount, drawTag))
	}

	if len(draws) > 0 {
		sb.WriteString("\nDraws detected (must be resolved in Phase 2):\n")
		for _, group := range draws {
			sb.WriteString(fmt.Sprintf("  Tied at %d votes: ", group[0].VoteCount))
			names := make([]string, len(group))
			for i, c := range group {
				names[i] = c.Name
			}
			sb.WriteString(strings.Join(names, ", ") + "\n")
		}
	}

	sb.WriteString("\n" + separator + "\n")
	sb.WriteString("Below the Line (no votes, alphabetical):\n")
	for _, c := range belowTheLine {
		sb.WriteString(fmt.Sprintf("  - %-20s [%s]\n", c.Name, c.TeamName))
	}

	return sb.String()
}

// FormatElected renders the elected members list.
func FormatElected(result ranking.Phase3Result) string {
	var sb strings.Builder

	sb.WriteString("\n=== ELECTED MEMBERS ===\n")
	sb.WriteString(separator + "\n")

	pos := 1
	sb.WriteString("Team Representatives:\n")
	for _, c := range result.TeamElected {
		sb.WriteString(fmt.Sprintf("  %d. %-20s [%s] - %d votes\n",
			pos, c.Name, c.TeamName, c.VoteCount))
		pos++
	}

	if len(result.VacancyElected) > 0 {
		sb.WriteString("\nVacancy Seats:\n")
		for _, c := range result.VacancyElected {
			sb.WriteString(fmt.Sprintf("  %d. %-20s [%s] - %d votes\n",
				pos, c.Name, c.TeamName, c.VoteCount))
			pos++
		}
	}

	return sb.String()
}

// FormatAlternates renders the alternates/replacement list.
func FormatAlternates(result ranking.Phase3Result) string {
	var sb strings.Builder

	sb.WriteString("\n=== ALTERNATES (Replacement List) ===\n")
	sb.WriteString(separator + "\n")

	if len(result.ReplacementList) == 0 {
		sb.WriteString("  (none)\n")
		return sb.String()
	}

	for i, c := range result.ReplacementList {
		sb.WriteString(fmt.Sprintf("  %d. %-20s [%s] - %d votes\n",
			i+1, c.Name, c.TeamName, c.VoteCount))
	}

	return sb.String()
}

// FormatFullResults renders the complete voting results including below-the-line.
func FormatFullResults(result ranking.Phase3Result, belowTheLine []model.Candidate) string {
	var sb strings.Builder

	sb.WriteString("\n=== FULL VOTING RESULTS ===\n")
	sb.WriteString(separator + "\n")

	// Collect all candidates in order: elected -> alternates -> below the line
	allCandidates := []model.Candidate{}
	allCandidates = append(allCandidates, result.TeamElected...)
	allCandidates = append(allCandidates, result.VacancyElected...)
	allCandidates = append(allCandidates, result.ReplacementList...)
	allCandidates = append(allCandidates, belowTheLine...)

	for i, c := range allCandidates {
		sb.WriteString(fmt.Sprintf("  %d. %-20s [%s] - %d votes\n",
			i+1, c.Name, c.TeamName, c.VoteCount))
	}

	return sb.String()
}

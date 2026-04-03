package parser

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"

    "github.com/debotush/vote-rank/internal/model"
)

// ParseFile reads the input file and returns ElectionData.
func ParseFile(path string) (*model.ElectionData, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, fmt.Errorf("could not open file: %w", err)
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)

    // Parse header line
    if !scanner.Scan() {
        return nil, fmt.Errorf("file is empty or missing header")
    }
    header := scanner.Text()
    config, err := parseHeader(header)
    if err != nil {
        return nil, fmt.Errorf("invalid header: %w", err)
    }

    // Parse candidate lines
    var candidates []model.Candidate
    lineNum := 1
    for scanner.Scan() {
        lineNum++
        line := scanner.Text()
        if strings.TrimSpace(line) == "" {
            continue
        }
        candidate, err := parseCandidate(line)
        if err != nil {
            return nil, fmt.Errorf("invalid candidate on line %d: %w", lineNum, err)
        }
        candidates = append(candidates, candidate)
    }

    return &model.ElectionData{
        Config:     config,
        Candidates: candidates,
    }, nil
}

func parseHeader(line string) (model.ElectionConfig, error) {
    parts := strings.Split(line, "\t")
    if len(parts) != 2 {
        return model.ElectionConfig{}, fmt.Errorf("expected 2 tab-separated fields, got %d", len(parts))
    }
    teamCount, err := strconv.Atoi(strings.TrimSpace(parts[0]))
    if err != nil {
        return model.ElectionConfig{}, fmt.Errorf("invalid team count: %w", err)
    }
    openPositions, err := strconv.Atoi(strings.TrimSpace(parts[1]))
    if err != nil {
        return model.ElectionConfig{}, fmt.Errorf("invalid open position count: %w", err)
    }
    return model.ElectionConfig{
        TeamCount:         teamCount,
        OpenPositionCount: openPositions,
    }, nil
}

func parseCandidate(line string) (model.Candidate, error) {
    parts := strings.Split(line, "\t")
    if len(parts) != 3 {
        return model.Candidate{}, fmt.Errorf("expected 3 tab-separated fields, got %d", len(parts))
    }
    votes, err := strconv.Atoi(strings.TrimSpace(parts[2]))
    if err != nil {
        return model.Candidate{}, fmt.Errorf("invalid vote count: %w", err)
    }
    return model.Candidate{
        Name:      strings.TrimSpace(parts[0]),
        TeamName:  strings.TrimSpace(parts[1]),
        VoteCount: votes,
    }, nil
}

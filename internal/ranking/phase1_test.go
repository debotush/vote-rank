package ranking_test

import (
    "testing"

    "github.com/debotush/vote-rank/internal/model"
    "github.com/debotush/vote-rank/internal/ranking"
)

func TestPhase1_QualifyingCandidates(t *testing.T) {
    candidates := []model.Candidate{
        {Name: "Charlie", TeamName: "Team A", VoteCount: 5},
        {Name: "Alice",   TeamName: "Team A", VoteCount: 10},
        {Name: "Diana",   TeamName: "Team C", VoteCount: 0},
        {Name: "Eve",     TeamName: "Team B", VoteCount: 3},
        {Name: "Bob",     TeamName: "Team B", VoteCount: 10},
    }

    result := ranking.Phase1(candidates)

    // First two should be Alice and Bob (both 10 votes, draw)
    if result.Qualified[0].VoteCount != 10 {
        t.Errorf("expected first qualified to have 10 votes, got %d", result.Qualified[0].VoteCount)
    }
    if result.Qualified[1].VoteCount != 10 {
        t.Errorf("expected second qualified to have 10 votes, got %d", result.Qualified[1].VoteCount)
    }

    // Third should be Charlie (5 votes)
    if result.Qualified[2].Name != "Charlie" {
        t.Errorf("expected third qualified to be Charlie, got %s", result.Qualified[2].Name)
    }

    // Fourth should be Eve (3 votes)
    if result.Qualified[3].Name != "Eve" {
        t.Errorf("expected fourth qualified to be Eve, got %s", result.Qualified[3].Name)
    }

    // Below the line: only Diana (0 votes), alphabetical
    if len(result.BelowTheLine) != 1 {
        t.Errorf("expected 1 below-the-line candidate, got %d", len(result.BelowTheLine))
    }
    if result.BelowTheLine[0].Name != "Diana" {
        t.Errorf("expected below-the-line candidate to be Diana, got %s", result.BelowTheLine[0].Name)
    }
}

func TestPhase1_AllZeroVotes(t *testing.T) {
    candidates := []model.Candidate{
        {Name: "Zara", TeamName: "Team A", VoteCount: 0},
        {Name: "Anna", TeamName: "Team B", VoteCount: 0},
    }

    result := ranking.Phase1(candidates)

    if len(result.Qualified) != 0 {
        t.Errorf("expected 0 qualified, got %d", len(result.Qualified))
    }
    if len(result.BelowTheLine) != 2 {
        t.Errorf("expected 2 below-the-line, got %d", len(result.BelowTheLine))
    }
    // Must be alphabetical
    if result.BelowTheLine[0].Name != "Anna" {
        t.Errorf("expected first below-the-line to be Anna, got %s", result.BelowTheLine[0].Name)
    }
}

func TestPhase1_DrawsDetected(t *testing.T) {
    candidates := []model.Candidate{
        {Name: "Alice", TeamName: "Team A", VoteCount: 10},
        {Name: "Bob",   TeamName: "Team B", VoteCount: 10},
        {Name: "Eve",   TeamName: "Team B", VoteCount: 3},
    }

    result := ranking.Phase1(candidates)

    if len(result.Draws) != 1 {
        t.Errorf("expected 1 draw group, got %d", len(result.Draws))
    }
    if len(result.Draws[0]) != 2 {
        t.Errorf("expected draw group to have 2 candidates, got %d", len(result.Draws[0]))
    }
}

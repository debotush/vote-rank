package ranking_test

import (
    "testing"

    "github.com/debotush/vote-rank/internal/model"
    "github.com/debotush/vote-rank/internal/ranking"
)

func TestPhase2_NoDraws(t *testing.T) {
    qualified := []model.Candidate{
        {Name: "Alice", TeamName: "Team A", VoteCount: 10},
        {Name: "Bob",   TeamName: "Team B", VoteCount: 5},
    }

    // No draws, order should stay the same
    result := ranking.Phase2(qualified, [][]model.Candidate{})

    if len(result) != 2 {
        t.Fatalf("expected 2 candidates, got %d", len(result))
    }
    if result[0].Name != "Alice" {
        t.Errorf("expected Alice first, got %s", result[0].Name)
    }
}

func TestPhase2_ResolveSingleDraw(t *testing.T) {
    qualified := []model.Candidate{
        {Name: "Alice", TeamName: "Team A", VoteCount: 10},
        {Name: "Bob",   TeamName: "Team B", VoteCount: 10},
        {Name: "Eve",   TeamName: "Team B", VoteCount: 3},
    }

    draws := [][]model.Candidate{
        {
            {Name: "Alice", TeamName: "Team A", VoteCount: 10},
            {Name: "Bob",   TeamName: "Team B", VoteCount: 10},
        },
    }

    // User resolves draw: Bob wins the lottery, so Bob goes before Alice
    resolution := map[string][]string{
        "10": {"Bob", "Alice"},
    }

    result := ranking.Phase2WithResolution(qualified, draws, resolution)

    if result[0].Name != "Bob" {
        t.Errorf("expected Bob first after draw resolution, got %s", result[0].Name)
    }
    if result[1].Name != "Alice" {
        t.Errorf("expected Alice second after draw resolution, got %s", result[1].Name)
    }
    if result[2].Name != "Eve" {
        t.Errorf("expected Eve third, got %s", result[2].Name)
    }
}

func TestPhase2_MultipleDrawGroups(t *testing.T) {
    qualified := []model.Candidate{
        {Name: "Alice",   TeamName: "Team A", VoteCount: 10},
        {Name: "Bob",     TeamName: "Team B", VoteCount: 10},
        {Name: "Charlie", TeamName: "Team A", VoteCount: 5},
        {Name: "Diana",   TeamName: "Team C", VoteCount: 5},
    }

    draws := [][]model.Candidate{
        {
            {Name: "Alice", TeamName: "Team A", VoteCount: 10},
            {Name: "Bob",   TeamName: "Team B", VoteCount: 10},
        },
        {
            {Name: "Charlie", TeamName: "Team A", VoteCount: 5},
            {Name: "Diana",   TeamName: "Team C", VoteCount: 5},
        },
    }

    resolution := map[string][]string{
        "10": {"Bob", "Alice"},
        "5":  {"Diana", "Charlie"},
    }

    result := ranking.Phase2WithResolution(qualified, draws, resolution)

    if result[0].Name != "Bob" {
        t.Errorf("expected Bob first, got %s", result[0].Name)
    }
    if result[1].Name != "Alice" {
        t.Errorf("expected Alice second, got %s", result[1].Name)
    }
    if result[2].Name != "Diana" {
        t.Errorf("expected Diana third, got %s", result[2].Name)
    }
    if result[3].Name != "Charlie" {
        t.Errorf("expected Charlie fourth, got %s", result[3].Name)
    }
}

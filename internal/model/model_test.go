package model_test

import (
	"testing"

	"github.com/debotush/vote-rank/internal/model"
)

func TestCandidate(t *testing.T) {
	c := model.Candidate{
		Name:      "Alice",
		TeamName:  "Team A",
		VoteCount: 5,
	}

	if c.Name != "Alice" {
		t.Errorf("expected Name 'Alice', got '%s'", c.Name)
	}
	if c.TeamName != "Team A" {
		t.Errorf("expected TeamName 'Team A', got '%s'", c.TeamName)
	}
	if c.VoteCount != 5 {
		t.Errorf("expected VoteCount 5, got %d", c.VoteCount)
	}
}

func TestElectionConfig(t *testing.T) {
	cfg := model.ElectionConfig{
		TeamCount:         3,
		OpenPositionCount: 5,
	}

	if cfg.TeamCount != 3 {
		t.Errorf("expected TeamCount 3, got %d", cfg.TeamCount)
	}
	if cfg.OpenPositionCount != 5 {
		t.Errorf("expected OpenPositionCount 5, got %d", cfg.OpenPositionCount)
	}
}

func TestElectionData(t *testing.T) {
	data := model.ElectionData{
		Config: model.ElectionConfig{
			TeamCount:         2,
			OpenPositionCount: 3,
		},
		Candidates: []model.Candidate{
			{Name: "Alice", TeamName: "Team A", VoteCount: 10},
			{Name: "Bob", TeamName: "Team B", VoteCount: 0},
		},
	}

	if len(data.Candidates) != 2 {
		t.Errorf("expected 2 candidates, got %d", len(data.Candidates))
	}
	if data.Candidates[0].Name != "Alice" {
		t.Errorf("expected first candidate 'Alice', got '%s'", data.Candidates[0].Name)
	}
}

package ranking_test

import (
	"testing"

	"github.com/debotush/vote-rank/internal/model"
	"github.com/debotush/vote-rank/internal/ranking"
)

func TestPhase3_BasicAllocation(t *testing.T) {
	// 2 team seats, 1 extra vacancy = 3 total open positions
	config := model.ElectionConfig{
		TeamCount:         2,
		OpenPositionCount: 3,
	}

	// Draw already resolved — order is final
	qualified := []model.Candidate{
		{Name: "Alice", TeamName: "Team A", VoteCount: 10},
		{Name: "Bob", TeamName: "Team B", VoteCount: 8},
		{Name: "Charlie", TeamName: "Team A", VoteCount: 6},
		{Name: "Diana", TeamName: "Team B", VoteCount: 4},
	}

	result := ranking.Phase3(config, qualified)

	// Alice and Bob should be team reps (one per team)
	if len(result.TeamElected) != 2 {
		t.Fatalf("expected 2 team elected, got %d", len(result.TeamElected))
	}
	if result.TeamElected[0].Name != "Alice" {
		t.Errorf("expected Alice as first team rep, got %s", result.TeamElected[0].Name)
	}
	if result.TeamElected[1].Name != "Bob" {
		t.Errorf("expected Bob as second team rep, got %s", result.TeamElected[1].Name)
	}

	// 1 vacancy left (3 positions - 2 team seats)
	// Charlie is next with most votes and not yet elected
	if len(result.VacancyElected) != 1 {
		t.Fatalf("expected 1 vacancy elected, got %d", len(result.VacancyElected))
	}
	if result.VacancyElected[0].Name != "Charlie" {
		t.Errorf("expected Charlie in vacancy, got %s", result.VacancyElected[0].Name)
	}

	// Diana is on replacement list
	if len(result.ReplacementList) != 1 {
		t.Fatalf("expected 1 on replacement list, got %d", len(result.ReplacementList))
	}
	if result.ReplacementList[0].Name != "Diana" {
		t.Errorf("expected Diana on replacement list, got %s", result.ReplacementList[0].Name)
	}
}

func TestPhase3_SkipsAlreadyRepresentedTeam(t *testing.T) {
	config := model.ElectionConfig{
		TeamCount:         1,
		OpenPositionCount: 2,
	}

	qualified := []model.Candidate{
		{Name: "Alice", TeamName: "Team A", VoteCount: 10},
		{Name: "Bob", TeamName: "Team A", VoteCount: 8},
		{Name: "Charlie", TeamName: "Team A", VoteCount: 6},
	}

	result := ranking.Phase3(config, qualified)

	if len(result.TeamElected) != 1 {
		t.Fatalf("expected 1 team elected, got %d", len(result.TeamElected))
	}
	if result.TeamElected[0].Name != "Alice" {
		t.Errorf("expected Alice as team rep, got %s", result.TeamElected[0].Name)
	}

	if len(result.VacancyElected) != 1 {
		t.Fatalf("expected 1 vacancy elected, got %d", len(result.VacancyElected))
	}
	if result.VacancyElected[0].Name != "Bob" {
		t.Errorf("expected Bob in vacancy, got %s", result.VacancyElected[0].Name)
	}

	if len(result.ReplacementList) != 1 {
		t.Fatalf("expected 1 replacement, got %d", len(result.ReplacementList))
	}
	if result.ReplacementList[0].Name != "Charlie" {
		t.Errorf("expected Charlie on replacement list, got %s", result.ReplacementList[0].Name)
	}
}

func TestPhase3_MoreSeatsThanCandidates(t *testing.T) {
	config := model.ElectionConfig{
		TeamCount:         3,
		OpenPositionCount: 10,
	}

	qualified := []model.Candidate{
		{Name: "Alice", TeamName: "Team A", VoteCount: 5},
		{Name: "Bob", TeamName: "Team B", VoteCount: 3},
	}

	result := ranking.Phase3(config, qualified)

	// Both get elected (team reps)
	if len(result.TeamElected) != 2 {
		t.Fatalf("expected 2 team elected, got %d", len(result.TeamElected))
	}

	// No vacancy candidates left
	if len(result.VacancyElected) != 0 {
		t.Errorf("expected 0 vacancy elected, got %d", len(result.VacancyElected))
	}

	// No replacement list
	if len(result.ReplacementList) != 0 {
		t.Errorf("expected 0 replacements, got %d", len(result.ReplacementList))
	}
}

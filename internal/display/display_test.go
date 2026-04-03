package display_test

import (
    "strings"
    "testing"

    "github.com/debotush/vote-rank/internal/display"
    "github.com/debotush/vote-rank/internal/model"
    "github.com/debotush/vote-rank/internal/ranking"
)

func makePhase3Result() ranking.Phase3Result {
    return ranking.Phase3Result{
        TeamElected: []model.Candidate{
            {Name: "Alice", TeamName: "Team A", VoteCount: 10},
            {Name: "Bob",   TeamName: "Team B", VoteCount: 8},
        },
        VacancyElected: []model.Candidate{
            {Name: "Charlie", TeamName: "Team A", VoteCount: 6},
        },
        ReplacementList: []model.Candidate{
            {Name: "Diana", TeamName: "Team B", VoteCount: 4},
        },
    }
}

func makeBelowTheLine() []model.Candidate {
    return []model.Candidate{
        {Name: "Eve", TeamName: "Team C", VoteCount: 0},
    }
}

func TestFormatElected(t *testing.T) {
    result := makePhase3Result()
    output := display.FormatElected(result)

    if !strings.Contains(output, "ELECTED MEMBERS") {
        t.Error("expected 'ELECTED MEMBERS' header")
    }
    if !strings.Contains(output, "Alice") {
        t.Error("expected Alice in elected output")
    }
    if !strings.Contains(output, "Bob") {
        t.Error("expected Bob in elected output")
    }
    if !strings.Contains(output, "Charlie") {
        t.Error("expected Charlie in elected output")
    }
    if strings.Contains(output, "Diana") {
        t.Error("did not expect Diana in elected output")
    }
}

func TestFormatAlternates(t *testing.T) {
    result := makePhase3Result()
    output := display.FormatAlternates(result)

    if !strings.Contains(output, "ALTERNATES") {
        t.Error("expected 'ALTERNATES' header")
    }
    if !strings.Contains(output, "Diana") {
        t.Error("expected Diana in alternates output")
    }
    if strings.Contains(output, "Alice") {
        t.Error("did not expect Alice in alternates output")
    }
}

func TestFormatFullResults(t *testing.T) {
    result := makePhase3Result()
    belowTheLine := makeBelowTheLine()
    output := display.FormatFullResults(result, belowTheLine)

    if !strings.Contains(output, "FULL VOTING RESULTS") {
        t.Error("expected 'FULL VOTING RESULTS' header")
    }
    if !strings.Contains(output, "Alice") {
        t.Error("expected Alice in full results")
    }
    if !strings.Contains(output, "Eve") {
        t.Error("expected Eve in full results")
    }
    if !strings.Contains(output, "10") {
        t.Error("expected vote count 10 in full results")
    }
    if !strings.Contains(output, "0") {
        t.Error("expected vote count 0 in full results")
    }
}

func TestFormatPhase1State(t *testing.T) {
    qualified := []model.Candidate{
        {Name: "Alice", TeamName: "Team A", VoteCount: 10},
        {Name: "Bob",   TeamName: "Team B", VoteCount: 10},
        {Name: "Eve",   TeamName: "Team B", VoteCount: 3},
    }
    belowTheLine := []model.Candidate{
        {Name: "Diana", TeamName: "Team C", VoteCount: 0},
    }
    draws := [][]model.Candidate{
        {
            {Name: "Alice", TeamName: "Team A", VoteCount: 10},
            {Name: "Bob",   TeamName: "Team B", VoteCount: 10},
        },
    }

    output := display.FormatPhase1State(qualified, belowTheLine, draws)

    if !strings.Contains(output, "PHASE 1") {
        t.Error("expected 'PHASE 1' header")
    }
    if !strings.Contains(output, "DRAW") {
        t.Error("expected 'DRAW' indicator for tied candidates")
    }
    if !strings.Contains(output, "Diana") {
        t.Error("expected Diana in below-the-line section")
    }
}

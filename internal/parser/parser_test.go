package parser_test

import (
    "testing"

    "github.com/debotush/vote-rank/internal/parser"
)

func TestParseFile_ValidInput(t *testing.T) {
    data, err := parser.ParseFile("../../testdata/sample.txt")
    if err != nil {
        t.Fatalf("expected no error, got: %v", err)
    }

    if data.Config.TeamCount != 3 {
        t.Errorf("expected TeamCount 3, got %d", data.Config.TeamCount)
    }
    if data.Config.OpenPositionCount != 2 {
        t.Errorf("expected OpenPositionCount 2, got %d", data.Config.OpenPositionCount)
    }
    if len(data.Candidates) != 5 {
        t.Errorf("expected 5 candidates, got %d", len(data.Candidates))
    }

    alice := data.Candidates[0]
    if alice.Name != "Alice" {
        t.Errorf("expected first candidate 'Alice', got '%s'", alice.Name)
    }
    if alice.TeamName != "Team A" {
        t.Errorf("expected team 'Team A', got '%s'", alice.TeamName)
    }
    if alice.VoteCount != 10 {
        t.Errorf("expected votes 10, got %d", alice.VoteCount)
    }

    diana := data.Candidates[3]
    if diana.VoteCount != 0 {
        t.Errorf("expected Diana to have 0 votes, got %d", diana.VoteCount)
    }
}

func TestParseFile_MissingFile(t *testing.T) {
    _, err := parser.ParseFile("../../testdata/nonexistent.txt")
    if err == nil {
        t.Error("expected error for missing file, got nil")
    }
}

func TestParseFile_InvalidHeader(t *testing.T) {
    _, err := parser.ParseFile("../../testdata/invalid_header.txt")
    if err == nil {
        t.Error("expected error for invalid header, got nil")
    }
}

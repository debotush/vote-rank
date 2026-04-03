package model

// Candidate represents a single candidate in the election.
type Candidate struct {
    Name      string
    TeamName  string
    VoteCount int
}

// ElectionConfig holds the top-level election parameters from the input file header.
type ElectionConfig struct {
    TeamCount         int
    OpenPositionCount int
}

// ElectionData is the full parsed input — config + all candidates.
type ElectionData struct {
    Config     ElectionConfig
    Candidates []Candidate
}

package ranking

import (
	"sort"

	"github.com/debotush/vote-rank/internal/model"
)

// Phase1Result holds the output of Phase 1.
type Phase1Result struct {
	Qualified    []model.Candidate
	BelowTheLine []model.Candidate
	Draws        [][]model.Candidate
}

// Phase1 separates and sorts candidates according to Phase 1 rules.
func Phase1(candidates []model.Candidate) Phase1Result {
	var qualified []model.Candidate
	var belowTheLine []model.Candidate

	for _, c := range candidates {
		if c.VoteCount > 0 {
			qualified = append(qualified, c)
		} else {
			belowTheLine = append(belowTheLine, c)
		}
	}

	// Sort qualified by votes descending
	sort.SliceStable(qualified, func(i, j int) bool {
		return qualified[i].VoteCount > qualified[j].VoteCount
	})

	// Sort below-the-line alphabetically
	sort.SliceStable(belowTheLine, func(i, j int) bool {
		return belowTheLine[i].Name < belowTheLine[j].Name
	})

	draws := detectDraws(qualified)

	return Phase1Result{
		Qualified:    qualified,
		BelowTheLine: belowTheLine,
		Draws:        draws,
	}
}

// detectDraws finds groups of candidates with the same vote count (size > 1).
func detectDraws(qualified []model.Candidate) [][]model.Candidate {
	var draws [][]model.Candidate

	i := 0
	for i < len(qualified) {
		j := i + 1
		for j < len(qualified) && qualified[j].VoteCount == qualified[i].VoteCount {
			j++
		}
		if j-i > 1 {
			draws = append(draws, qualified[i:j])
		}
		i = j
	}

	return draws
}

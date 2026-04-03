package ranking

import (
	"fmt"
	"strings"

	"github.com/debotush/vote-rank/internal/model"
)

// Phase2 returns the qualified list unchanged when there are no draws.
func Phase2(qualified []model.Candidate, draws [][]model.Candidate) []model.Candidate {
	if len(draws) == 0 {
		return qualified
	}
	return qualified
}

// Phase2WithResolution resolves draws by reordering tied candidates
func Phase2WithResolution(
	qualified []model.Candidate,
	draws [][]model.Candidate,
	resolution map[string][]string) []model.Candidate {

	resolvedOrder := make(map[int][]string)
	for voteStr, names := range resolution {
		var votes int
		fmt.Sscanf(voteStr, "%d", &votes)
		resolvedOrder[votes] = names
	}

	byName := make(map[string]model.Candidate)
	for _, c := range qualified {
		byName[strings.ToLower(c.Name)] = c
	}

	result := make([]model.Candidate, 0, len(qualified))
	visited := make(map[string]bool)

	i := 0
	for i < len(qualified) {
		c := qualified[i]
		if order, hasResolution := resolvedOrder[c.VoteCount]; hasResolution && !visited[c.Name] {

			for _, name := range order {
				if candidate, ok := byName[strings.ToLower(name)]; ok {
					result = append(result, candidate)
					visited[name] = true
				}
			}

			for i < len(qualified) && qualified[i].VoteCount == c.VoteCount {
				i++
			}
		} else if !visited[strings.ToLower(c.Name)] {
			result = append(result, c)
			visited[c.Name] = true
			i++
		} else {
			i++
		}
	}

	return result
}

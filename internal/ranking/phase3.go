package ranking

import (
	"github.com/debotush/vote-rank/internal/model"
)

// Phase3Result holds the outcome of Phase 3 seat allocation.
type Phase3Result struct {
	TeamElected     []model.Candidate
	VacancyElected  []model.Candidate
	ReplacementList []model.Candidate
}

// Phase3 allocates team seats and vacancies according to the ranking rules.
func Phase3(config model.ElectionConfig, qualified []model.Candidate) Phase3Result {
	teamSeats := config.TeamCount
	totalSeats := config.OpenPositionCount

	elected := make(map[string]bool)
	teamRepresented := make(map[string]bool)

	var teamElected []model.Candidate
	var vacancyElected []model.Candidate

	//Step 1: Allocate team seats
	for _, c := range qualified {
		if teamSeats <= 0 {
			break
		}
		if !teamRepresented[c.TeamName] {
			teamElected = append(teamElected, c)
			elected[c.Name] = true
			teamRepresented[c.TeamName] = true
			teamSeats--
		}
	}

	//Step 2: Calculate vacancies
	teamSeatsUsed := len(teamElected)
	vacancies := totalSeats - teamSeatsUsed

	//Step 3: Fill vacancies from remaining candidates
	for _, c := range qualified {
		if vacancies <= 0 {
			break
		}
		if !elected[c.Name] {
			vacancyElected = append(vacancyElected, c)
			elected[c.Name] = true
			vacancies--
		}
	}

	// --- Step 4: Build replacement list ---
	var replacementList []model.Candidate
	for _, c := range qualified {
		if !elected[c.Name] {
			replacementList = append(replacementList, c)
		}
	}

	return Phase3Result{
		TeamElected:     teamElected,
		VacancyElected:  vacancyElected,
		ReplacementList: replacementList,
	}
}

# Work Breakdown Structure

The goal is to build a fully functional voting rank CLI application.
Below is the breakdown of all work needed to get there.

## 1. Parse the Input File
- Read the header line to get team count and open positions
- Read each candidate line to get name, team, and vote count
- Show a clear error if the file is missing or has wrong format

## 2. Phase 1 - Sort Candidates
- Candidates with at least one vote go to the qualified list, sorted by votes (highest first)
- Candidates with zero votes go below the line, sorted alphabetically
- Detect and highlight candidates who are tied on the same vote count

## 3. Phase 2 - Resolve Draws
- Show the user which candidates are tied
- Ask the user to type the tiebreak order (draw result)
- Update the ranking based on the user's input

## 4. Phase 3 - Allocate Seats
- Go through the ranked list and give each team one representative seat
- Fill remaining seats with the next highest voted candidates
- Put everyone who was not elected onto the replacement list

## 5. Phase 4 - Show Results
- Print the list of elected members
- Print the replacement list (alternates)
- Print the full voting results with vote counts for every candidate

## 6. Testing
- Write a test for every rule in phases 1 through 4 before writing the code
- Test edge cases: all candidates tied, everyone has zero votes, more seats than candidates
- Test that the parser rejects bad input files cleanly

## 7. CI/CD Pipeline
- On every push: run go vet and staticcheck to catch code issues early
- On every push: run all tests with race detection enabled
- On every push: build the binary to make sure it compiles
- On a version tag (e.g. v1.0.0): automatically build binaries for Linux, Windows, and macOS and publish them as a release

## 8. Versioning and Branching
- All work is done on short-lived branches (e.g. feature/phase3, fix/draw-resolution)
- Branches are merged into master via pull requests after tests pass
- Master is always in a working, deployable state
- Releases are tagged using SemVer: vMAJOR.MINOR.PATCH
  - MAJOR: breaking change to input format or algorithm
  - MINOR: new feature added
  - PATCH: bug fixed

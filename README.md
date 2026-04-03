# vote-rank

A CLI application that ranks candidates based on voting results using a
4-phase algorithm. Built in Go with Cobra, developed using TDD.

## Requirements

- Go 1.22 or higher

## Installation
```bash
git clone https://github.com/debotush/vote-rank.git
cd vote-rank
go build -o vote-rank .
```

## Usage
```bash
./vote-rank run path/to/input.txt
```

On Windows:
```bash
.\vote-rank.exe run path\to\input.txt
```

## Input File Format

The input file is tab-separated with the following structure:

team-count    open-position-count
candidate-name    team-name    vote-count
candidate-name    team-name    vote-count

Example:

```
3    2
Alice    Team A    10
Bob      Team B    10
Charlie  Team A    5
Diana    Team C    0
Eve      Team B    3
```

## How It Works

The application walks through 4 phases interactively:

**Phase 1 — Qualifying candidates**
Candidates with at least one vote are ranked by votes descending.
Candidates with zero votes appear below the line in alphabetical order.
Tied candidates are highlighted for Phase 2.

**Phase 2 — Resolving draws**
If candidates are tied, the user enters the tiebreak order manually
to reflect the result of a draw (lottery).

**Phase 3 — Allocating seats**
Team seats are filled first, one representative per team.
Remaining seats are filled as vacancies from the ranked list.
Unelected candidates form the replacement list.

**Phase 4 — Results**
Displays elected members, alternates, and full voting results.

## Running Tests
```bash
go test ./...
```

With race detection:
```bash
go test -race ./...
```

## Project Structure

```
vote-rank/
├── cmd/              # CLI commands (Cobra)
├── internal/
│   ├── model/        # Core data structures
│   ├── parser/       # Input file parsing
│   ├── ranking/      # Phase 1, 2, 3 algorithm logic
│   └── display/      # Phase 4 output formatting
├── testdata/         # Sample input files for tests
├── docs/             # ADR and WBS documentation
└── .github/
    └── workflows/    # CI/CD pipeline (GitHub Actions)
```

## Documentation

- [Architectural Decision Records](docs/ADR.md)
- [Work Breakdown Structure](docs/WBS.md)

## AI Usage Disclosure

This project was built with assistance from Claude (Anthropic).

**How it was used:**
- Generating boilerplate structure for each phase (models, parser, ranking, display)
- Suggesting TDD structure — writing edge tests before implementation
- Debugging issues such as case-insensitive name matching in Phase 2
- Drafting documentation (ADR, WBS, README)

**How results were verified:**
- Every generated code block was run locally before accepting it
- All tests were executed after each step to confirm correctness
- The full application was run end-to-end with the sample input file
  after each phase was added
- Bugs found during manual runs (missing candidate after wrong-case
  input in Phase 2)

**Manual contributions:**
- All prompts and direction were written by me
- Identified the case-sensitivity bug during manual testing
- Made decisions on project structure, naming, and algorithm interpretation
- Verified the algorithm output matches the specification in the assessment
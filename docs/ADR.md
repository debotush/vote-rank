# Architectural Decision Records

## ADR-001: Language — Go

Go was chosen for its simplicity, fast compilation, strong standard
library, and trivial cross-platform binary builds via GOOS/GOARCH flags.

---

## ADR-002: CLI Framework — Cobra

Cobra is the de facto standard CLI library in the Go ecosystem, used by
Docker, Kubernetes. It handles subcommand routing, flag
parsing, and help generation out of the box. Adding new subcommands in
the future requires no changes to existing code.

---

## ADR-003: Project Structure — Internal Packages

All business logic lives under internal/ in focused packages: model,
parser, ranking, display. The cmd/ package only handles CLI wiring and
user interaction. This separation allows each package to be tested
independently without spinning up the CLI.

---

## ADR-004: Development Approach — TDD

Tests are written before implementation for all business logic. The
ranking algorithm has precise rules that are easy to get wrong silently,
for example draw resolution and team seat skipping. Writing tests first
makes each rule explicit and catches regressions immediately.

---

## ADR-005: Versioning and Branching

Branching follows GitHub Flow — master is always deployable, feature
work happens on short-lived branches merged via Pull Requests.

Versioning follows SemVer (vMAJOR.MINOR.PATCH):
- MAJOR: breaking changes to input format or algorithm behavior
- MINOR: new features such as JSON export or additional display formats
- PATCH: bug fixes such as case-insensitive name matching

**Repository Purpose**: This repo is a small Go load-generator/tooling project. The primary entry is under cmd/loadgen and core packages live under internal/. Treat this as a single-binary utility (not a long-running server).

**High-level Architecture**:
- **cmd/**: Hosts entrypoint(s). Example: cmd/loadgen/main.go is the binary entry.
- **internal/engine/**: Core engine types and logic (see internal/engine/types.go).
- **internal/stats/**: Metrics/summary helpers (see internal/stats/summary.go).
- **go.mod**: dependency root — use Go tooling for builds/tests.

**What to edit vs. what to avoid**:
- Edit: implementation inside `internal/*` to add features or change behavior. Keep exported APIs minimal and stable.
- Avoid: adding exported packages at repo root. Respect `internal/` visibility — code here is intended to be internal-only.
- Prefer small, focused changes to `cmd/loadgen` for CLI flags and wiring; keep heavy business logic in `internal/engine`.

**Developer workflows (concrete commands)**:
- Build the binary: `go build -o bin/loadgen ./cmd/loadgen`
- Run tests: `go test ./...`
- Quick lint/format: `gofmt -w .` and `go vet ./...`
- Tidy deps: `go mod tidy`

**Project-specific patterns & examples**:
- Single binary pattern: follow the `cmd/<name>/main.go` pattern for any new CLI tools.
- Core types live in `types.go` files (example: internal/engine/types.go) — add related methods in the same package.
- Stats/aggregation: keep summaries and reporting in `internal/stats` to avoid leaking metrics concerns into engine logic.

**PR/changes guidance for AI agents**:
- Small, atomic commits with a focused message: what changed and why (e.g., "engine: add request throttling" ).
- When adding surface API or flags, update `cmd/loadgen/main.go` and a single internal package only.
- Add unit tests alongside changes under the same package (use `_test.go`) and ensure `go test ./...` passes.

**Integration & external deps**:
- This project relies on standard Go tooling. Do not assume presence of third-party CI or linters unless added to go.mod or repository files.

**Notes to the next AI/editing agent**:
- Repository skeleton files may be empty initially (check file contents before modifying). Example paths to inspect first:
  - cmd/loadgen/main.go
  - internal/engine/types.go
  - internal/stats/summary.go
- If you change package APIs, update all call sites inside `internal/` and adjust tests.

If any section is unclear or you'd like conventions expanded (examples: CLI flag style, logging approach, testing thresholds), tell me which area to elaborate.

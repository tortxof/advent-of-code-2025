# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Advent of Code 2025 solutions implemented in Go. The project uses a multi-executable architecture where each day's puzzle is a separate CLI program.

## Architecture

**Multi-executable pattern**: Each day's solution is an independent program in `cmd/dayNN/main.go`. This allows:
- Building/running individual days without dependencies on others
- Each day can have its own approach and data structures
- No shared state between puzzles

**Shared utilities**: Common helper functions go in `internal/util/` for code reuse across days. The `internal/` directory prevents external imports per Go conventions.

**Input files**: Puzzle inputs are stored in `inputs/dayNN.txt` (gitignored per Advent of Code guidelines).

## Commands

Run a specific day's solution:
```bash
go run ./cmd/day01
```

Build a specific day:
```bash
go build -o day01 ./cmd/day01
./day01
```

Build all days:
```bash
for day in cmd/day*; do
    go build -o ${day#cmd/} ./$day
done
```

Format code:
```bash
go fmt ./...
```

## Import Paths

Module name is `advent-of-code-2025`, so imports from `internal/` look like:
```go
import "advent-of-code-2025/internal/util"
```

## Development Pattern

When starting a new day:
1. Create `cmd/dayNN/` directory with `main.go`
2. Add puzzle input to `inputs/dayNN.txt`
3. Each solution is self-contained; copy patterns from existing days if helpful
4. Add common utilities to `internal/util/` as patterns emerge

# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build & Test Commands

- **Run tests:** `make test` or `go test -v ./...`
- **Run a single test:** `go test -v -run TestDriverByPosition ./...`
- **Check formatting:** `make fmt`
- **Lint:** `make lint` (requires golangci-lint)
- **Build:** `go build -v ./...`
- **Install deps:** `go get -v ./...`

## Git Workflow

Never commit directly to main. Always create a feature branch for changes.

## Architecture

This is a Go library package (no main) that wraps the Formula 1 API's `/v1/event-tracker` endpoint to retrieve the latest race results. It uses `github.com/rpunt/simplehttp` as its HTTP client.

**Single-file library:** All code lives in `f1apireader.go`. The public API is:

- `RaceResults()` — fetches the latest event data from the F1 API, returns `*Event`
- `Event.Status()` — returns the race state (e.g. "completed") by finding the "Race" entry in timetables
- `Event.Winner()` — returns the 1st place `RaceResult`
- `Event.DriverByPosition(n)` — returns podium driver (positions 1-3 only)

**Data model:** The `Event` struct mirrors the F1 API JSON response. `RaceResult` contains driver/team info for podium finishers. `PositionNumber` is a string in the JSON, converted via `strconv.Atoi` for comparisons.

**Tests:** `TestF1apireader` hits the live API. `TestDriverByPosition` and `TestResultParsing` use fixture data from `test_data/verstappen_wins.json`.

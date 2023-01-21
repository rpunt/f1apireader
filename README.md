[![Go](https://github.com/rpunt/f1apireader/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/rpunt/f1apireader/actions/workflows/go.yml)

# f1apireader

A Go module to consume the latest race result from the F1 API

## Examples

### Return an `Event` struct for the most recent race

```go
# race is an Event object
race, err := f1apireader.RaceResults()
...
```

### Return the status of the target race

```go
raceStatus, err := race.Status()

...

fmt.Printf("race is %s\n", raceStatus)
// e.g. "completed"
```

### Return a RaceResult struct containing the race winner

```go
raceWinner, err := race.Winner()

...

fmt.Println("Driver TLA: %s\n", raceWinner.DriverTLA)
// e.g. "HAM"
```

### Return a RaceResult struct for the 3rd place driver (only positions 1-3 are returnable)

```go
driver, err := race.DriverByPosition(3)

...

fmt.Println("3rd place driver: %s\n", driver.DriverTLA)
// e.g. "LEC"
```

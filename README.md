# f1apireader

A Go module to consume the latest race result from the F1 API

## Examples

### Return a `Race_Status` struct for the most recent race

```go
# race is a Race_Status object
race, err := f1apireader.RaceResults("https://api.formula1.com/v1/event-tracker")
if err != nil {
  fmt.Println("err:", err)
}

# Write the TLA for the winning driver back to teh Race_Status struct

for _, event := range race.SeasonContext.Timetables {
  if event.Description == "Race" {
    switch event.State {
    case "upcoming":
      race.CurrentRaceStatus = "The race hasn't started yet"
    case "started":
      race.CurrentRaceStatus = "They're racing now"
    case "completed":
      race.CurrentRaceStatus = "They're done racing"
      for _, result := range race.RaceResults {
        position, err := strconv.Atoi(result.PositionNumber)
        if err != nil {
          fmt.Println("err:", err)
        }
        if position == 1 {
          race.Winner = result.DriverTLA
        }
      }
    }
  }
}
```

# f1apireader

A Go module to consume the latest race result from the F1 API

```go
race, err := f1apireader.RaceResults("https://api.formula1.com/v1/event-tracker")
if err != nil {
	fmt.Println("err:", err)
}

# race is a Race_Status object
```

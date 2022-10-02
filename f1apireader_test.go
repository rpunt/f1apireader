package f1apireader

import (
	"testing"
)

func TestF1apireader(t *testing.T) {
	// race_results :=
	// res, err = RaceResults("")
	// if err != nil {
	// 	fmt.Println("err:", err)
	// }
	// want := "Empty URL"

	msg, err := RaceResults("https://api.formula1.com/v1/event-tracker")
	if err != nil {
		t.Fatalf(`RaceResults("") = %q, %v, want "", error`, msg, err)
	}
}

func TestF1apireaderEmpty(t *testing.T) {
	// race_results := Race_Status("https://api.formula1.com/v1/event-tracker")
	// res, err = RaceResults("")
	// if err != nil {
	// 	fmt.Println("err:", err)
	// }
	// want := "Empty URL"

	msg, err := RaceResults("")
	if msg != nil || err == nil {
		t.Fatalf(`RaceResults("") = %q, %v, want "", error`, msg, err)
	}
}

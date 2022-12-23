package f1apireader

import (
	"encoding/json"
	"log"
	"os"
	"testing"
)

func TestF1apireader(t *testing.T) {
	race, err := RaceResults("https://api.formula1.com/v1/event-tracker")
	if err != nil {
		t.Errorf(`RaceResults("") = %q, %v, want "", error`, race, err)
	}
}

func TestF1apireaderEmpty(t *testing.T) {
	msg, err := RaceResults("")
	if msg != nil || err == nil {
		t.Fatalf(`RaceResults("") = %q, %v, want "", error`, msg, err)
	}
}

func TestResultParsing(t *testing.T) {
	race_test_data := ReadTestData("test_data/verstappen_wins.json")
	var want string
	want = "VER"
	if got, _ := race_test_data.Winner(); got != want {
		t.Errorf(`Winner() = %q, want "%v"`, got, want)
	}

	want = "completed"
	if got, _ := race_test_data.Status(); got != want {
		t.Errorf(`Status() = %q, want "%v"`, got, want)
	}
}

func ReadTestData(filename string) (Race_Status) {
	content, err := os.ReadFile(filename)
	CheckErr(err)

	race := Race_Status{}
	jsonErr := json.Unmarshal(content, &race)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return race
}

// Check for and log errors
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

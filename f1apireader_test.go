package f1apireader

import (
	"testing"
)

func TestF1apireader(t *testing.T) {
	msg, err := RaceResults("https://api.formula1.com/v1/event-tracker")
	if err != nil {
		t.Fatalf(`RaceResults("") = %q, %v, want "", error`, msg, err)
	}
}

func TestF1apireaderEmpty(t *testing.T) {
	msg, err := RaceResults("")
	if msg != nil || err == nil {
		t.Fatalf(`RaceResults("") = %q, %v, want "", error`, msg, err)
	}
}

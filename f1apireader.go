package f1apireader

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Race_Status struct {
	RaceHubID              string
	Locale                 string
	CreatedAt              string
	UpdatedAt              string
	FomRaceID              string
	BrandColourHexadecimal string
	CircuitSmallImage      struct {
		Title string
		Path  string
		Url   string
	}
	Links []struct {
		Text string
		Url  string
	}
	SeasonContext struct {
		Id                      string
		ContentType             string
		CreatedAt               time.Time
		UpdatedAt               time.Time
		Locale                  string
		SeasonYear              string
		CurrentOrNextMeetingKey string
		State                   string
		EventState              string
		LiveEventId             string
		LiveTimingsSource       string
		LiveBlog                struct {
			ContentType     string
			Title           string
			ScribbleEventId string
		}
		SeasonState                  string
		RaceListingOverride          int
		DriverAndTeamListingOverride int
		Timetables                   []struct {
			State       string
			Session     string
			GmtOffset   string
			Description string
			EndTime     string //time.Time
			StartTime   string //time.Time
		}
		ReplayBaseUrl        string
		SeasonContextUIState int
	}
	RaceResults []struct {
		DriverTLA        string
		DriverFirstName  string
		DriverLastName   string
		TeamName         string
		PositionNumber   string
		RaceTime         string
		TeamColourCode   string
		GapToLeader      string
		DriverImage      string
		DriverNameFormat string
	}
	Race struct {
		MeetingCountryName  string
		MeetingStartDate    time.Time
		MeetingOfficialName string // "FORMULA 1 PIRELLI GRAN PREMIO D’ITALIA 2022",
		MeetingEndDate      time.Time
	}
	SeasonYearImage string
	SessionLinkSets struct {
		ReplayLinks []struct {
			Session  string
			Text     string
			Url      string
			LinkType string
		}
	}
}

// Consume the F1 API for the most recent race results
func RaceResults(URL string) (*Race_Status, error) {
	if URL == "" {
		return nil, errors.New("empty url")
	}

	client := &http.Client{
		// CheckRedirect: redirectPolicyFunc,
	}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		fmt.Printf("reponse error: %s", err)
	}
	req.Header.Add("apiKey", "qPgPPRJyGCIPxFT3el4MF7thXHyJCzAP")
	req.Header.Add("locale", "en")

	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("reponse error: %s", err)
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)

	// body, err := os.ReadFile("../didhamiltonwin/api_results/20220904-132607.json")

	race_status := Race_Status{}
	jsonErr := json.Unmarshal(body, &race_status)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return &race_status, nil
}

func (r Race_Status) Status() (string, error) {
	for _, event := range r.SeasonContext.Timetables {
		if event.Description == "Race" {
			return event.State, nil
		}
	}

	return "", errors.New("unable to retrieve race timetable: no \"Race\" block")
}

func (r Race_Status) Winner() (string, error) {
	var driverTLA string
	for _, result := range r.RaceResults {
		position, err := strconv.Atoi(result.PositionNumber)
		if err != nil {
			return "", err
		}
		if position == 1 {
			driverTLA = result.DriverTLA
		}
	}
	return driverTLA, nil
}

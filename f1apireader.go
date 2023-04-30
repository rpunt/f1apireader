package f1apireader

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/rpunt/simplehttp"
)

type Event struct {
	RaceHubID              string            `json:"raceHubId,omitempty"`
	Locale                 string            `json:"locale,omitempty"`
	CreatedAt              time.Time         `json:"createdAt,omitempty"`
	UpdatedAt              time.Time         `json:"updatedAt,omitempty"`
	FomRaceID              string            `json:"fomRaceId,omitempty"`
	BrandColourHexadecimal string            `json:"brandColourHexadecimal,omitempty"`
	Headline               string            `json:"headline,omitempty"`
	CuratedSection         CuratedSection    `json:"curatedSection,omitempty"`
	CircuitSmallImage      CircuitSmallImage `json:"circuitSmallImage,omitempty"`
	Links                  []Link            `json:"links,omitempty"`
	SeasonContext          SeasonContext     `json:"seasonContext,omitempty"`
	RaceResults            []RaceResult      `json:"raceResults,omitempty"`
	Race                   Race              `json:"race,omitempty"`
	SeasonYearImage        string            `json:"seasonYearImage,omitempty"`
	SessionLinkSets        SessionLinkSets   `json:"sessionLinkSets,omitempty"`
}
type Image struct {
	Title string `json:"title,omitempty"`
	Path  string `json:"path,omitempty"`
	URL   string `json:"url,omitempty"`
}
type Thumbnail struct {
	Caption string `json:"caption,omitempty"`
	Image   Image  `json:"image,omitempty"`
}
type Items struct {
	ID              string    `json:"id,omitempty"`
	UpdatedAt       time.Time `json:"updatedAt,omitempty"`
	RealUpdatedAt   time.Time `json:"realUpdatedAt,omitempty"`
	Locale          string    `json:"locale,omitempty"`
	Title           string    `json:"title,omitempty"`
	Slug            string    `json:"slug,omitempty"`
	ArticleType     string    `json:"articleType,omitempty"`
	MetaDescription string    `json:"metaDescription,omitempty"`
	Thumbnail       Thumbnail `json:"thumbnail,omitempty"`
	MediaIcon       string    `json:"mediaIcon,omitempty"`
}
type CuratedSection struct {
	ContentType string  `json:"contentType,omitempty"`
	Title       string  `json:"title,omitempty"`
	Items       []Items `json:"items,omitempty"`
}
type CircuitSmallImage struct {
	Title string `json:"title,omitempty"`
	Path  string `json:"path,omitempty"`
	URL   string `json:"url,omitempty"`
}
type Link struct {
	Text string `json:"text,omitempty"`
	URL  string `json:"url,omitempty"`
}
type LiveBlog struct {
	ContentType     string `json:"contentType,omitempty"`
	Title           string `json:"title,omitempty"`
	ScribbleEventID string `json:"scribbleEventId,omitempty"`
}
type Timetables struct {
	State       string `json:"state,omitempty"`
	Session     string `json:"session,omitempty"`
	GmtOffset   string `json:"gmtOffset,omitempty"`
	Description string `json:"description,omitempty"`
	EndTime     string `json:"endTime,omitempty"`
	StartTime   string `json:"startTime,omitempty"`
}
type SeasonContext struct {
	ID                           string       `json:"id,omitempty"`
	ContentType                  string       `json:"contentType,omitempty"`
	CreatedAt                    time.Time    `json:"createdAt,omitempty"`
	UpdatedAt                    time.Time    `json:"updatedAt,omitempty"`
	Locale                       string       `json:"locale,omitempty"`
	SeasonYear                   string       `json:"seasonYear,omitempty"`
	CurrentOrNextMeetingKey      string       `json:"currentOrNextMeetingKey,omitempty"`
	State                        string       `json:"state,omitempty"`
	EventState                   string       `json:"eventState,omitempty"`
	LiveEventID                  string       `json:"liveEventId,omitempty"`
	LiveTimingsSource            string       `json:"liveTimingsSource,omitempty"`
	LiveBlog                     LiveBlog     `json:"liveBlog,omitempty"`
	SeasonState                  string       `json:"seasonState,omitempty"`
	RaceListingOverride          int          `json:"raceListingOverride,omitempty"`
	DriverAndTeamListingOverride int          `json:"driverAndTeamListingOverride,omitempty"`
	Timetables                   []Timetables `json:"timetables,omitempty"`
	ReplayBaseURL                string       `json:"replayBaseUrl,omitempty"`
	SeasonContextUIState         int          `json:"seasonContextUIState,omitempty"`
}
type RaceResult struct {
	DriverTLA        string `json:"driverTLA,omitempty"`
	DriverFirstName  string `json:"driverFirstName,omitempty"`
	DriverLastName   string `json:"driverLastName,omitempty"`
	TeamName         string `json:"teamName,omitempty"`
	PositionNumber   string `json:"positionNumber,omitempty"`
	RaceTime         string `json:"raceTime,omitempty"`
	TeamColourCode   string `json:"teamColourCode,omitempty"`
	GapToLeader      string `json:"gapToLeader,omitempty"`
	DriverImage      string `json:"driverImage,omitempty"`
	DriverNameFormat string `json:"driverNameFormat,omitempty"`
}
type Race struct {
	MeetingCountryName  string    `json:"meetingCountryName,omitempty"`
	MeetingStartDate    time.Time `json:"meetingStartDate,omitempty"`
	MeetingOfficialName string    `json:"meetingOfficialName,omitempty"`
	MeetingEndDate      time.Time `json:"meetingEndDate,omitempty"`
}
type ReplayLinks struct {
	Session  string `json:"session,omitempty"`
	Text     string `json:"text,omitempty"`
	URL      string `json:"url,omitempty"`
	LinkType string `json:"linkType,omitempty"`
}
type SessionLinkSets struct {
	ReplayLinks []ReplayLinks `json:"replayLinks,omitempty"`
}

// Consume the F1 API for the most recent race results
func RaceResults() (*Event, error) {
	client := simplehttp.New("https://api.formula1.com")
	client.Headers["apiKey"] = "qPgPPRJyGCIPxFT3el4MF7thXHyJCzAP"
	client.Headers["locale"] = "en"

	raceStatus := Event{}

	response, err := client.Get("/v1/event-tracker")
	if err != nil {
		return &raceStatus, err
	}

	jsonErr := json.Unmarshal([]byte(response.Body), &raceStatus)
	if jsonErr != nil {
		return &raceStatus, jsonErr
	}

	return &raceStatus, nil
}

func (r Event) Status() (string, error) {
	for _, event := range r.SeasonContext.Timetables {
		if event.Description == "Race" {
			return event.State, nil
		}
	}

	return "", errors.New("unable to retrieve race timetable: no \"Race\" block")
}

func (r Event) Winner() (RaceResult, error) {
	driver, err := r.DriverByPosition(1)
	if err != nil {
		return RaceResult{}, err
	}
	return driver, nil
}

func (r Event) DriverByPosition(desiredPosition int) (RaceResult, error) {
	var driver RaceResult
	for _, result := range r.RaceResults {
		position, err := strconv.Atoi(result.PositionNumber)
		if err != nil {
			return driver, err
		}
		if desiredPosition < 1 || desiredPosition > 3 {
			// only positions 1-3 are tracked as RaceResults
			return driver, errors.New("DriverByPosition(): valid values are 1, 2, 3")
		}
		if position == desiredPosition {
			driver = result
		}
	}
	return driver, nil
}

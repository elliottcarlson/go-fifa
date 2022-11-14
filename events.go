package go_fifa

import (
	"time"
)

type EventMatchStatus string

const (
	Started   EventMatchStatus = "EVENT_STATUS_STARTED"
	Finished  EventMatchStatus = "EVENT_STATUS_FINISHED"
	Scheduled EventMatchStatus = "EVENT_STATUS_SCHEDULED"
)

type GetLiveEventsResponse struct {
	Events            []Event `json:"events"`
	NextPageToken     string  `json:"next_page_token"`
	PreviousPageToken string  `json:"previous_page_token"`
}

type Event struct {
	Description  string           `json:"description"`
	ID           string           `json:"id"`
	IsTest       bool             `json:"is_test"`
	Location     EventLocation    `json:"location"`
	Metadata     EventMetadata    `json:"metadata"`
	Organizer    string           `json:"organizer"`
	PosterURL    string           `json:"poster_url"`
	StartTime    time.Time        `json:"start_time"`
	Status       EventMatchStatus `json:"status"`
	ThumbnailURL string           `json:"thumbnail_url"`
	Timezone     string           `json:"timezone"`
	Title        string           `json:"title"`
	Views        string           `json:"views"`
}

type EventLocation struct {
	Physical EventLocationPhysical `json:"physical"`
}

type EventLocationPhysical struct {
	City          string                   `json:"city"`
	ContinentCode string                   `json:"continent_code"`
	Coordinates   EventLocationCoordinates `json:"coordinates"`
	CountryCode   string                   `json:"country_code"`
	Venue         string                   `json:"venue"`
}

type EventLocationCoordinates struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type EventMetadata struct {
	AwayTeam         EventMetadataTeam `json:"away_team"`
	Category         string            `json:"category"`
	CompetitionID    string            `json:"competition_id"`
	CompetitionName  string            `json:"competition_name"`
	CUID             string            `json:"cuid"`
	HasChat          bool              `json:"has_chat"`
	HomeTeam         EventMetadataTeam `json:"home_team"`
	LastModifiedDate string            `json:"last_modified_date"`
	LastModifiedTime string            `json:"last_modified_time"`
	// LiveStartedAt      EventLiveTimestamp      `json:"live_started_at"`
	// LiveStoppedAt      EventLiveTimestamp      `json:"live_stopped_at"`
	MatchRecapAssetURL string                  `json:"match_recap_asset_url"`
	MCLSInternalData   []EventMetadataInternal `json:"mcls_internal_data"`
	MYCEventID         string                  `json:"myc_event_id"`
	PPVData            []interface{}           `json:"ppv_data"`
	Privacy            string                  `json:"privacy"`
	Referees           []interface{}           `json:"referees"`
	SeasonID           string                  `json:"season_id"`
	SeasonName         string                  `json:"season_name"`
	TVID               string                  `json:"tv_id"`
}

type EventMetadataTeam struct {
	Abbreviation string        `json:"abbreviation"`
	Bench        []interface{} `json:"bench"`
	Color        string        `json:"color"`
	Logo         string        `json:"logo"`
	Manager      string        `json:"manager"`
	Name         string        `json:"name"`
	Score        int           `json:"score"`
	Squad        []interface{} `json:"squad"`
	TeamID       string        `json:"team_id"`
}

type EventLiveTimestamp struct {
	Long struct {
		Timestamp time.Time `json:"timestamp-millis"`
	} `json:"long"`
}

type EventMetadataInternal struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type eventOptions struct {
	OrderBy   string    `url:"order_by"`
	PageSize  int       `url:"page_size"`
	StartTime time.Time `url:"start_after_time,omitempty"`
	EndTime   time.Time `url:"start_before_time,omitempty"`
	After     string    `json:"after,omitempty"`
	ID        string    `json:"competition"`
}

func (c *Client) GetLiveEvents() ([]Event, error) {
	st, err := time.Parse("2006-01-02T15:04:05", "2022-09-08T07:00:00")
	if err != nil {
		return nil, err
	}
	be, err := time.Parse("2006-01-02T15:04:05", "2022-09-09T07:00:00")
	if err != nil {
		return nil, err
	}
	opts := eventOptions{
		OrderBy:   "ORDER_START_TIME_ASC",
		PageSize:  8,
		ID:        "cjp46zxxt00kf0hv3m6f3x7wb",
		StartTime: st,
		EndTime:   be,
	}
	url := "/bff/events/v1beta1"
	var events []Event
	c.ApiBaseURL = "https://mcls-api.mycujoo.tv"
	for {
		var resp GetLiveEventsResponse
		_, err := c.get(url, &resp, opts)
		if err != nil {
			return nil, err
		}
		events = append(events, resp.Events...)
		if resp.NextPageToken == "" {
			break
		}
		opts.After = resp.NextPageToken
		break
	}
	c.ApiBaseURL = ""

	return events, nil
}

// func (c *Client) GetMatchEvents(options *GetMatchEventOptions) (*GetMatchEventsResponse, error) {
// 	if options.CompetitionId == "" {
// 		return nil, errors.New("competitionId is required but not provIded")
// 	}
// 	if options.SeasonId == "" {
// 		return nil, errors.New("seasonId is required but not provIded")
// 	}
// 	if options.StageId == "" {
// 		return nil, errors.New("stageId is required but not provIded")
// 	}
// 	if options.MatchId == "" {
// 		return nil, errors.New("matchId is required but not provIded")
// 	}
// 	url := fmt.Sprintf("/timelines/%s/%s/%s/%s", options.CompetitionId, options.SeasonId, options.StageId, options.MatchId)
// 	var respData GetMatchEventsResponse
// 	_, err := c.get(url, &respData, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &respData, nil
// }

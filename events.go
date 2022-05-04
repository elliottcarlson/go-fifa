package go_fifa

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type GetMatchEventOptions struct {
	CompetitionID string
	SeasonID      string
	StageID       string
	MatchID       string
}

type GetMatchEventsResponse struct {
	StageID       string          `json:"IdStage"`
	MatchID       string          `json:"IdMatch"`
	CompetitionID string          `json:"IdCompetition"`
	SeasonID      string          `json:"IdSeason"`
	GroupID       string          `json:"IdGroup"`
	Event         []EventResponse `json:"Event"`
	Properties    interface{}     `json:"Properties"`
	IsUpdateable  bool            `json:"IsUpdateable"`
}

type EventResponse struct {
	Context             string                       `json:"Context"`
	ID                  string                       `json:"IdEvent"`
	Guid                uuid.UUID                    `json:"Guid"`
	TeamID              string                       `json:"IdTeam"`
	PlayerID            string                       `json:"IdPlayer"`
	SubPlayerID         string                       `json:"IdSubPlayer"`
	PersonID            string                       `json:"IdPerson"`
	SubTeamID           string                       `json:"IdSubTeam"`
	Timestamp           time.Time                    `json:"Timestamp"`
	DateTimeUTC         time.Time                    `json:"DateTimeUTC"`
	MatchMinute         string                       `json:"MatchMinute"`
	Period              PeriodEnum                   `json:"Period"`
	HomeGoals           int                          `json:"HomeGoals"`
	AwayGoals           int                          `json:"AwayGoals"`
	Type                MatchEvent                   `json:"Type"`
	TypeLocalized       []DefaultDescriptionResponse `json:"TypeLocalized"`
	PositionX           float32                      `json:"PositionX"`
	PositionY           float32                      `json:"PositionY"`
	GoalGatePositionX   float32                      `json:"GoalGatePositionX"`
	GoalGatePositionY   float32                      `json:"GoalGatePositionY"`
	GoalGatePositionZ   float32                      `json:"GoalGatePositionZ"`
	VarDetail           string                       `json:"VarDetail"`
	VarNotificationData string                       `json:"VarNotificationData"`
	HomePenaltyGoals    int                          `json:"HomePenaltyGoals"`
	AwayPenaltyGoals    int                          `json:"AwayPenaltyGoals"`
	EventDescription    []DefaultDescriptionResponse `json:"EventDescription"`
}

func (c *client) GetMatchEvents(options *GetMatchEventOptions) (*GetMatchEventsResponse, error) {
	if options.CompetitionID == "" {
		return nil, errors.New("competitionID is required but not provided")
	}
	if options.SeasonID == "" {
		return nil, errors.New("seasonID is required but not provided")
	}
	if options.StageID == "" {
		return nil, errors.New("stageID is required but not provided")
	}
	if options.MatchID == "" {
		return nil, errors.New("matchID is required but not provided")
	}
	url := fmt.Sprintf("/timelines/%s/%s/%s/%s", options.CompetitionID, options.SeasonID, options.StageID, options.MatchID)
	var respData GetMatchEventsResponse
	_, err := c.get(url, &respData, nil)
	if err != nil {
		return nil, err
	}
	return &respData, nil
}

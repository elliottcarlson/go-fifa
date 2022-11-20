package go_fifa

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type GetMatchEventOptions struct {
	CompetitionId string
	SeasonId      string
	StageId       string
	MatchId       string
}

type GetMatchEventsResponse struct {
	StageId       string          `json:"IdStage"`
	MatchId       string          `json:"IdMatch"`
	CompetitionId string          `json:"IdCompetition"`
	SeasonId      string          `json:"IdSeason"`
	GroupId       string          `json:"IdGroup"`
	Events        []EventResponse `json:"Event"`
	Properties    interface{}     `json:"Properties"`
	IsUpdateable  bool            `json:"IsUpdateable"`
}

type EventResponse struct {
	Context             string                       `json:"Context"`
	Id                  string                       `json:"EventId"`
	GuId                uuid.UUID                    `json:"GuId"`
	TeamId              string                       `json:"IdTeam"`
	PlayerId            string                       `json:"IdPlayer"`
	SubPlayerId         string                       `json:"IdSubPlayer"`
	PersonId            string                       `json:"IdPerson"`
	SubTeamId           string                       `json:"IdSubTeam"`
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
	VarNotificationData VarNotificationDataResponse  `json:"VarNotificationData"`
	HomePenaltyGoals    int                          `json:"HomePenaltyGoals"`
	AwayPenaltyGoals    int                          `json:"AwayPenaltyGoals"`
	EventDescription    []DefaultDescriptionResponse `json:"EventDescription"`
}

func (c *Client) GetMatchEvents(options *GetMatchEventOptions) (*GetMatchEventsResponse, error) {
	if options.CompetitionId == "" {
		return nil, errors.New("competitionId is required but not provIded")
	}
	if options.SeasonId == "" {
		return nil, errors.New("seasonId is required but not provIded")
	}
	if options.StageId == "" {
		return nil, errors.New("stageId is required but not provIded")
	}
	if options.MatchId == "" {
		return nil, errors.New("matchId is required but not provIded")
	}
	url := fmt.Sprintf("/timelines/%s/%s/%s/%s", options.CompetitionId, options.SeasonId, options.StageId, options.MatchId)
	var respData GetMatchEventsResponse
	_, err := c.get(url, &respData, nil)
	if err != nil {
		return nil, err
	}
	return &respData, nil
}

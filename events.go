package go_fifa

import (
	"errors"
	"fmt"
)

// GetMatchEventOptions represents the inputs used to get data about matches
type GetMatchEventOptions struct {
	CompetitionId string
	SeasonId      string
	StageId       string
	MatchId       string
}

// GetMatchEvents returns the events from a match
func (c *Client) GetMatchEvents(options *GetMatchEventOptions) (TimelineResponse, error) {
	if options.CompetitionId == "" {
		return TimelineResponse{}, errors.New("competitionId is required but not provIded")
	}
	if options.SeasonId == "" {
		return TimelineResponse{}, errors.New("seasonId is required but not provIded")
	}
	if options.StageId == "" {
		return TimelineResponse{}, errors.New("stageId is required but not provIded")
	}
	if options.MatchId == "" {
		return TimelineResponse{}, errors.New("matchId is required but not provIded")
	}
	url := fmt.Sprintf("/timelines/%s/%s/%s/%s", options.CompetitionId, options.SeasonId, options.StageId, options.MatchId)
	var respData TimelineResponse
	_, err := c.get(url, &respData, nil)
	if err != nil {
		return TimelineResponse{}, err
	}
	return respData, nil
}

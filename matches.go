package go_fifa

import (
	"time"
)

// GetMatchesOptions represents the inputs when looking up matches
type GetMatchesOptions struct {
	From          time.Time `url:"from,omitempty"`
	To            time.Time `url:"to,omitempty"`
	Count         int       `url:"Count,omitempty"`
	TeamId        string    `url:"IdTeam,omitempty"`
	SeasonId      string    `url:"IdSeason,omitempty"`
	CompetitionId string    `url:"IdCompetition,omitempty"`
	MatchId       string    `url:"IdMatch,omitempty"`
	StageId       string    `url:"IdSeason,omitempty"`
}

// GetCurrentMatches returns the list of active matches
func (c *Client) GetCurrentMatches() ([]LiveMatch, error) {
	var respData LiveResponse
	_, err := c.get("/live/football/now", &respData, nil)
	if err != nil {
		return nil, err
	}
	return respData.Results, nil
}

// GetMatches returns the information about queried matches
func (c *Client) GetMatches(opts *GetMatchesOptions) ([]MatchData, error) {
	var respData MatchesResponse
	if opts.Count == 0 {
		opts.Count = 500
	}
	_, err := c.get("/calendar/matches", &respData, opts)
	if err != nil {
		return nil, err
	}
	return respData.Results, nil
}

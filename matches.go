package go_fifa

import (
	"time"
)

type CurrentMatchesResponse struct {
	PaginatedResponse
	Results []MatchResponse `json:"Results"`
}

type GetCurrentMatchesOptions struct {
	Competitions []string
}

type GetMatchesOptions struct {
	From time.Time `url:"from"`
	To   time.Time `url:"to"`
}

type GetTeamMatchesOptions struct {
	Count         int    `url:"Count"`
	TeamId        string `url:"IdTeam"`
	SeasonId      string `url:"IdSeason"`
	CompetitionId string `url:"IdCompetition"`
}

func (c *Client) GetCurrentMatches() ([]MatchResponse, error) {
	var respData CurrentMatchesResponse
	_, err := c.get("/live/football/now", &respData, nil)
	if err != nil {
		return nil, err
	}
	return respData.Results, nil
}

func (c *Client) GetTodaysMatches() ([]MatchResponse, error) {
	now := time.Now().UTC()
	topHour := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, time.UTC)
	options := &GetMatchesOptions{
		From: topHour,
		To:   topHour.Add(time.Hour * 24),
	}
	var respData CurrentMatchesResponse
	_, err := c.get("/calendar/matches", &respData, options)
	if err != nil {
		return nil, err
	}
	return respData.Results, nil
}

func (c *Client) GetTeamMatches(opts *GetTeamMatchesOptions) ([]MatchResponse, error) {
	var respData CurrentMatchesResponse
	if opts.Count == 0 {
		opts.Count = 500
	}
	_, err := c.get("/calendar/matches", &respData, opts)
	if err != nil {
		return nil, err
	}
	return respData.Results, nil
}

package go_fifa

import (
	"fmt"
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

type GetMatchDataOptions struct {
	CompetitionId string
	SeasonId      string
	StageId       string
	MatchId       string
}

func (c *Client) GetCurrentMatches() ([]MatchResponse, error) {
	var respData CurrentMatchesResponse
	_, err := c.get("/live/football/now", &respData, nil)
	if err != nil {
		return nil, err
	}
	return respData.Results, nil
}

func (c *Client) GetUpcomingMatches() ([]MatchResponse, error) {
	now := time.Now()
	startHour := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, now.Location()).UTC()
	options := &GetMatchesOptions{
		From: startHour,
		To:   startHour.Add(time.Hour * 24),
	}
	var respData CurrentMatchesResponse
	_, err := c.get("/calendar/matches", &respData, options)
	if err != nil {
		return nil, err
	}
	return respData.Results, nil
}

func (c *Client) GetTodaysMatches() ([]MatchResponse, error) {
	now := time.Now()
	startDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).UTC()
	options := &GetMatchesOptions{
		From: startDay,
		To:   startDay.Add(time.Hour * 24),
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

func (c *Client) GetMatchData(options *GetMatchDataOptions) (MatchDataResponse, error) {
	var respData MatchDataResponse
	url := fmt.Sprintf("/live/football/%s/%s/%s/%s", options.CompetitionId, options.SeasonId, options.StageId, options.MatchId)
	_, err := c.get(url, &respData, nil)
	if err != nil {
		return MatchDataResponse{}, err
	}
	return respData, nil
}

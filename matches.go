package go_fifa

import "time"

type CurrentMatchesResponse struct {
	PaginatedResponse
	Results []MatchResponse `json:"Results"`
}

type GetCurrentMatchesOptions struct {
	Competitions []string
}

type GetTodaysMatchesOptions struct {
	From time.Time `query:"from"`
	To   time.Time `query:"to"`
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
	options := &GetTodaysMatchesOptions{
		From: time.Now(),
		To:   time.Now().Add(time.Hour * 24),
	}
	var respData CurrentMatchesResponse
	_, err := c.get("/calendar/matches", &respData, options)
	if err != nil {
		return nil, err
	}
	return respData.Results, nil
}

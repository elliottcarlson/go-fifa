package go_fifa

import (
	"errors"
	"fmt"
)

type GetCompetitionsResponse struct {
	PaginatedResponse
	Results []CompetitionResponse `json:"Results"`
}

type GetCompetitionsOptions struct {
	CompetitionID string
}

func (c *Client) GetCompetitions() ([]CompetitionResponse, error) {
	var respData GetCompetitionsResponse
	url := "/competitions"
	_, err := c.get(url, &respData, nil)
	if err != nil {
		return nil, err
	}
	return respData.Results, nil
}

func (c *Client) GetCompetition(options *GetCompetitionsOptions) (*CompetitionResponse, error) {
	if options.CompetitionID == "" {
		return nil, errors.New("competitionID is required but was not provided")
	}
	var respData CompetitionResponse
	url := fmt.Sprintf("/competitions/%s", options.CompetitionID)
	_, err := c.get(url, &respData, nil)
	if err != nil {
		return nil, err
	}
	return &respData, nil
}

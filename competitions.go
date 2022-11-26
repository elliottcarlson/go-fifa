package go_fifa

import (
	"errors"
	"fmt"
)

// GetCompetitionsResponse represents the input used when looking up a competition
type GetCompetitionsOptions struct {
	CompetitionId string
}

// GetCompetitions returns the list of competitions
func (c *Client) GetCompetitions() ([]CompetitionResponse, error) {
	var respData GetCompetitionsResponse
	url := "/competitions"
	_, err := c.get(url, &respData, nil)
	if err != nil {
		return nil, err
	}
	return respData.Results, nil
}

// GetCompetition returns the information about a competition
func (c *Client) GetCompetition(options *GetCompetitionsOptions) (*CompetitionResponse, error) {
	if options.CompetitionId == "" {
		return nil, errors.New("competitionId is required but was not provIded")
	}
	var respData CompetitionResponse
	url := fmt.Sprintf("/competitions/%s", options.CompetitionId)
	_, err := c.get(url, &respData, nil)
	if err != nil {
		return nil, err
	}
	return &respData, nil
}

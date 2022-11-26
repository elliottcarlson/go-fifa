package go_fifa

import "fmt"

// GetSeasonOptions represents the inputs when looking up a season
type GetSeasonOptions struct {
	SeasonId string
}

// GetSeason returns information about a season
func (c *Client) GetSeason(options *GetSeasonOptions) (*SeasonResponse, error) {
	var season SeasonResponse
	url := fmt.Sprintf("/seasons/%s", options.SeasonId)
	_, err := c.get(url, &season, nil)
	if err != nil {
		return nil, err
	}
	return &season, nil
}

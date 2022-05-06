package go_fifa

import "fmt"

type GetSeasonOptions struct {
	SeasonId string
}

func (c *Client) GetSeason(options *GetSeasonOptions) (*SeasonResponse, error) {
	var season SeasonResponse
	url := fmt.Sprintf("/seasons/%s", options.SeasonId)
	_, err := c.get(url, &season, nil)
	if err != nil {
		return nil, err
	}
	return &season, nil
}

type GetSeasonStandingsOptions struct {
	SeasonId      string
	CompetitionId string
	StageId       string
}

func (c *Client) GetSeasonStandings(opts *GetSeasonStandingsOptions) ([]StandingsResult, error) {
	var response StandingResponse
	url := fmt.Sprintf("/calendar/%s/%s/%s/Standing", opts.CompetitionId, opts.SeasonId, opts.StageId)
	_, err := c.get(url, &response, nil)
	if err != nil {
		return nil, err
	}
	return response.Results, nil
}

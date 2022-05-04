package go_fifa

import "fmt"

type GetTeamOptions struct {
	TeamID string
}

func (c *client) GetTeam(options *GetTeamOptions) (*TeamResponse, error) {
	var team TeamResponse
	url := fmt.Sprintf("/teams/%s", options.TeamID)
	_, err := c.get(url, &team, nil)
	if err != nil {
		return nil, err
	}
	return &team, nil
}

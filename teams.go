package go_fifa

import "fmt"

// GetTeamOptions represents the inputs when looking up a team
type GetTeamOptions struct {
	TeamId string
}

// GetTeam returns information about the provided team
func (c *Client) GetTeam(opts *GetTeamOptions) (*TeamResponse, error) {
	var team TeamResponse
	url := fmt.Sprintf("/teams/%s", opts.TeamId)
	_, err := c.get(url, &team, nil)
	if err != nil {
		return nil, err
	}
	return &team, nil
}

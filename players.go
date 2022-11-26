package go_fifa

import "fmt"

// GetPlayerOptions represents the inputs when looking up a player
type GetPlayerOptions struct {
	PlayerId string
}

// GetPlayer returns information about a player
func (c *Client) GetPlayer(options *GetPlayerOptions) (*PlayerResponse, error) {
	var player PlayerResponse
	url := fmt.Sprintf("/players/%s", options.PlayerId)
	_, err := c.get(url, &player, nil)
	if err != nil {
		return nil, err
	}
	return &player, nil
}

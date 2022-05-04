package go_fifa

import "fmt"

type GetPlayerOptions struct {
	PlayerID string
}

func (c *client) GetPlayer(options *GetPlayerOptions) (*PlayerResponse, error) {
	var player PlayerResponse
	url := fmt.Sprintf("/players/%s", options.PlayerID)
	_, err := c.get(url, &player, nil)
	if err != nil {
		return nil, err
	}
	return &player, nil
}

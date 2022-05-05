package go_fifa

import "fmt"

type GetPlayerOptions struct {
	PlayerId string
}

func (c *Client) GetPlayer(options *GetPlayerOptions) (*PlayerResponse, error) {
	var player PlayerResponse
	url := fmt.Sprintf("/players/%s", options.PlayerId)
	_, err := c.get(url, &player, nil)
	if err != nil {
		return nil, err
	}
	return &player, nil
}

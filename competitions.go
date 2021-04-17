package go_fifa

type GetCompetitionsResponse struct {
	PaginatedResponse
	Results []CompetitionResponse `json:"Results"`
}

func (c *Client) GetCompetitions() ([]CompetitionResponse, error) {
	var respData GetCompetitionsResponse
	_, err := c.get("/competitions", &respData, nil)
	if err != nil {
		return nil, err
	}
	return respData.Results, nil
}

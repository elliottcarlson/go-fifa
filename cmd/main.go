package main

import (
	"fmt"

	fifa "github.com/ImDevinC/go-fifa"
)

const (
	competitionID string = "2000000106" // 17 world cup
)

func main() {
	client, err := fifa.NewClient(nil)
	if err != nil {
		panic(err)
	}
	matches, err := client.GetCurrentMatches()
	if err != nil {
		panic(err)
	}
	for _, m := range matches {
		fmt.Printf("CompetitionID: %s, SeasonID: %s, StageID: %s, MatchID: %s\n", m.CompetitionID, m.SeasonID, m.StageID, m.ID)
	}
	comp, err := client.GetCompetition(&fifa.GetCompetitionsOptions{CompetitionID: competitionID})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", comp)
}

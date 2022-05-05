package main

import (
	"fmt"

	fifa "github.com/ImDevinC/go-fifa"
)

const (
	competitionId string = "2000000106" // 17 world cup
)

func main() {
	// client, err := fifa.NewClient(nil)
	client := fifa.Client{}
	if err != nil {
		panic(err)
	}
	matches, err := client.GetCurrentMatches()
	if err != nil {
		panic(err)
	}
	for _, m := range matches {
		fmt.Printf("CompetitionId: %s, SeasonId: %s, StageId: %s, MatchId: %s\n", m.CompetitionId, m.SeasonId, m.StageId, m.Id)
	}
	comp, err := client.GetCompetition(&fifa.GetCompetitionsOptions{CompetitionId: competitionId})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", comp)
}

package main

import (
	"fmt"
	"log"
	"strings"

	fifa "github.com/ImDevinC/go-fifa"
)

const (
	competitionId string = "17" // 17 world cup
)

func showCompetitions(client *fifa.Client) error {
	comps, err := client.GetCompetitions()
	if err != nil {
		return err
	}
	for _, c := range comps {
		log.Printf("Id: %s", c.CompetitionId)
		log.Printf("Name: %s", c.Name[0].Description)
	}
	return nil
}

func getCompetitionsByName(client *fifa.Client, input string) error {
	comps, err := client.GetCompetitions()
	if err != nil {
		return err
	}
	input = strings.ToLower(input)
	var results []fifa.CompetitionResponse
	for _, c := range comps {
		name := strings.ToLower(c.Name[0].Description)
		if strings.Contains(name, input) {
			results = append(results, c)
		}
	}
	if len(results) > 1 {
		fmt.Println("Multiple tournaments returned")
		fmt.Printf("%+v\n", results)
	} else if len(results) == 0 {
		fmt.Println("Found " + results[0].Name[0].Description)
	} else {
		fmt.Println("No results found")
	}
	return nil
}

func getLiveMatches(client *fifa.Client) error {
	matches, err := client.GetCurrentMatches()
	if err != nil {
		return err
	}
	for _, m := range matches {
		fmt.Printf("%s: %s vs %s (%s/%s/%s/%s)\n", m.Competition[0].Description, m.HomeTeam.Name[0].Description, m.AwayTeam.Name[0].Description, m.CompetitionId, m.SeasonId, m.StageId, m.Id)
	}
	return nil
}

func getMatchData(client *fifa.Client) error {
	events, err := client.GetMatchEvents(&fifa.GetMatchEventOptions{
		CompetitionId: "2000000005",
		SeasonId:      "400250052",
		StageId:       "b1ayaoa4q68n6464fy4orklqs",
		MatchId:       "3y748w6ppuxciynnoonrt9jx0",
	})
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", events)
	return nil
}

func main() {
	client := &fifa.Client{}
	getMatchData(client)

}

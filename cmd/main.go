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
func main() {
	client := &fifa.Client{}
	// showCompetitions(client)
	getCompetitionsByName(client, "world cup")
}

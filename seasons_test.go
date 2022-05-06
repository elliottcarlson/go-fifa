package go_fifa_test

import (
	"fmt"
	"testing"

	fifa "github.com/ImDevinC/go-fifa"
	"github.com/stretchr/testify/assert"
)

func TestGetSeason(t *testing.T) {
	client := fifa.Client{}
	_, err := client.GetSeason(&fifa.GetSeasonOptions{
		SeasonId: "400198457",
	})
	if ok := assert.Nil(t, err, "expected no error with GetSeason, got: %s", err); !ok {
		t.FailNow()
	}
}

func TestGetSeasonStandings(t *testing.T) {
	client := fifa.Client{}
	resp, err := client.GetSeasonStandings(&fifa.GetSeasonStandingsOptions{
		CompetitionId: "2000001049",
		SeasonId:      "400198457",
		StageId:       "400198931",
	})
	if ok := assert.Nil(t, err, "expected no error with GetSeasonStandings, got: %s", err); !ok {
		t.FailNow()
	}
	fmt.Printf("%+v", resp[0])
}

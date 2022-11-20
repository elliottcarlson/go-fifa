package go_fifa_test

import (
	"testing"

	fifa "github.com/ImDevinC/go-fifa"
	"github.com/stretchr/testify/assert"
)

func TestGetSeason(t *testing.T) {
	t.Parallel()
	client := fifa.Client{}
	_, err := client.GetSeason(&fifa.GetSeasonOptions{
		SeasonId: "400198457",
	})
	if ok := assert.Nil(t, err, "expected no error with GetSeason, got: %s", err); !ok {
		t.FailNow()
	}
}

func TestGetSeasonStandings(t *testing.T) {
	t.Parallel()
	client := fifa.Client{}
	resp, err := client.GetSeasonStandings(&fifa.GetSeasonStandingsOptions{
		CompetitionId: "2000001049",
		SeasonId:      "400198457",
		StageId:       "400198931",
	})
	if ok := assert.Nil(t, err, "expected no error with GetSeasonStandings, got: %s", err); !ok {
		t.FailNow()
	}
	if ok := assert.Greater(t, len(resp), 0, "expected more than one result, got none"); !ok {
		t.FailNow()
	}
	if ok := assert.Equal(t, resp[0].CompetitionId, "2000001049", "incorrect competition id"); !ok {
		t.FailNow()
	}
}

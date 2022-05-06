package go_fifa_test

import (
	"testing"

	fifa "github.com/ImDevinC/go-fifa"
	"github.com/stretchr/testify/assert"
)

func TestGetCurrentMatches(t *testing.T) {
	client := fifa.Client{}
	_, err := client.GetCurrentMatches()
	if ok := assert.Nil(t, err, "expected no error with GetCurrentMatches, got: %s", err); !ok {
		t.FailNow()
	}
}

func TestGetTodaysMatches(t *testing.T) {
	client := fifa.Client{}
	_, err := client.GetTodaysMatches()
	if ok := assert.Nil(t, err, "expected no error with GetTodaysMatches, got: %s", err); !ok {
		t.FailNow()
	}
}

func TestGetTeammatches(t *testing.T) {
	client := fifa.Client{}
	_, err := client.GetTeamMatches(&fifa.GetTeamMatchesOptions{
		TeamId:        "2000019544",
		SeasonId:      "400198457",
		CompetitionId: "2000001049",
	})
	if ok := assert.Nil(t, err, "expected no error with GetTeamMatches, got: %s", err); !ok {
		t.FailNow()
	}
}

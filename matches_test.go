package go_fifa_test

import (
	"testing"
	"time"

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
	now := time.Now().UTC()
	client := fifa.Client{}
	resp, err := client.GetTodaysMatches()
	if ok := assert.Nil(t, err, "expected no error with GetTodaysMatches, got: %s", err); !ok {
		t.FailNow()
	}
	if ok := assert.GreaterOrEqual(t, len(resp), 1, "expected at least one match"); !ok {
		t.FailNow()
	}
	topHour := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), 0, 0, 0, time.UTC)
	for _, m := range resp {
		if ok := assert.False(t, m.Date.Before(topHour), "date %s should not be before %s", m.Date, now); !ok {
			t.FailNow()
		}
	}
}

func TestGetTeamMatches(t *testing.T) {
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

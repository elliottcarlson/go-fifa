package go_fifa_test

import (
	"fmt"
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
	// now := time.Now()
	client := fifa.Client{}
	resp, err := client.GetTodaysMatches()
	if ok := assert.Nil(t, err, "expected no error with GetTodaysMatches, got: %s", err); !ok {
		t.FailNow()
	}
	if ok := assert.GreaterOrEqual(t, len(resp), 1, "expected at least one match"); !ok {
		t.FailNow()
	}
	// startDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).UTC()
	if ok := assert.Greater(t, len(resp), 0); !ok {
		t.FailNow()
	}
	fmt.Printf("%+v\n", resp[0])
	m := resp[0]
	evts, err := client.GetMatchEvents(&fifa.GetMatchEventOptions{
		CompetitionId: m.CompetitionId,
		SeasonId:      m.SeasonId,
		StageId:       m.StageId,
		MatchId:       m.Id,
	})
	if ok := assert.Nil(t, err, "expected no error with GetMatchEvents, got: %s", err); !ok {
		t.FailNow()
	}
	if ok := assert.Greater(t, len(evts.Event), 0, "expected more results"); !ok {
		t.FailNow()
	}
	fmt.Println(evts.Event[0])
	// for _, m := range resp {
	// 	if ok := assert.False(t, m.Date.Before(startDate), "date %s should not be before %s", m.Date, now); !ok {
	// 		t.FailNow()
	// 	}
	// }
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

func TestGetUpcomingMatches(t *testing.T) {
	now := time.Now()
	client := fifa.Client{}
	resp, err := client.GetUpcomingMatches()
	if ok := assert.Nil(t, err, "expected no error with GetTodaysMatches, got: %s", err); !ok {
		t.FailNow()
	}
	if ok := assert.GreaterOrEqual(t, len(resp), 1, "expected at least one match"); !ok {
		t.FailNow()
	}
	startDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).UTC()
	for _, m := range resp {
		if ok := assert.False(t, m.Date.Before(startDate), "date %s should not be before %s", m.Date, now); !ok {
			t.FailNow()
		}
	}
}

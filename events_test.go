package go_fifa_test

import (
	"testing"

	fifa "github.com/ImDevinC/go-fifa"
	"github.com/stretchr/testify/assert"
)

func TestGetMatchEvents(t *testing.T) {
	t.Parallel()
	client := fifa.Client{}
	resp, err := client.GetMatchEvents(&fifa.GetMatchEventOptions{
		CompetitionId: "17",
		SeasonId:      "255711",
		StageId:       "285063",
		MatchId:       "400128082",
	})
	if ok := assert.Nil(t, err, "expected no error with GetMatchEvents, got: %s", err); !ok {
		t.FailNow()
	}
	if ok := assert.Greater(t, len(resp.Events), 0, "expected at least on result, got none"); !ok {
		t.FailNow()
	}
	if ok := assert.Equal(t, resp.Events[0].Type, fifa.CoinToss, "expected a coin toss event, got %d instead", resp.Events[0].Type); !ok {
		t.Fail()
	}
}

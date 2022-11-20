package go_fifa_test

import (
	"testing"

	fifa "github.com/ImDevinC/go-fifa"
	"github.com/stretchr/testify/assert"
)

func TestGetMatchEvents(t *testing.T) {
	t.Parallel()
	client := fifa.Client{}
	_, err := client.GetMatchEvents(&fifa.GetMatchEventOptions{
		CompetitionId: "17",
		SeasonId:      "255711",
		StageId:       "285063",
		MatchId:       "400128082",
	})
	if ok := assert.Nil(t, err, "expected no error with GetMatchEvents, got: %s", err); !ok {
		t.FailNow()
	}
}

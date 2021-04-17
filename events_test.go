package go_fifa_test

import (
	"testing"

	fifa "github.com/ImDevinC/go-fifa"
	"github.com/stretchr/testify/assert"
)

func TestGetMatchEvents(t *testing.T) {
	client, err := fifa.NewClient(&fifa.Options{})
	if ok := assert.Nil(t, err, "expected no error with NewClient, got: %s", err); !ok {
		t.FailNow()
	}
	_, err = client.GetMatchEvents(&fifa.GetMatchEventOptions{
		CompetitionID: "2000000104",
		SeasonID:      "400156963",
		StageID:       "400157262",
		MatchID:       "400157188",
	})
	if ok := assert.Nil(t, err, "expected no error with GetMatchEvents, got: %s", err); !ok {
		t.FailNow()
	}
}

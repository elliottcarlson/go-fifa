package go_fifa_test

import (
	"testing"

	fifa "github.com/ImDevinC/go-fifa"
	"github.com/stretchr/testify/assert"
)

func TestGetSeason(t *testing.T) {
	t.Parallel()
	client := fifa.Client{}
	resp, err := client.GetSeason(&fifa.GetSeasonOptions{
		SeasonId: "255711",
	})
	if ok := assert.Nil(t, err, "expected no error with GetSeason, got: %s", err); !ok {
		t.FailNow()
	}
	if ok := assert.Equal(t, resp.Abbreviation, "FWC 2022"); !ok {
		t.Fail()
	}
}

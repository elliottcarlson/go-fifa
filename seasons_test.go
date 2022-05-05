package go_fifa_test

import (
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

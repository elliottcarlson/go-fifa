package go_fifa_test

import (
	"testing"

	fifa "github.com/ImDevinC/go-fifa"
	"github.com/stretchr/testify/assert"
)

func TestGetCurrentMatches(t *testing.T) {
	// TODO: We probably need to update this test with a mock client since
	// we can't guarantee matches will always be available
	t.Parallel()
	client := fifa.Client{}
	_, err := client.GetCurrentMatches()
	if ok := assert.Nil(t, err, "expected no error with GetCurrentMatches, got: %s", err); !ok {
		t.FailNow()
	}
}

func TestGetMatches(t *testing.T) {
	t.Parallel()
	client := fifa.Client{}
	resp, err := client.GetMatches(&fifa.GetMatchesOptions{
		CompetitionId: "17",    // World cup
		TeamId:        "43927", // Ecuador
		SeasonId:      "255711",
	})
	if ok := assert.Nil(t, err, "expected no error, got: %s", err); !ok {
		t.FailNow()
	}
	if ok := assert.Greater(t, len(resp), 0, "expected at least on result"); !ok {
		t.FailNow()
	}
}

package go_fifa_test

import (
	"testing"

	fifa "github.com/ImDevinC/go-fifa"
	"github.com/stretchr/testify/assert"
)

func TestGetCompetitions(t *testing.T) {
	client := fifa.Client{}
	resp, err := client.GetCompetitions()
	if ok := assert.Nil(t, err, "expected no error with GetCompetitions, got: %s", err); !ok {
		t.FailNow()
	}
	if ok := assert.Greater(t, len(resp), 1, "expected more than 1 result, got: %d", len(resp)); !ok {
		t.FailNow()
	}
}

func TestGetCompetitionById(t *testing.T) {
	client := fifa.Client{}
	resp, err := client.GetCompetition(&fifa.GetCompetitionsOptions{CompetitionId: "17"})
	if ok := assert.Nil(t, err, "expected no error with GetCompetitions, got: %s", err); !ok {
		t.FailNow()
	}
	if ok := assert.Equal(t, resp.Name[0].Description, "FIFA World Cup™", "expected `FIFA World Cup`, got %s", resp.Name[0].Description); !ok {
		t.FailNow()
	}
}

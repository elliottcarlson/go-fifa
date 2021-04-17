package go_fifa_test

import (
	"testing"

	fifa "github.com/ImDevinC/go-fifa"
	"github.com/stretchr/testify/assert"
)

func TestGetTeam(t *testing.T) {
	client, err := fifa.NewClient(&fifa.Options{})
	if ok := assert.Nil(t, err, "expected no error with NewClient, got: %s", err); !ok {
		t.FailNow()
	}
	_, err = client.GetTeam(&fifa.GetTeamOptions{
		TeamID: "1884381",
	})
	if ok := assert.Nil(t, err, "expected no error with GetTeam, got: %s", err); !ok {
		t.FailNow()
	}
}

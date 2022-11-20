package go_fifa_test

import (
	"testing"

	fifa "github.com/ImDevinC/go-fifa"
	"github.com/stretchr/testify/assert"
)

func TestGetTeam(t *testing.T) {
	t.Parallel()
	client := fifa.Client{}
	_, err := client.GetTeam(&fifa.GetTeamOptions{
		TeamId: "1884381",
	})
	if ok := assert.Nil(t, err, "expected no error with GetTeam, got: %s", err); !ok {
		t.FailNow()
	}
}

package go_fifa_test

import (
	"testing"

	fifa "github.com/ImDevinC/go-fifa"
	"github.com/stretchr/testify/assert"
)

func TestGetTeam(t *testing.T) {
	t.Parallel()
	client := fifa.Client{}
	resp, err := client.GetTeam(&fifa.GetTeamOptions{
		TeamId: "43921",
	})
	if ok := assert.Nil(t, err, "expected no error with GetTeam, got: %s", err); !ok {
		t.FailNow()
	}
	if resp.CountryId != "USA" {
		t.Errorf("expected countryId USA but got %s", resp.CountryId)
	}
}

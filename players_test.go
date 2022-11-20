package go_fifa_test

import (
	"testing"

	fifa "github.com/ImDevinC/go-fifa"
	"github.com/stretchr/testify/assert"
)

func TestGetPlayer(t *testing.T) {
	t.Parallel()
	client := fifa.Client{}
	resp, err := client.GetPlayer(&fifa.GetPlayerOptions{
		PlayerId: "200704206",
	})
	if ok := assert.Nil(t, err, "expected no error with GetTeam, got: %s", err); !ok {
		t.FailNow()
	}
	if ok := assert.NotEmpty(t, resp.Name, "no names available"); !ok {
		t.FailNow()
	}
	if ok := assert.Equal(t, "Edgar HERNÁNDEZ", resp.Name[0].Description, "got unexpected value"); !ok {
		t.FailNow()
	}
}

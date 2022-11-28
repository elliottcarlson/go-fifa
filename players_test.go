package go_fifa_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	fifa "github.com/imdevinc/go-fifa"
	"github.com/stretchr/testify/assert"
)

var expectedPlayer = fifa.PlayerResponse{
	Id: "229397",
	Name: []fifa.LocaleDescription{{
		Locale:      "en-GB",
		Description: "Lionel MESSI",
	}},
	Alias: []fifa.LocaleDescription{{
		Locale:      "en-GB",
		Description: "MESSI",
	}},
	BirthDate:                time.Date(1987, 6, 24, 0, 0, 0, 0, time.UTC),
	Weight:                   72,
	Height:                   170,
	Birthplace:               "Rosario",
	CountryId:                "ARG",
	InternationalCaps:        167,
	InternationalDebut:       nil,
	TopCompetitionDebut:      nil,
	TwitterAccount:           nil,
	PreferredFoot:            nil,
	MediaContent:             []any{},
	LocalizedTwitterAccounts: nil,
	Goals:                    93,
	Properties: fifa.Properties{
		StatsPerformIfesId: "229397",
		IdStatsPerform:     "c5ryhn04g9goikd0blmh83aol",
		IdIFES:             "229397",
	},
	IsUpdateable: false,
}

func TestGetPlayer(t *testing.T) {
	t.Parallel()
	client := fifa.Client{
		Client: &TestClient{},
	}
	resp, err := client.GetPlayer(&fifa.GetPlayerOptions{
		PlayerId: "229397",
	})
	if ok := assert.Nil(t, err, "expected no error with GetTeam, got: %s", err); !ok {
		t.FailNow()
	}
	diff := cmp.Diff(*resp, expectedPlayer)
	if ok := assert.Equal(t, diff, "", diff); !ok {
		t.Fail()
	}
}

package go_fifa_test

import (
	"testing"
	"time"

	fifa "github.com/ImDevinC/go-fifa"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

var expectedPlayer = fifa.PlayerResponse{
	Id: "229397",
	Name: []fifa.DefaultDescriptionResponse{{
		Locale:      "en-GB",
		Description: "Lionel MESSI",
	}},
	Alias: []fifa.DefaultDescriptionResponse{{
		Locale:      "en-GB",
		Description: "MESSI",
	}},
	Birthdate:                time.Date(1987, 6, 24, 0, 0, 0, 0, time.UTC),
	Weight:                   72,
	Height:                   170,
	Birthplace:               "Rosario",
	CountryId:                "ARG",
	InternationalCaps:        166,
	InternationalDebut:       0,
	TopCompetitionDebut:      0,
	PictureURL:               "",
	ThumbnailURL:             "",
	TwitterAccount:           "",
	PreferredFoot:            "",
	MediaContent:             []string{},
	LocalizedTwitterAccounts: nil,
	Goals:                    92,
	Properties: map[string]any{
		"StatsPerformIfesId": "229397",
		"IdStatsPerform":     "c5ryhn04g9goikd0blmh83aol",
		"IdIFES":             "229397",
	},
	IsUpdateable: false,
}

func TestGetPlayer(t *testing.T) {
	t.Parallel()
	client := fifa.Client{}
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

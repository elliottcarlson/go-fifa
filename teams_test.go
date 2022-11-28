package go_fifa_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	fifa "github.com/imdevinc/go-fifa"
	"github.com/stretchr/testify/assert"
)

var expectedTeam = fifa.TeamResponse{
	Id:              "43921",
	ConfederationId: "CONCACAF",
	Type:            1,
	AgeType:         7,
	FootballType:    0,
	Gender:          fifa.MaleGender,
	Name: []fifa.LocaleDescription{{
		Locale:      "en-GB",
		Description: "United States",
	}},
	AssociationId: "USA",
	City:          "CHICAGO",
	PostalCode:    "IL 60616",
	ShortClubName: "USA",
	CountryId:     "USA",
	Abbreviation:  "USA",
	Street:        "US Soccer Federation, US Soccer House, 1801 S. Prairie Avenue",
	PictureUrl:    "https://api.fifa.com/api/v3/picture/flags-{format}-{size}/USA",
	DisplayName: []fifa.LocaleDescription{{
		Locale:      "en-GB",
		Description: "USA",
	}},
	Content: []any{},
	Properties: fifa.Properties{
		IdIFES:             "43921",
		StatsPerformIfesId: "43921",
		IdStatsPerform:     "9vh2u1p4ppm597tjfahst2m3n",
	},
	IsUpdateable: false,
}

func TestGetTeam(t *testing.T) {
	t.Parallel()
	client := fifa.Client{}
	resp, err := client.GetTeam(&fifa.GetTeamOptions{
		TeamId: "43921",
	})
	if ok := assert.Nil(t, err, "expected no error with GetTeam, got: %s", err); !ok {
		t.FailNow()
	}
	diff := cmp.Diff(expectedTeam, *resp)
	if ok := assert.Equal(t, diff, "", diff); !ok {
		t.Fail()
	}
}

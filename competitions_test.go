package go_fifa_test

import (
	"testing"

	fifa "github.com/ImDevinC/go-fifa"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

var expCompetitionsResp = []fifa.Competition{{
	Id: "2000000000",
	Name: []fifa.LocaleDescription{{
		Locale:      "en-gb",
		Description: "Premier League",
	}},
	ConfederationIds:     []string{"UEFA"},
	MemberAssociationIds: []string{"ENG"},
	OwnerId:              "UEFA",
	Gender:               fifa.MaleGender,
	FootballType:         0,
	TeamType:             0,
	Type:                 2,
	AgeType:              7,
	Properties: fifa.Properties{
		IdStatsPerform: "2kwbbcootiqqgmrzs6o5inle5",
	},
	IsUpdateable: false,
}}

func TestGetCompetitions(t *testing.T) {
	t.Parallel()
	client := fifa.Client{
		Client: &TestClient{},
	}
	resp, err := client.GetCompetitions()
	if ok := assert.Nil(t, err, "expected no error with GetCompetitions, got: %s", err); !ok {
		t.FailNow()
	}
	diff := cmp.Diff(resp, expCompetitionsResp)
	if ok := assert.Equal(t, diff, "", diff); !ok {
		t.Fail()
	}
}

func TestGetCompetitionById(t *testing.T) {
	t.Parallel()
	client := fifa.Client{
		Client: &TestClient{},
	}
	resp, err := client.GetCompetition(&fifa.GetCompetitionsOptions{CompetitionId: "2000000000"})
	if ok := assert.Nil(t, err, "expected no error with GetCompetitions, got: %s", err); !ok {
		t.FailNow()
	}
	diff := cmp.Diff(resp, &expCompetitionsResp[0])
	if ok := assert.Equal(t, diff, "", diff); !ok {
		t.Fail()
	}
}

package go_fifa_test

import (
	"testing"
	"time"

	fifa "github.com/ImDevinC/go-fifa"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

var expectedSeason = fifa.SeasonResponse{
	Id: "255711",
	Name: []fifa.DefaultDescriptionResponse{{
		Locale:      "en-GB",
		Description: "FIFA World Cup Qatar 2022™",
	}},
	ShortName:    []fifa.DefaultDescriptionResponse{},
	Abbreviation: "FWC 2022",
	MemberAssociations: []string{
		"QAT", "GER", "DEN", "BRA", "FRA", "BEL", "CRO", "ESP", "SRB", "ENG", "SUI", "NED", "ARG", "IRN", "KOR", "JPN", "KSA", "ECU", "URU", "CAN", "GHA", "SEN", "POR", "POL", "MAR", "TUN", "CMR", "USA", "MEX", "AUS", "CRC", "WAL",
	},
	Confederations:      []string{"AFC", "UEFA", "CONMEBOL", "CONCACAF", "CAF"},
	CompetitionId:       "17",
	StartDate:           time.Date(2022, 11, 20, 16, 0, 0, 0, time.UTC),
	EndDate:             time.Date(2022, 12, 18, 0, 0, 0, 0, time.UTC),
	PictureURL:          "https://api.fifa.com/api/v3/picture/tournaments-{format}-{size}/255711",
	MascotPictureURL:    "https://api.fifa.com/api/v3/picture/tournaments-mascot-{format}-{size}/255711",
	MatchBallPictureURL: "https://api.fifa.com/api/v3/picture/tournaments-matchball-{format}-{size}/255711",
	HostTeams: []any{map[string]any{
		"IdTeam": string("43834"),
	}},
	SportType: 0,
	Properties: fifa.Properties{
		IdIFES:    "255711",
		Providers: "IMM",
	},
	IsUpdateable: false,
}

func TestGetSeason(t *testing.T) {
	t.Parallel()
	client := fifa.Client{}
	resp, err := client.GetSeason(&fifa.GetSeasonOptions{
		SeasonId: "255711",
	})
	if ok := assert.Nil(t, err, "expected no error with GetSeason, got: %s", err); !ok {
		t.FailNow()
	}
	diff := cmp.Diff(*resp, expectedSeason)
	if ok := assert.Equal(t, diff, "", diff); !ok {
		t.Fail()
	}
}

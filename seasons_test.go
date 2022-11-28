package go_fifa_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	fifa "github.com/imdevinc/go-fifa"
	"github.com/stretchr/testify/assert"
)

var expectedSeason = fifa.SeasonResponse{
	Id: "255711",
	Name: []fifa.LocaleDescription{{
		Locale:      "en-GB",
		Description: "FIFA World Cup Qatar 2022™",
	}},
	Abbreviation: "FWC 2022",
	MemberAssociationIds: []string{
		"QAT", "GER", "DEN", "BRA", "FRA", "BEL", "CRO", "ESP", "SRB", "ENG", "SUI", "NED", "ARG", "IRN", "KOR", "JPN", "KSA", "ECU", "URU", "CAN", "GHA", "SEN", "POR", "POL", "MAR", "TUN", "CMR", "USA", "MEX", "AUS", "CRC", "WAL",
	},
	ConfederationIds:    []string{"AFC", "UEFA", "CONMEBOL", "CONCACAF", "CAF"},
	CompetitionId:       "17",
	StartDate:           time.Date(2022, 11, 20, 16, 0, 0, 0, time.UTC),
	EndDate:             time.Date(2022, 12, 18, 0, 0, 0, 0, time.UTC),
	PictureUrl:          "https://api.fifa.com/api/v3/picture/tournaments-{format}-{size}/255711",
	MascotPictureUrl:    "https://api.fifa.com/api/v3/picture/tournaments-mascot-{format}-{size}/255711",
	MatchBallPictureUrl: "https://api.fifa.com/api/v3/picture/tournaments-matchball-{format}-{size}/255711",
	HostTeams: []fifa.SeasonHostTeam{{
		TeamId: "43834",
	}},
	SportType: 0,
	Properties: fifa.Properties{
		IdIFES:    "255711",
		Providers: "IMM",
	},
	IsUpdateable: false,
	ShortName:    []string{},
}

func TestGetSeason(t *testing.T) {
	t.Parallel()
	client := fifa.Client{
		Client: &TestClient{},
	}
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

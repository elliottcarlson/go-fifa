package go_fifa_test

import (
	"testing"
	"time"

	fifa "github.com/ImDevinC/go-fifa"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

var expectedMatch = fifa.MatchResponse{
	Id:              "400128082",
	StageId:         "285063",
	GroupId:         "285065",
	SeasonId:        "255711",
	SeasonShortName: []any{},
	CompetitionId:   "17",
	Weather: fifa.WeatherResponse{
		Humidity:    "61",
		Temperature: "24",
		WindSpeed:   "8",
		Type:        10,
		TypeLocalized: []fifa.DefaultDescriptionResponse{{
			Locale:      "en-GB",
			Description: "Clear Night",
		}},
	},
	StageName: []fifa.DefaultDescriptionResponse{{
		Locale:      "en-GB",
		Description: "First stage",
	}},
	GroupName: []fifa.DefaultDescriptionResponse{{
		Locale:      "en-GB",
		Description: "Group A",
	}},
	Competition: []fifa.DefaultDescriptionResponse{{
		Locale:      "en-GB",
		Description: "FIFA World Cup™",
	}},
	Season: []fifa.DefaultDescriptionResponse{{
		Locale:      "en-GB",
		Description: "FIFA World Cup Qatar 2022™",
	}},
	Date:      time.Date(2022, 11, 20, 16, 00, 00, 0, time.UTC),
	LocalDate: time.Date(2022, 11, 20, 19, 00, 00, 0, time.UTC),
	HomeTeam: fifa.MatchTeamResponse{
		Score:      0,
		Side:       "",
		Id:         "43834",
		PictureURL: "https://api.fifa.com/api/v3/picture/flags-{format}-{size}/QAT",
		CountryId:  "QAT",
		Tactics:    "5-3-2",
		Type:       1,
		AgeType:    7,
		Name: []fifa.DefaultDescriptionResponse{{
			Locale:      "en-GB",
			Description: "Qatar",
		}},
		Abbreviation:  "QAT",
		ShortClubName: "Qatar",
		FootballType:  0,
		Gender:        fifa.MaleGender,
		AssociationId: "QAT",
	},
	AwayTeam: fifa.MatchTeamResponse{
		Score:      2,
		Side:       "",
		Id:         "43927",
		PictureURL: "https://api.fifa.com/api/v3/picture/flags-{format}-{size}/ECU",
		CountryId:  "ECU",
		Tactics:    "4-4-2",
		Type:       1,
		AgeType:    7,
		Name: []fifa.DefaultDescriptionResponse{{
			Locale:      "en-GB",
			Description: "Ecuador",
		}},
		Abbreviation:  "ECU",
		ShortClubName: "Ecuador",
		FootballType:  0,
		Gender:        fifa.MaleGender,
		AssociationId: "ECU",
	},
	HomeTeamScore:          0,
	AwayTeamScore:          2,
	AggregateHomeTeamScore: 0,
	AggregateAwayTeamScore: 0,
	HomeTeamPenaltyScore:   0,
	AwayTeamPenaltyScore:   0,
	Stadium: fifa.StadiumResponse{
		Id: "400090319",
		Name: []fifa.DefaultDescriptionResponse{{
			Locale:      "en-GB",
			Description: "Al Bayt Stadium",
		}},
		Capacity:   0,
		WebAddress: "",
		Built:      "",
		HasRoof:    false,
		Turf:       "",
		CityId:     "400030664",
		City: []fifa.DefaultDescriptionResponse{{
			Locale:      "en-GB",
			Description: "Al Khor",
		}},
		CountryId:          "QAT",
		PostalCode:         "",
		Street:             "",
		Email:              "",
		Fax:                "",
		Phone:              "",
		AffiliationCountry: "",
		AffiliationRegion:  "",
		Latitude:           0,
		Longitude:          0,
		Length:             "",
		Width:              "",
		Properties: map[string]any{
			"IdIFES": string("5047367"),
		},
		IsUpdateable: false,
	},
	MatchTime:           "0'",
	MatchDay:            "1",
	SecondHalfTime:      "",
	FirstHalfTime:       "",
	FirstHalfExtraTime:  0,
	SecondHalfExtraTime: 0,
	WinnerId:            "43927",
	BallPossession: fifa.BallPossessionResponse{
		Intervals:   []any{},
		LastX:       []any{},
		OverallHome: 47.088547,
		OverallAway: 52.911457,
	},
	Officials: []fifa.OfficialResponse{
		{
			CountryId: "ITA",
			Id:        "315593",
			ShortName: []fifa.DefaultDescriptionResponse{{
				Locale:      "en-GB",
				Description: "Daniele ORSATO",
			}},
			Name: []fifa.DefaultDescriptionResponse{{
				Locale:      "en-GB",
				Description: "Daniele Orsato",
			}},
			Type: 1,
			TypeLocalized: []fifa.DefaultDescriptionResponse{{
				Locale:      "en-GB",
				Description: "Referee",
			}},
		},
		{
			CountryId: "ITA",
			Id:        "415967",
			ShortName: []fifa.DefaultDescriptionResponse{{
				Locale:      "en-GB",
				Description: "Ciro CARBONE",
			}},
			Name: []fifa.DefaultDescriptionResponse{{
				Locale:      "en-GB",
				Description: "Ciro Carbone",
			}},
			Type: 2,
			TypeLocalized: []fifa.DefaultDescriptionResponse{{
				Locale:      "en-GB",
				Description: "Assistant Referee 1",
			}},
		},
		{
			CountryId: "ITA",
			Id:        "360574",
			ShortName: []fifa.DefaultDescriptionResponse{{
				Locale:      "en-GB",
				Description: "Alessandro GIALLATINI",
			}},
			Name: []fifa.DefaultDescriptionResponse{{
				Locale:      "en-GB",
				Description: "Alessandro Giallatini",
			}},
			Type: 3,
			TypeLocalized: []fifa.DefaultDescriptionResponse{{
				Locale:      "en-GB",
				Description: "Assistant Referee 2",
			}},
		},
		{
			CountryId: "ROU",
			Id:        "314605",
			ShortName: []fifa.DefaultDescriptionResponse{{
				Locale:      "en-GB",
				Description: "Istvan KOVACS",
			}},
			Name: []fifa.DefaultDescriptionResponse{{
				Locale:      "en-GB",
				Description: "Istvan Kovacs",
			}},
			Type: 4,
			TypeLocalized: []fifa.DefaultDescriptionResponse{{
				Locale:      "en-GB",
				Description: "Fourth official",
			}},
		},
		{
			CountryId: "ITA",
			Id:        "399108",
			ShortName: []fifa.DefaultDescriptionResponse{{
				Locale:      "en-GB",
				Description: "Massimiliano IRRATI",
			}},
			Name: []fifa.DefaultDescriptionResponse{{
				Locale:      "en-GB",
				Description: "Massimiliano Irrati",
			}},
			Type: 5,
			TypeLocalized: []fifa.DefaultDescriptionResponse{{
				Locale:      "en-GB",
				Description: "Video Assistant Referee (VAR)",
			}},
		},
		{
			CountryId: "POL",
			Id:        "327977",
			ShortName: []fifa.DefaultDescriptionResponse{{
				Locale:      "en-GB",
				Description: "Tomasz LISTKIEWICZ",
			}},
			Name: []fifa.DefaultDescriptionResponse{{
				Locale:      "en-GB",
				Description: "Tomasz LISTKIEWICZ",
			}},
			Type: 7,
			TypeLocalized: []fifa.DefaultDescriptionResponse{{
				Locale:      "en-GB",
				Description: "Offside VAR",
			}},
		},
		{
			CountryId: "ITA",
			Id:        "328434",
			ShortName: []fifa.DefaultDescriptionResponse{{
				Locale:      "en-GB",
				Description: "Paolo VALERI",
			}},
			Name: []fifa.DefaultDescriptionResponse{{
				Locale:      "en-GB",
				Description: "Paolo Valeri",
			}},
			Type: 8,
			TypeLocalized: []fifa.DefaultDescriptionResponse{{
				Locale:      "en-GB",
				Description: "Assistant VAR",
			}},
		},
		{
			CountryId: "FRA",
			Id:        "373782",
			ShortName: []fifa.DefaultDescriptionResponse{{
				Locale:      "en-GB",
				Description: "Benoit MILLOT",
			}},
			Name: []fifa.DefaultDescriptionResponse{{
				Locale:      "en-GB",
				Description: "Benoit MILLOT",
			}},
			Type: 9,
			TypeLocalized: []fifa.DefaultDescriptionResponse{{
				Locale:      "en-GB",
				Description: "Support VAR",
			}},
		},
		{
			CountryId: "ROU",
			Id:        "361587",
			ShortName: []fifa.DefaultDescriptionResponse{{
				Locale:      "en-GB",
				Description: "Mihai ARTENE",
			}},
			Name: []fifa.DefaultDescriptionResponse{{
				Locale:      "en-GB",
				Description: "Mihai Ovidiu Artene",
			}},
			Type: 10,
			TypeLocalized: []fifa.DefaultDescriptionResponse{{
				Locale:      "en-GB",
				Description: "Reserve Assistant Referee",
			}},
		},
	},
	PlaceHolderA:      "A1",
	PlaceHolderB:      "A2",
	Status:            0,
	ResultType:        1,
	TimeDefined:       true,
	OfficialityStatus: 1,
	MatchNumber:       1,
	Properties: map[string]any{
		"IdIFES": "128084",
	},
	IsUpdateable: false,
}

func TestGetCurrentMatches(t *testing.T) {
	// TODO: We probably need to update this test with a mock client since
	// we can't guarantee matches will always be available
	t.Parallel()
	client := fifa.Client{}
	_, err := client.GetCurrentMatches()
	if ok := assert.Nil(t, err, "expected no error with GetCurrentMatches, got: %s", err); !ok {
		t.FailNow()
	}
}

func TestGetMatches(t *testing.T) {
	t.Parallel()
	client := fifa.Client{}
	resp, err := client.GetMatches(&fifa.GetMatchesOptions{
		CompetitionId: "17",     // World cup
		TeamId:        "43927",  // Ecuador
		SeasonId:      "255711", // World Cup Qatar 2022
		Count:         1,
	})
	if ok := assert.Nil(t, err, "expected no error, got: %s", err); !ok {
		t.FailNow()
	}
	if ok := assert.Greater(t, len(resp), 0, "expected at least on result"); !ok {
		t.FailNow()
	}
	diff := cmp.Diff(resp[0], expectedMatch)
	if ok := assert.Equal(t, diff, "", diff); !ok {
		t.Fail()
	}
}

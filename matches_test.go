package go_fifa_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	fifa "github.com/imdevinc/go-fifa"
	"github.com/stretchr/testify/assert"
)

var expLiveMatches = fifa.LiveMatch{
	MatchId:       "bn0ffkqrvppd3jahiflnzs1lg",
	StageId:       "9d5x17v8h81sp5r6q4hbx5hqs",
	SeasonId:      "9cy0d6sf8dui1o0rpwynx3190",
	CompetitionId: "af79lqrc0ntom74zq13ccjslo",
	CompetitionName: []fifa.LocaleDescription{{
		Locale:      "en-gb",
		Description: "Challenger Pro League",
	}},
	SeasonName: []fifa.LocaleDescription{{
		Locale:      "en-gb",
		Description: "Challenger Pro League 2022/2023",
	}},
	SeasonShortName: []any{},
	Stadium: fifa.Stadium{
		Name: []fifa.LocaleDescription{{
			Locale:      "en-gb",
			Description: "Lotto Park",
		}},
		Capacity: 28063,
		Built:    time.Date(1917, 01, 01, 0, 0, 0, 0, time.UTC),
		HasRoof:  false,
		CityId:   "400241509",
		CityName: []fifa.LocaleDescription{{
			Locale:      "en-GB",
			Description: "Bruxelles (Brussel)",
		}},
		CountryId: "BEL",
		Street:    "Avenue Théo Verbeeck 2",
		Latitude:  50.834182156,
		Longitude: 4.298352599,
	},
	ResultType:          1,
	MatchDay:            "15",
	Date:                time.Date(2022, 11, 26, 19, 0, 0, 0, time.UTC),
	LocalDate:           time.Date(2022, 11, 26, 20, 0, 0, 0, time.UTC),
	MatchTime:           "66'",
	FirstHalfExtraTime:  0,
	SecondHalfExtraTime: 0,
	Period:              fifa.SecondPeriod,
	HomeTeam: fifa.Team{
		Score:      1,
		Id:         "d011rbp4etj5styzfmpd7wa38",
		PictureUrl: "https://api.fifa.com/api/v3/picture/teams-{format}-{size}/d011rbp4etj5styzfmpd7wa38",
		CountryId:  "BEL",
		Type:       0,
		AgeType:    7,
		Tactics:    "5-3-2",
		Name: []fifa.LocaleDescription{{
			Locale:      "en-gb",
			Description: "RSC Anderlecht II",
		}},
		Abbreviation: "ANF",
		Coaches: []fifa.TeamCoach{{
			Id:        "4dvbxh16w3txv1jy5e2mat638",
			CountryId: "NED",
			Name: []fifa.LocaleDescription{{
				Locale:      "en-gb",
				Description: "Robin Veldeman",
			}},
			Alias: []fifa.LocaleDescription{{
				Locale:      "en-gb",
				Description: "R. Veldman",
			}},
			Role: 0,
		}},
		Players: []fifa.TeamPlayer{{
			Id:          "c716rboejsazxhoosf6l0v6dw",
			TeamId:      "d011rbp4etj5styzfmpd7wa38",
			ShirtNumber: 63,
			Status:      1,
			Captain:     false,
			Name: []fifa.LocaleDescription{{
				Locale:      "en-gb",
				Description: "Timon Vanhoutte",
			}},
			ShortName: []fifa.LocaleDescription{{
				Locale:      "en-gb",
				Description: "T. Vanhoutte",
			}},
			Position:    0,
			FieldStatus: 1,
		}},
		Bookings: []fifa.Booking{},
		Goals: []fifa.Goal{{
			Type:           2,
			PlayerId:       "rao8xq2rg3beoc418zv1gzro",
			Minute:         "11:35",
			AssistPlayerId: "358110",
			Period:         3,
			TeamId:         "d011rbp4etj5styzfmpd7wa38",
		}},
		Substitutions: []fifa.TeamSubstitution{{
			Period:      fifa.SecondPeriod,
			Reason:      2,
			Position:    3,
			PlayerOffId: "71euf80lf18fu40b9lika9ysq",
			PlayerOnId:  "dzl3nhd1wmcttn6qkm0sqb7ro",
			PlayerOffName: []fifa.LocaleDescription{{
				Locale:      "en-gb",
				Description: "Nils De Wilde",
			}},
			PlayerOnName: []fifa.LocaleDescription{{
				Locale:      "en-gb",
				Description: "Simion Michez",
			}},
			Minute: "64:52",
			TeamId: "d011rbp4etj5styzfmpd7wa38",
		}},
		FootballType:  0,
		Gender:        fifa.MaleGender,
		AssociationId: "BEL",
		ShortClubName: "Anderlecht II",
	},
	AwayTeam: fifa.Team{
		Score:      2,
		Id:         "2000019799",
		PictureUrl: "https://api.fifa.com/api/v3/picture/teams-{format}-{size}/2000019799",
		CountryId:  "BEL",
		Type:       0,
		AgeType:    7,
		Tactics:    "3-4-3",
		Name: []fifa.LocaleDescription{{
			Locale:      "en-gb",
			Description: "Beerschot",
		}},
		Abbreviation: "BEE",
		Coaches: []fifa.TeamCoach{{
			Id:        "co3yqoz2wqu4bx8jh2t8gjrl6",
			CountryId: "AUT",
			Name: []fifa.LocaleDescription{{
				Locale:      "en-gb",
				Description: "Andreas Wieland",
			}},
			Alias: []fifa.LocaleDescription{{
				Locale:      "en-gb",
				Description: "A. Wieland",
			}},
			Role: 0,
		}},
		Players: []fifa.TeamPlayer{{
			Id:          "8r3r6z1khjq55pweqzfqezpm2",
			TeamId:      "2000019799",
			ShirtNumber: 1,
			Status:      1,
			Captain:     true,
			Name: []fifa.LocaleDescription{{
				Locale:      "en-gb",
				Description: "Bill Lathouwers",
			}},
			ShortName: []fifa.LocaleDescription{{
				Locale:      "en-gb",
				Description: "B. Lathouwers",
			}},
			Position:    0,
			FieldStatus: 1,
		}},
		Bookings: []fifa.Booking{{
			Card:     1,
			Period:   fifa.FirstPeriod,
			PlayerId: "2ky2kn7gpjorkyg9zyg68pk6i",
			TeamId:   "2000019799",
			Minute:   "38:56",
			Reason:   "OffTheBallFoul",
		}},
		Goals: []fifa.Goal{{
			Type:           2,
			PlayerId:       "cfxzvlgvt9jj17qxq41t9sxcl",
			Minute:         "33:28",
			AssistPlayerId: "7q0v538ed6j5xy3wgykh56w2x",
			Period:         fifa.FirstPeriod,
			TeamId:         "2000019799",
		}},
		Substitutions: []fifa.TeamSubstitution{},
		FootballType:  0,
		Gender:        fifa.MaleGender,
		AssociationId: "BEL",
		ShortClubName: "Beerschot",
	},
	BallPossession: fifa.BallPossession{
		Intervals:   []any{},
		LastX:       []any{},
		OverallHome: 54.2,
		OverallAway: 45.8,
	},
	Officials: []fifa.Official{{
		CountryId: "BEL",
		Id:        "4sabn8pw0tbyrrnwcch72o05x",
		ShortName: []fifa.LocaleDescription{{
			Locale:      "en-gb",
			Description: "Q. Pirard",
		}},
		Name: []fifa.LocaleDescription{{
			Locale:      "en-gb",
			Description: "Quentin Pirard",
		}},
		Type: 1,
		TypeLocalized: []fifa.LocaleDescription{{
			Locale:      "en-GB",
			Description: "Referee",
		}},
	}},
	MatchStatus: 3,
	GroupName:   []any{},
	StageName: []fifa.LocaleDescription{{
		Locale:      "en-gb",
		Description: "Regular Season",
	}},
	OfficialityStatus: 0,
	TimeDefined:       true,
	Properties: fifa.Properties{
		IdStatsPerform: "bn0ffkqrvppd3jahiflnzs1lg",
	},
	IsUpdateable: false,
}

func TestGetCurrentMatches(t *testing.T) {
	// TODO: We probably need to update this test with a mock client since
	// we can't guarantee matches will always be available
	t.Parallel()
	client := fifa.Client{
		Client: &TestClient{},
	}
	resp, err := client.GetCurrentMatches()
	if ok := assert.Nil(t, err, "expected no error with GetCurrentMatches, got: %s", err); !ok {
		t.FailNow()
	}

	match := resp[0]
	diff := cmp.Diff(match, expLiveMatches)
	if ok := assert.Equal(t, diff, "", diff); !ok {
		t.Error()
	}
}

var expMatches = fifa.MatchData{
	CompetitionId: "17",
	SeasonId:      "255711",
	StageId:       "285063",
	GroupId:       "285065",
	Weather: fifa.Weather{
		Humidity:    "61",
		Temperature: "24",
		WindSpeed:   "8",
		Type:        10,
		TypeLocalized: []fifa.LocaleDescription{{
			Locale:      "en-GB",
			Description: "Clear Night",
		}},
	},
	Attendance: "67372",
	Id:         "400128082",
	StageName: []fifa.LocaleDescription{{
		Locale:      "en-GB",
		Description: "First stage",
	}},
	GroupName: []fifa.LocaleDescription{{
		Locale:      "en-GB",
		Description: "Group A",
	}},
	CompetitionName: []fifa.LocaleDescription{{
		Locale:      "en-GB",
		Description: "FIFA World Cup™",
	}},
	SeasonName: []fifa.LocaleDescription{{
		Locale:      "en-GB",
		Description: "FIFA World Cup Qatar 2022™",
	}},
	SeasonShortName: []any{},
	Date:            time.Date(2022, 11, 20, 16, 0, 0, 0, time.UTC),
	LocalDate:       time.Date(2022, 11, 20, 19, 0, 0, 0, time.UTC),
	Home: fifa.Team{
		Score:      0,
		Id:         "43834",
		PictureUrl: "https://api.fifa.com/api/v3/picture/flags-{format}-{size}/QAT",
		CountryId:  "QAT",
		Tactics:    "5-3-2",
		Type:       1,
		AgeType:    7,
		Name: []fifa.LocaleDescription{{
			Locale:      "en-GB",
			Description: "Qatar",
		}},
		Abbreviation:  "QAT",
		ShortClubName: "Qatar",
		FootballType:  0,
		Gender:        fifa.MaleGender,
		AssociationId: "QAT",
	},
	Away: fifa.Team{
		Score:      2,
		Id:         "43927",
		PictureUrl: "https://api.fifa.com/api/v3/picture/flags-{format}-{size}/ECU",
		CountryId:  "ECU",
		Tactics:    "4-4-2",
		Type:       1,
		AgeType:    7,
		Name: []fifa.LocaleDescription{{
			Locale:      "en-GB",
			Description: "Ecuador",
		}},
		Abbreviation:  "ECU",
		ShortClubName: "Ecuador",
		FootballType:  0,
		Gender:        fifa.MaleGender,
		AssociationId: "ECU",
	},
	HomeTeamScore:        0,
	AwayTeamScore:        2,
	HomeTeamPenaltyScore: 0,
	AwayTeamPenaltyScore: 0,
	IsHomeMatch:          false,
	Stadium: fifa.Stadium{
		Id: "400090319",
		Name: []fifa.LocaleDescription{{
			Locale:      "en-GB",
			Description: "Al Bayt Stadium",
		}},
		HasRoof: false,
		CityId:  "400030664",
		CityName: []fifa.LocaleDescription{{
			Locale:      "en-GB",
			Description: "Al Khor",
		}},
		CountryId: "QAT",
		Properties: fifa.Properties{
			IdIFES: "5047367",
		},
	},
	IsTicketSalesAllowed: false,
	MatchTime:            "0'",
	Winner:               "43927",
	PlaceholderA:         "A1",
	PlaceholderB:         "A2",
	BallPossession: fifa.BallPossession{
		Intervals:   []any{},
		LastX:       []any{},
		OverallHome: 47.088547,
		OverallAway: 52.911457,
	},
	Officials: []fifa.Official{
		{
			CountryId: "ITA",
			Id:        "315593",
			ShortName: []fifa.LocaleDescription{{
				Locale:      "en-GB",
				Description: "Daniele ORSATO",
			}},
			Name: []fifa.LocaleDescription{{
				Locale:      "en-GB",
				Description: "Daniele Orsato",
			}},
			Type: 1,
			TypeLocalized: []fifa.LocaleDescription{{
				Locale:      "en-GB",
				Description: "Referee",
			}},
		},
		{
			CountryId: "ITA",
			Id:        "415967",
			ShortName: []fifa.LocaleDescription{{
				Locale:      "en-GB",
				Description: "Ciro CARBONE",
			}},
			Name: []fifa.LocaleDescription{{
				Locale:      "en-GB",
				Description: "Ciro Carbone",
			}},
			Type: 2,
			TypeLocalized: []fifa.LocaleDescription{{
				Locale:      "en-GB",
				Description: "Assistant Referee 1",
			}},
		},
		{
			CountryId: "ITA",
			Id:        "360574",
			ShortName: []fifa.LocaleDescription{{
				Locale:      "en-GB",
				Description: "Alessandro GIALLATINI",
			}},
			Name: []fifa.LocaleDescription{{
				Locale:      "en-GB",
				Description: "Alessandro Giallatini",
			}},
			Type: 3,
			TypeLocalized: []fifa.LocaleDescription{{
				Locale:      "en-GB",
				Description: "Assistant Referee 2",
			}},
		},
		{
			CountryId: "ROU",
			Id:        "314605",
			ShortName: []fifa.LocaleDescription{{
				Locale:      "en-GB",
				Description: "Istvan KOVACS",
			}},
			Name: []fifa.LocaleDescription{{
				Locale:      "en-GB",
				Description: "Istvan Kovacs",
			}},
			Type: 4,
			TypeLocalized: []fifa.LocaleDescription{{
				Locale:      "en-GB",
				Description: "Fourth official",
			}},
		},
		{
			CountryId: "ITA",
			Id:        "399108",
			ShortName: []fifa.LocaleDescription{{
				Locale:      "en-GB",
				Description: "Massimiliano IRRATI",
			}},
			Name: []fifa.LocaleDescription{{
				Locale:      "en-GB",
				Description: "Massimiliano Irrati",
			}},
			Type: 5,
			TypeLocalized: []fifa.LocaleDescription{{
				Locale:      "en-GB",
				Description: "Video Assistant Referee (VAR)",
			}},
		},
		{
			CountryId: "POL",
			Id:        "327977",
			ShortName: []fifa.LocaleDescription{{
				Locale:      "en-GB",
				Description: "Tomasz LISTKIEWICZ",
			}},
			Name: []fifa.LocaleDescription{{
				Locale:      "en-GB",
				Description: "Tomasz LISTKIEWICZ",
			}},
			Type: 7,
			TypeLocalized: []fifa.LocaleDescription{{
				Locale:      "en-GB",
				Description: "Offside VAR",
			}},
		},
		{
			CountryId: "ITA",
			Id:        "328434",
			ShortName: []fifa.LocaleDescription{{
				Locale:      "en-GB",
				Description: "Paolo VALERI",
			}},
			Name: []fifa.LocaleDescription{{
				Locale:      "en-GB",
				Description: "Paolo Valeri",
			}},
			Type: 8,
			TypeLocalized: []fifa.LocaleDescription{{
				Locale:      "en-GB",
				Description: "Assistant VAR",
			}},
		},
		{
			CountryId: "FRA",
			Id:        "373782",
			ShortName: []fifa.LocaleDescription{{
				Locale:      "en-GB",
				Description: "Benoit MILLOT",
			}},
			Name: []fifa.LocaleDescription{{
				Locale:      "en-GB",
				Description: "Benoit MILLOT",
			}},
			Type: 9,
			TypeLocalized: []fifa.LocaleDescription{{
				Locale:      "en-GB",
				Description: "Support VAR",
			}},
		},
		{
			CountryId: "ROU",
			Id:        "361587",
			ShortName: []fifa.LocaleDescription{{
				Locale:      "en-GB",
				Description: "Mihai ARTENE",
			}},
			Name: []fifa.LocaleDescription{{
				Locale:      "en-GB",
				Description: "Mihai Ovidiu Artene",
			}},
			Type: 10,
			TypeLocalized: []fifa.LocaleDescription{{
				Locale:      "en-GB",
				Description: "Reserve Assistant Referee",
			}},
		},
	},
	MatchStatus:       0,
	ResultType:        1,
	MatchNumber:       1,
	TimeDefined:       true,
	OfficialityStatus: 1,
	Properties: fifa.Properties{
		IdIFES: "128084",
	},
}

func TestGetMatches(t *testing.T) {
	t.Parallel()
	client := fifa.Client{
		Client: &TestClient{},
	}
	resp, err := client.GetMatches(&fifa.GetMatchesOptions{
		CompetitionId: "17",     // World cup
		SeasonId:      "255711", // World Cup Qatar 2022
		Count:         1,
	})
	if ok := assert.Nil(t, err, "expected no error, got: %s", err); !ok {
		t.FailNow()
	}
	diff := cmp.Diff(resp[0], expMatches)
	if ok := assert.Equal(t, diff, "", diff); !ok {
		t.Fail()
	}
}

// This tests against a real match to make sure our marshalling is correct
func TestLiveMatch(t *testing.T) {
	t.Parallel()
	client := fifa.Client{}
	_, err := client.GetMatches(&fifa.GetMatchesOptions{
		CompetitionId: "17",
		SeasonId:      "255711",
		StageId:       "285063",
		From:          time.Date(2022, 11, 28, 0, 0, 0, 0, time.UTC),
		To:            time.Date(2022, 11, 29, 0, 0, 0, 0, time.UTC),
	})
	if ok := assert.Nil(t, err, "expected no error, got: %s", err); !ok {
		t.FailNow()
	}
}

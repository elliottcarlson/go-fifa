package go_fifa_test

import (
	"io"
	"net/http"
	"strings"
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

type TestClient struct{}

func (c *TestClient) Do(req *http.Request) (*http.Response, error) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Status:     http.StatusText(http.StatusOK),
		Body:       io.NopCloser(strings.NewReader(expectedLiveResponse)),
	}
	return resp, nil
}

var _ fifa.HTTPClient = (*TestClient)(nil)

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

	if ok := assert.Len(t, resp, 3, "expected 3 results, got %d", len(resp)); !ok {
		t.Error()
	}

	match := resp[0]

	if ok := assert.Equal(t, match.HomeTeam.Name[0].Description, "RSC Anderlecht II", "expected RSC Anderlecht II but got %s", match.HomeTeam.Name[0].Description); !ok {
		t.Error()
	}

	if ok := assert.Equal(t, match.AwayTeam.Name[0].Description, "Beerschot", "expected Beerschot but got %s", match.AwayTeam.Name[0].Description); !ok {
		t.Error()
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

const expectedLiveResponse = `
{
    "ContinuationToken": null,
    "ContinuationHash": null,
    "Results": [
        {
            "IdMatch": "bn0ffkqrvppd3jahiflnzs1lg",
            "IdStage": "9d5x17v8h81sp5r6q4hbx5hqs",
            "IdGroup": null,
            "IdSeason": "9cy0d6sf8dui1o0rpwynx3190",
            "IdCompetition": "af79lqrc0ntom74zq13ccjslo",
            "CompetitionName": [
                {
                    "Locale": "en-gb",
                    "Description": "Challenger Pro League"
                }
            ],
            "SeasonName": [
                {
                    "Locale": "en-gb",
                    "Description": "Challenger Pro League 2022/2023"
                }
            ],
            "SeasonShortName": [],
            "Stadium": {
                "IdStadium": null,
                "Name": [
                    {
                        "Locale": "en-gb",
                        "Description": "Lotto Park"
                    }
                ],
                "Capacity": 28063,
                "WebAddress": null,
                "Built": "1917-01-01T00:00:00Z",
                "Roof": false,
                "Turf": null,
                "IdCity": "400241509",
                "CityName": [
                    {
                        "Locale": "en-GB",
                        "Description": "Bruxelles (Brussel)"
                    }
                ],
                "IdCountry": "BEL",
                "PostalCode": null,
                "Street": "Avenue Théo Verbeeck 2",
                "Email": null,
                "Fax": null,
                "Phone": null,
                "AffiliationCountry": null,
                "AffiliationRegion": null,
                "Latitude": 50.834182156,
                "Longitude": 4.298352599,
                "Length": null,
                "Width": null,
                "Properties": {},
                "IsUpdateable": null
            },
            "ResultType": 1,
            "MatchDay": "15",
            "HomeTeamPenaltyScore": null,
            "AwayTeamPenaltyScore": null,
            "AggregateHomeTeamScore": null,
            "AggregateAwayTeamScore": null,
            "Weather": null,
            "Attendance": null,
            "Date": "2022-11-26T19:00:00Z",
            "LocalDate": "2022-11-26T20:00:00Z",
            "MatchTime": "66'",
            "SecondHalfTime": null,
            "FirstHalfTime": null,
            "FirstHalfExtraTime": 0,
            "SecondHalfExtraTime": 0,
            "Winner": null,
            "Period": 5,
            "HomeTeam": {
                "Score": 1,
                "Side": null,
                "IdTeam": "d011rbp4etj5styzfmpd7wa38",
                "PictureUrl": "https://api.fifa.com/api/v3/picture/teams-{format}-{size}/d011rbp4etj5styzfmpd7wa38",
                "IdCountry": "BEL",
                "TeamType": 0,
                "AgeType": 7,
                "Tactics": "5-3-2",
                "TeamName": [
                    {
                        "Locale": "en-gb",
                        "Description": "RSC Anderlecht II"
                    }
                ],
                "Abbreviation": "ANF",
                "Coaches": [
                    {
                        "IdCoach": "4dvbxh16w3txv1jy5e2mat638",
                        "IdCountry": "NED",
                        "PictureUrl": null,
                        "Name": [
                            {
                                "Locale": "en-gb",
                                "Description": "Robin Veldeman"
                            }
                        ],
                        "Alias": [
                            {
                                "Locale": "en-gb",
                                "Description": "R. Veldman"
                            }
                        ],
                        "Role": 0,
                        "SpecialStatus": null
                    },
                    {
                        "IdCoach": "290826",
                        "IdCountry": "BEL",
                        "PictureUrl": null,
                        "Name": [
                            {
                                "Locale": "en-gb",
                                "Description": "Guillaume Gillet"
                            }
                        ],
                        "Alias": [
                            {
                                "Locale": "en-gb",
                                "Description": "G. Gillet"
                            }
                        ],
                        "Role": 0,
                        "SpecialStatus": null
                    }
                ],
                "Players": [
                    {
                        "IdPlayer": "c716rboejsazxhoosf6l0v6dw",
                        "IdTeam": "d011rbp4etj5styzfmpd7wa38",
                        "ShirtNumber": 63,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Timon Vanhoutte"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "T. Vanhoutte"
                            }
                        ],
                        "Position": 0,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "5is9bcj143orulo4biogwxzl6",
                        "IdTeam": "d011rbp4etj5styzfmpd7wa38",
                        "ShirtNumber": 47,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Lucas Lissens"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "L. Lissens"
                            }
                        ],
                        "Position": 1,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "e3mwjkfs86jtko5qiqecqy7tg",
                        "IdTeam": "d011rbp4etj5styzfmpd7wa38",
                        "ShirtNumber": 58,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Ethan Butera"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "E. Butera"
                            }
                        ],
                        "Position": 1,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "8lt54ei6pcu25u8uerigmx5yi",
                        "IdTeam": "d011rbp4etj5styzfmpd7wa38",
                        "ShirtNumber": 65,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Alonzo Engwanda"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "A. Engwanda"
                            }
                        ],
                        "Position": 1,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "1e7snjsrbh93xt5ti1oyuqxka",
                        "IdTeam": "d011rbp4etj5styzfmpd7wa38",
                        "ShirtNumber": 46,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Anouar Ait El Hadj"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "A. Ait El Hadj"
                            }
                        ],
                        "Position": 2,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "anx0hiu9yn8szbjk3994u7njd",
                        "IdTeam": "d011rbp4etj5styzfmpd7wa38",
                        "ShirtNumber": 71,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": true,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Theo Leoni"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "T. Leoni"
                            }
                        ],
                        "Position": 2,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "358110",
                        "IdTeam": "d011rbp4etj5styzfmpd7wa38",
                        "ShirtNumber": 88,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "David Hubert"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "D. Hubert"
                            }
                        ],
                        "Position": 2,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "71euf80lf18fu40b9lika9ysq",
                        "IdTeam": "d011rbp4etj5styzfmpd7wa38",
                        "ShirtNumber": 60,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Nils De Wilde"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "N. De Wilde"
                            }
                        ],
                        "Position": 3,
                        "PlayerPicture": null,
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "rao8xq2rg3beoc418zv1gzro",
                        "IdTeam": "d011rbp4etj5styzfmpd7wa38",
                        "ShirtNumber": 76,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Lucas Stassin"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "L. Stassin"
                            }
                        ],
                        "Position": 3,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "a0hdaimb60ls8qj00c6q9las",
                        "IdTeam": "d011rbp4etj5styzfmpd7wa38",
                        "ShirtNumber": 57,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Ilay Camara"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "I. Camara"
                            }
                        ],
                        "Position": 5,
                        "PlayerPicture": null,
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "9vjnul6ac65jk6muwmbil16vo",
                        "IdTeam": "d011rbp4etj5styzfmpd7wa38",
                        "ShirtNumber": 77,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Luca Ferrara"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "L. Ferrara"
                            }
                        ],
                        "Position": 5,
                        "PlayerPicture": null,
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "2zo1hevzos5iwq9vzbzpxz3e1",
                        "IdTeam": "d011rbp4etj5styzfmpd7wa38",
                        "ShirtNumber": 33,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Rik Vercauteren"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "R. Vercauteren"
                            }
                        ],
                        "Position": 0,
                        "PlayerPicture": null,
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "6juck9s3m79lonj4ajb30ik4q",
                        "IdTeam": "d011rbp4etj5styzfmpd7wa38",
                        "ShirtNumber": 64,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Loic Masscho"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "L. Masscho"
                            }
                        ],
                        "Position": 1,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "5auzs4fw49qr80b89baetcwt0",
                        "IdTeam": "d011rbp4etj5styzfmpd7wa38",
                        "ShirtNumber": 67,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Chris Lokesa"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "C. Lokesa"
                            }
                        ],
                        "Position": 2,
                        "PlayerPicture": null,
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "do8nqlhvf43wmla5rakw1chsk",
                        "IdTeam": "d011rbp4etj5styzfmpd7wa38",
                        "ShirtNumber": 73,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Amando Lapage"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "A. Lapage"
                            }
                        ],
                        "Position": 2,
                        "PlayerPicture": null,
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "23snplvy5445j7qepgluw9ppm",
                        "IdTeam": "d011rbp4etj5styzfmpd7wa38",
                        "ShirtNumber": 22,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Ilias Takidine"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "I. Takidine"
                            }
                        ],
                        "Position": 3,
                        "PlayerPicture": null,
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "6es2qusg5h5q9ricqm51j1xck",
                        "IdTeam": "d011rbp4etj5styzfmpd7wa38",
                        "ShirtNumber": 72,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Enock Agyei"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "E. Agyei"
                            }
                        ],
                        "Position": 3,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "dzl3nhd1wmcttn6qkm0sqb7ro",
                        "IdTeam": "d011rbp4etj5styzfmpd7wa38",
                        "ShirtNumber": 75,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Simion Michez"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "S. Michez"
                            }
                        ],
                        "Position": 3,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    }
                ],
                "Bookings": [],
                "Goals": [
                    {
                        "Type": 2,
                        "IdPlayer": "rao8xq2rg3beoc418zv1gzro",
                        "Minute": "11:35",
                        "IdAssistPlayer": "358110",
                        "Period": 3,
                        "IdGoal": null,
                        "IdTeam": "d011rbp4etj5styzfmpd7wa38"
                    }
                ],
                "Substitutions": [
                    {
                        "IdEvent": null,
                        "Period": 5,
                        "Reason": 2,
                        "SubstitutePosition": 3,
                        "IdPlayerOff": "71euf80lf18fu40b9lika9ysq",
                        "IdPlayerOn": "dzl3nhd1wmcttn6qkm0sqb7ro",
                        "PlayerOffName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Nils De Wilde"
                            }
                        ],
                        "PlayerOnName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Simion Michez"
                            }
                        ],
                        "Minute": "64:52",
                        "IdTeam": "d011rbp4etj5styzfmpd7wa38"
                    },
                    {
                        "IdEvent": null,
                        "Period": 5,
                        "Reason": 2,
                        "SubstitutePosition": 3,
                        "IdPlayerOff": "a0hdaimb60ls8qj00c6q9las",
                        "IdPlayerOn": "6es2qusg5h5q9ricqm51j1xck",
                        "PlayerOffName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Ilay Camara"
                            }
                        ],
                        "PlayerOnName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Enock Agyei"
                            }
                        ],
                        "Minute": "65:31",
                        "IdTeam": "d011rbp4etj5styzfmpd7wa38"
                    },
                    {
                        "IdEvent": null,
                        "Period": 5,
                        "Reason": 2,
                        "SubstitutePosition": 1,
                        "IdPlayerOff": "9vjnul6ac65jk6muwmbil16vo",
                        "IdPlayerOn": "6juck9s3m79lonj4ajb30ik4q",
                        "PlayerOffName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Luca Ferrara"
                            }
                        ],
                        "PlayerOnName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Loic Masscho"
                            }
                        ],
                        "Minute": "65:50",
                        "IdTeam": "d011rbp4etj5styzfmpd7wa38"
                    }
                ],
                "FootballType": 0,
                "Gender": 1,
                "IdAssociation": "BEL",
                "ShortClubName": "Anderlecht II"
            },
            "AwayTeam": {
                "Score": 2,
                "Side": null,
                "IdTeam": "2000019799",
                "PictureUrl": "https://api.fifa.com/api/v3/picture/teams-{format}-{size}/2000019799",
                "IdCountry": "BEL",
                "TeamType": 0,
                "AgeType": 7,
                "Tactics": "3-4-3",
                "TeamName": [
                    {
                        "Locale": "en-gb",
                        "Description": "Beerschot"
                    }
                ],
                "Abbreviation": "BEE",
                "Coaches": [
                    {
                        "IdCoach": "co3yqoz2wqu4bx8jh2t8gjrl6",
                        "IdCountry": "AUT",
                        "PictureUrl": null,
                        "Name": [
                            {
                                "Locale": "en-gb",
                                "Description": "Andreas Wieland"
                            }
                        ],
                        "Alias": [
                            {
                                "Locale": "en-gb",
                                "Description": "A. Wieland"
                            }
                        ],
                        "Role": 0,
                        "SpecialStatus": null
                    }
                ],
                "Players": [
                    {
                        "IdPlayer": "8r3r6z1khjq55pweqzfqezpm2",
                        "IdTeam": "2000019799",
                        "ShirtNumber": 1,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Bill Lathouwers"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "B. Lathouwers"
                            }
                        ],
                        "Position": 0,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "b1h4yvtkjhvtkcgu3gtgiytx",
                        "IdTeam": "2000019799",
                        "ShirtNumber": 2,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Jan Van den Bergh"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "J. Van den Bergh"
                            }
                        ],
                        "Position": 1,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "d6eq2hlmwh2hpbx4m3qjiuxud",
                        "IdTeam": "2000019799",
                        "ShirtNumber": 3,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Hervé Matthys"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "H. Matthys"
                            }
                        ],
                        "Position": 1,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "2ky2kn7gpjorkyg9zyg68pk6i",
                        "IdTeam": "2000019799",
                        "ShirtNumber": 66,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Apostolos Konstantopoulos"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "A. Konstantopoulos"
                            }
                        ],
                        "Position": 1,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "7q0v538ed6j5xy3wgykh56w2x",
                        "IdTeam": "2000019799",
                        "ShirtNumber": 16,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Léo Seydoux"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "L. Seydoux"
                            }
                        ],
                        "Position": 2,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "cfxzvlgvt9jj17qxq41t9sxcl",
                        "IdTeam": "2000019799",
                        "ShirtNumber": 18,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": true,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Ryan Sanusi"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "R. Sanusi"
                            }
                        ],
                        "Position": 2,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "389572",
                        "IdTeam": "2000019799",
                        "ShirtNumber": 35,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Dante Rigo"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "RIGO"
                            }
                        ],
                        "Position": 2,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "aeue0dvxs088u4ucahccoe1as",
                        "IdTeam": "2000019799",
                        "ShirtNumber": 40,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Ilias Sebaoui"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "I. Sebaoui"
                            }
                        ],
                        "Position": 2,
                        "PlayerPicture": null,
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "b157xbo2jovsxkgr3wie3rip6",
                        "IdTeam": "2000019799",
                        "ShirtNumber": 7,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Thibo Baeten"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "T. Baeten"
                            }
                        ],
                        "Position": 3,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "6g9x1y7xmjzwhk646v1iyuup5",
                        "IdTeam": "2000019799",
                        "ShirtNumber": 10,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Thibaud Verlinden"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "T. Verlinden"
                            }
                        ],
                        "Position": 3,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "380782",
                        "IdTeam": "2000019799",
                        "ShirtNumber": 17,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Nökkvi Þeyr Þórisson"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "N. Þórisson"
                            }
                        ],
                        "Position": 3,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "7qkoogt2mpaqmmwyjgdzqe6zt",
                        "IdTeam": "2000019799",
                        "ShirtNumber": 71,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Davor Matijas"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "D. Matijaš"
                            }
                        ],
                        "Position": 0,
                        "PlayerPicture": null,
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "cjpcrd2e3exb4k1izyfot4y22",
                        "IdTeam": "2000019799",
                        "ShirtNumber": 5,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Robbe Quirynen"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "R. Quirynen"
                            }
                        ],
                        "Position": 1,
                        "PlayerPicture": null,
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "axz9twlt1z4ghq0ovf58vgv6x",
                        "IdTeam": "2000019799",
                        "ShirtNumber": 6,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Luca Meisl"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "L. Meisl"
                            }
                        ],
                        "Position": 1,
                        "PlayerPicture": null,
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "73g7emt47agukey5jmc78zaol",
                        "IdTeam": "2000019799",
                        "ShirtNumber": 28,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Marco Weymans"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "M. Weymans"
                            }
                        ],
                        "Position": 1,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "6w45ga1ltkcn72xtv9ps5um7d",
                        "IdTeam": "2000019799",
                        "ShirtNumber": 8,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Ibrahim Alhassan"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "A. Alhassan"
                            }
                        ],
                        "Position": 2,
                        "PlayerPicture": null,
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "404254",
                        "IdTeam": "2000019799",
                        "ShirtNumber": 11,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Ramiro VACA"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "R. Vaca"
                            }
                        ],
                        "Position": 2,
                        "PlayerPicture": null,
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "1ng8mptm1t32jirhkhfq2wyui",
                        "IdTeam": "2000019799",
                        "ShirtNumber": 22,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Andi Koshi"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "A. Koshi"
                            }
                        ],
                        "Position": 2,
                        "PlayerPicture": null,
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    }
                ],
                "Bookings": [
                    {
                        "Card": 1,
                        "Period": 3,
                        "IdEvent": null,
                        "EventNumber": null,
                        "IdPlayer": "2ky2kn7gpjorkyg9zyg68pk6i",
                        "IdCoach": null,
                        "IdTeam": "2000019799",
                        "Minute": "38:56",
                        "Reason": "OffTheBallFoul"
                    }
                ],
                "Goals": [
                    {
                        "Type": 2,
                        "IdPlayer": "cfxzvlgvt9jj17qxq41t9sxcl",
                        "Minute": "33:28",
                        "IdAssistPlayer": "7q0v538ed6j5xy3wgykh56w2x",
                        "Period": 3,
                        "IdGoal": null,
                        "IdTeam": "2000019799"
                    },
                    {
                        "Type": 2,
                        "IdPlayer": "380782",
                        "Minute": "56:37",
                        "IdAssistPlayer": "7q0v538ed6j5xy3wgykh56w2x",
                        "Period": 5,
                        "IdGoal": null,
                        "IdTeam": "2000019799"
                    }
                ],
                "Substitutions": [
                    {
                        "IdEvent": null,
                        "Period": 5,
                        "Reason": 2,
                        "SubstitutePosition": 1,
                        "IdPlayerOff": "aeue0dvxs088u4ucahccoe1as",
                        "IdPlayerOn": "73g7emt47agukey5jmc78zaol",
                        "PlayerOffName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Ilias Sebaoui"
                            }
                        ],
                        "PlayerOnName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Marco Weymans"
                            }
                        ],
                        "Minute": "45:00",
                        "IdTeam": "2000019799"
                    }
                ],
                "FootballType": 0,
                "Gender": 1,
                "IdAssociation": "BEL",
                "ShortClubName": "Beerschot"
            },
            "BallPossession": {
                "Intervals": [],
                "LastX": [],
                "OverallHome": 54.2,
                "OverallAway": 45.8
            },
            "TerritorialPossesion": null,
            "TerritorialThirdPossesion": null,
            "Officials": [
                {
                    "IdCountry": "BEL",
                    "OfficialId": "4sabn8pw0tbyrrnwcch72o05x",
                    "NameShort": [
                        {
                            "Locale": "en-gb",
                            "Description": "Q. Pirard"
                        }
                    ],
                    "Name": [
                        {
                            "Locale": "en-gb",
                            "Description": "Quentin Pirard"
                        }
                    ],
                    "OfficialType": 1,
                    "TypeLocalized": [
                        {
                            "Locale": "en-GB",
                            "Description": "Referee"
                        }
                    ]
                },
                {
                    "IdCountry": "BEL",
                    "OfficialId": "2kzam9ap51k4annm3s50z7dd6",
                    "NameShort": [
                        {
                            "Locale": "en-gb",
                            "Description": "J. Van den Borre"
                        }
                    ],
                    "Name": [
                        {
                            "Locale": "en-gb",
                            "Description": "Jordan Van Dden Borre"
                        }
                    ],
                    "OfficialType": 2,
                    "TypeLocalized": [
                        {
                            "Locale": "en-GB",
                            "Description": "Assistant Referee 1"
                        }
                    ]
                },
                {
                    "IdCountry": "NED",
                    "OfficialId": "b86gr42gds0le0tg5d6r10b9w",
                    "NameShort": [
                        {
                            "Locale": "en-gb",
                            "Description": "J. van Dyck"
                        }
                    ],
                    "Name": [
                        {
                            "Locale": "en-gb",
                            "Description": "Jonas van Dyck"
                        }
                    ],
                    "OfficialType": 3,
                    "TypeLocalized": [
                        {
                            "Locale": "en-GB",
                            "Description": "Assistant Referee 2"
                        }
                    ]
                },
                {
                    "IdCountry": "BEL",
                    "OfficialId": "428449",
                    "NameShort": [
                        {
                            "Locale": "en-gb",
                            "Description": "J. Vermeire"
                        }
                    ],
                    "Name": [
                        {
                            "Locale": "en-GB",
                            "Description": "Jordy Vermeire"
                        }
                    ],
                    "OfficialType": 4,
                    "TypeLocalized": [
                        {
                            "Locale": "en-GB",
                            "Description": "Fourth official"
                        }
                    ]
                }
            ],
            "MatchStatus": 3,
            "GroupName": [],
            "StageName": [
                {
                    "Locale": "en-gb",
                    "Description": "Regular Season"
                }
            ],
            "OfficialityStatus": 0,
            "TimeDefined": true,
            "Properties": {
                "IdStatsPerform": "bn0ffkqrvppd3jahiflnzs1lg"
            },
            "IsUpdateable": null
        },
        {
            "IdMatch": "400235462",
            "IdStage": "285063",
            "IdGroup": "285067",
            "IdSeason": "255711",
            "IdCompetition": "17",
            "CompetitionName": [
                {
                    "Locale": "en-GB",
                    "Description": "FIFA World Cup™"
                }
            ],
            "SeasonName": [
                {
                    "Locale": "en-GB",
                    "Description": "FIFA World Cup Qatar 2022™"
                }
            ],
            "SeasonShortName": [],
            "Stadium": {
                "IdStadium": "400090323",
                "Name": [
                    {
                        "Locale": "en-GB",
                        "Description": "Lusail Stadium"
                    }
                ],
                "Capacity": null,
                "WebAddress": null,
                "Built": null,
                "Roof": false,
                "Turf": null,
                "IdCity": "400157371",
                "CityName": [
                    {
                        "Locale": "en-GB",
                        "Description": "Al Daayen"
                    }
                ],
                "IdCountry": "QAT",
                "PostalCode": null,
                "Street": null,
                "Email": null,
                "Fax": null,
                "Phone": null,
                "AffiliationCountry": null,
                "AffiliationRegion": null,
                "Latitude": null,
                "Longitude": null,
                "Length": null,
                "Width": null,
                "Properties": {
                    "IdIFES": "5047371"
                },
                "IsUpdateable": null
            },
            "ResultType": 0,
            "MatchDay": "2",
            "HomeTeamPenaltyScore": 0,
            "AwayTeamPenaltyScore": 0,
            "AggregateHomeTeamScore": null,
            "AggregateAwayTeamScore": null,
            "Weather": {
                "Humidity": "71",
                "Temperature": "24",
                "WindSpeed": "1",
                "Type": 10,
                "TypeLocalized": [
                    {
                        "Locale": "en-GB",
                        "Description": "Clear Night"
                    }
                ]
            },
            "Attendance": "88966",
            "Date": "2022-11-26T19:00:00Z",
            "LocalDate": "2022-11-26T22:00:00Z",
            "MatchTime": "64'",
            "SecondHalfTime": null,
            "FirstHalfTime": null,
            "FirstHalfExtraTime": null,
            "SecondHalfExtraTime": null,
            "Winner": null,
            "Period": 5,
            "HomeTeam": {
                "Score": 1,
                "Side": null,
                "IdTeam": "43922",
                "PictureUrl": "https://api.fifa.com/api/v3/picture/flags-{format}-{size}/ARG",
                "IdCountry": "ARG",
                "TeamType": 1,
                "AgeType": 7,
                "Tactics": "4-4-2",
                "TeamName": [
                    {
                        "Locale": "en-GB",
                        "Description": "Argentina"
                    }
                ],
                "Abbreviation": "ARG",
                "Coaches": [
                    {
                        "IdCoach": "153933",
                        "IdCountry": "ARG",
                        "PictureUrl": null,
                        "Name": [
                            {
                                "Locale": "en-GB",
                                "Description": "Roberto AYALA"
                            }
                        ],
                        "Alias": [
                            {
                                "Locale": "en-GB",
                                "Description": "Roberto AYALA"
                            }
                        ],
                        "Role": 1,
                        "SpecialStatus": null
                    },
                    {
                        "IdCoach": "183380",
                        "IdCountry": "ARG",
                        "PictureUrl": "https://digitalhub.fifa.com/transform/1cd0729f-827e-4547-b7cb-93b186c209a5/1442739826",
                        "Name": [
                            {
                                "Locale": "en-GB",
                                "Description": "Lionel SCALONI"
                            }
                        ],
                        "Alias": [
                            {
                                "Locale": "en-GB",
                                "Description": "Lionel SCALONI"
                            }
                        ],
                        "Role": 0,
                        "SpecialStatus": null
                    },
                    {
                        "IdCoach": "3076",
                        "IdCountry": "ARG",
                        "PictureUrl": null,
                        "Name": [
                            {
                                "Locale": "en-GB",
                                "Description": "Walter SAMUEL"
                            }
                        ],
                        "Alias": [
                            {
                                "Locale": "en-GB",
                                "Description": "Walter SAMUEL"
                            }
                        ],
                        "Role": 1,
                        "SpecialStatus": null
                    },
                    {
                        "IdCoach": "178142",
                        "IdCountry": "ARG",
                        "PictureUrl": null,
                        "Name": [
                            {
                                "Locale": "en-GB",
                                "Description": "Pablo AIMAR"
                            }
                        ],
                        "Alias": [
                            {
                                "Locale": "en-GB",
                                "Description": "Pablo AIMAR"
                            }
                        ],
                        "Role": 1,
                        "SpecialStatus": null
                    },
                    {
                        "IdCoach": "380527",
                        "IdCountry": "ARG",
                        "PictureUrl": null,
                        "Name": [
                            {
                                "Locale": "en-GB",
                                "Description": "Matias MANNA"
                            }
                        ],
                        "Alias": [
                            {
                                "Locale": "en-GB",
                                "Description": "Matias MANNA"
                            }
                        ],
                        "Role": 1,
                        "SpecialStatus": null
                    }
                ],
                "Players": [
                    {
                        "IdPlayer": "308300",
                        "IdTeam": "43922",
                        "ShirtNumber": 23,
                        "Status": 1,
                        "SpecialStatus": 2,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Damian MARTINEZ"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "MARTINEZ D."
                            }
                        ],
                        "Position": 0,
                        "PlayerPicture": {
                            "Id": "EC57583F-85ED-4B6E-B6198EC94DE13B8F",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/ec57583f-85ed-4b6e-b619-8ec94de13b8f/1442740128"
                        },
                        "FieldStatus": 1,
                        "LineupX": 10.0,
                        "LineupY": 1.0
                    },
                    {
                        "IdPlayer": "402926",
                        "IdTeam": "43922",
                        "ShirtNumber": 4,
                        "Status": 1,
                        "SpecialStatus": 2,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Gonzalo MONTIEL"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "MONTIEL"
                            }
                        ],
                        "Position": 1,
                        "PlayerPicture": {
                            "Id": "D37E7B7F-E91D-4ED4-B6018494B56AF199",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/d37e7b7f-e91d-4ed4-b601-8494b56af199/1442742819"
                        },
                        "FieldStatus": 2,
                        "LineupX": 18.0,
                        "LineupY": 4.0
                    },
                    {
                        "IdPlayer": "401204",
                        "IdTeam": "43922",
                        "ShirtNumber": 8,
                        "Status": 1,
                        "SpecialStatus": 2,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Marcos ACUNA"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "ACUÑA"
                            }
                        ],
                        "Position": 1,
                        "PlayerPicture": {
                            "Id": "B1D9EBE0-067F-4246-881DED9DAB03835E",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/b1d9ebe0-067f-4246-881d-ed9dab03835e/1442742836"
                        },
                        "FieldStatus": 1,
                        "LineupX": 2.0,
                        "LineupY": 4.0
                    },
                    {
                        "IdPlayer": "310116",
                        "IdTeam": "43922",
                        "ShirtNumber": 19,
                        "Status": 1,
                        "SpecialStatus": 2,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Nicolas OTAMENDI"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "OTAMENDI"
                            }
                        ],
                        "Position": 1,
                        "PlayerPicture": {
                            "Id": "592D999D-3DED-4A1C-8DF10DDCAD20C18B",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/592d999d-3ded-4a1c-8df1-0ddcad20c18b/1442741360"
                        },
                        "FieldStatus": 1,
                        "LineupX": 14.0,
                        "LineupY": 4.0
                    },
                    {
                        "IdPlayer": "402921",
                        "IdTeam": "43922",
                        "ShirtNumber": 25,
                        "Status": 1,
                        "SpecialStatus": 2,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Lisandro MARTINEZ"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "MARTINEZ"
                            }
                        ],
                        "Position": 1,
                        "PlayerPicture": {
                            "Id": "5C6BC06C-6694-4863-B543E74B672741A3",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/5c6bc06c-6694-4863-b543-e74b672741a3/1442801367"
                        },
                        "FieldStatus": 1,
                        "LineupX": 6.0,
                        "LineupY": 4.0
                    },
                    {
                        "IdPlayer": "428882",
                        "IdTeam": "43922",
                        "ShirtNumber": 7,
                        "Status": 1,
                        "SpecialStatus": 2,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Rodrigo DE PAUL"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "DE PAUL"
                            }
                        ],
                        "Position": 2,
                        "PlayerPicture": {
                            "Id": "427CC23A-3489-42F1-892F881B85B857D8",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/427cc23a-3489-42f1-892f-881b85b857d8/1442742174"
                        },
                        "FieldStatus": 1,
                        "LineupX": 14.0,
                        "LineupY": 6.0
                    },
                    {
                        "IdPlayer": "266800",
                        "IdTeam": "43922",
                        "ShirtNumber": 11,
                        "Status": 1,
                        "SpecialStatus": 2,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Angel DI MARIA"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "DI MARÍA"
                            }
                        ],
                        "Position": 2,
                        "PlayerPicture": {
                            "Id": "1CED2AF1-6B06-4896-BA4B111C9DE737E9",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/1ced2af1-6b06-4896-ba4b-111c9de737e9/1442742357"
                        },
                        "FieldStatus": 1,
                        "LineupX": 18.0,
                        "LineupY": 8.0
                    },
                    {
                        "IdPlayer": "392671",
                        "IdTeam": "43922",
                        "ShirtNumber": 18,
                        "Status": 1,
                        "SpecialStatus": 2,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "RODRIGUEZ GUIDO"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "G. RODRIGUEZ"
                            }
                        ],
                        "Position": 2,
                        "PlayerPicture": {
                            "Id": "7819E5F3-30CA-40B3-82F1E0D30B83E384",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/7819e5f3-30ca-40b3-82f1-e0d30b83e384/1442743421"
                        },
                        "FieldStatus": 2,
                        "LineupX": 6.0,
                        "LineupY": 6.0
                    },
                    {
                        "IdPlayer": "430628",
                        "IdTeam": "43922",
                        "ShirtNumber": 20,
                        "Status": 1,
                        "SpecialStatus": 2,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Alexis MAC ALLISTER"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "MAC ALLISTER"
                            }
                        ],
                        "Position": 2,
                        "PlayerPicture": {
                            "Id": "587F8FF4-72A1-4499-BF71C38BBC87C9EC",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/587f8ff4-72a1-4499-bf71-c38bbc87c9ec/1442742544"
                        },
                        "FieldStatus": 1,
                        "LineupX": 2.0,
                        "LineupY": 8.0
                    },
                    {
                        "IdPlayer": "229397",
                        "IdTeam": "43922",
                        "ShirtNumber": 10,
                        "Status": 1,
                        "SpecialStatus": 3,
                        "Captain": true,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Lionel MESSI"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "MESSI"
                            }
                        ],
                        "Position": 3,
                        "PlayerPicture": {
                            "Id": "40E6D6B5-9742-4123-8FB8D69662C3B24A",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/40e6d6b5-9742-4123-8fb8-d69662c3b24a/1442770580"
                        },
                        "FieldStatus": 1,
                        "LineupX": 14.0,
                        "LineupY": 12.0
                    },
                    {
                        "IdPlayer": "402920",
                        "IdTeam": "43922",
                        "ShirtNumber": 22,
                        "Status": 1,
                        "SpecialStatus": 2,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Lautaro MARTINEZ"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "L. MARTINEZ"
                            }
                        ],
                        "Position": 3,
                        "PlayerPicture": {
                            "Id": "CE6502BD-6A21-4A83-A4349EE5DA63E99E",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/ce6502bd-6a21-4a83-a434-9ee5da63e99e/1442742404"
                        },
                        "FieldStatus": 2,
                        "LineupX": 6.0,
                        "LineupY": 12.0
                    },
                    {
                        "IdPlayer": "398422",
                        "IdTeam": "43922",
                        "ShirtNumber": 1,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Franco ARMANI"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "ARMANI"
                            }
                        ],
                        "Position": 0,
                        "PlayerPicture": {
                            "Id": "55E9FE36-B328-44F5-86D81CBA96A4D05E",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/55e9fe36-b328-44f5-86d8-1cba96a4d05e/1442740056"
                        },
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "394824",
                        "IdTeam": "43922",
                        "ShirtNumber": 12,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Geronimo RULLI"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "RULLI"
                            }
                        ],
                        "Position": 0,
                        "PlayerPicture": {
                            "Id": "08C7E932-41A2-460E-B14B04903AC5CA8C",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/08c7e932-41a2-460e-b14b-04903ac5ca8c/1442774135"
                        },
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "402918",
                        "IdTeam": "43922",
                        "ShirtNumber": 2,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Juan FOYTH"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "FOYTH"
                            }
                        ],
                        "Position": 4,
                        "PlayerPicture": {
                            "Id": "FD88628F-36A5-4BEF-875EFCB52581207A",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/fd88628f-36a5-4bef-875e-fcb52581207a/1442742652"
                        },
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "308322",
                        "IdTeam": "43922",
                        "ShirtNumber": 3,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Nicolas TAGLIAFICO"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "TAGLIAFICO"
                            }
                        ],
                        "Position": 4,
                        "PlayerPicture": {
                            "Id": "9C5788AF-5F9C-4D80-BA26296EDD06E352",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/9c5788af-5f9c-4d80-ba26-296edd06e352/1442740186"
                        },
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "332847",
                        "IdTeam": "43922",
                        "ShirtNumber": 5,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Leandro PAREDES"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "PAREDES"
                            }
                        ],
                        "Position": 4,
                        "PlayerPicture": {
                            "Id": "DF4EB81D-B234-432C-A5D147F10FD6B844",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/df4eb81d-b234-432c-a5d1-47f10fd6b844/1442742223"
                        },
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "332979",
                        "IdTeam": "43922",
                        "ShirtNumber": 6,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "German PEZZELLA"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "PEZZELLA"
                            }
                        ],
                        "Position": 4,
                        "PlayerPicture": {
                            "Id": "57C8D393-5328-40AB-A5F259EFD1A042DF",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/57c8d393-5328-40ab-a5f2-59efd1a042df/1442742197"
                        },
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "416081",
                        "IdTeam": "43922",
                        "ShirtNumber": 9,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Julian ALVAREZ"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "J. ALVAREZ"
                            }
                        ],
                        "Position": 4,
                        "PlayerPicture": {
                            "Id": "86F9D09D-336E-49E9-868F78E8426D6BFB",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/86f9d09d-336e-49e9-868f-78e8426d6bfb/1442743531"
                        },
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "431196",
                        "IdTeam": "43922",
                        "ShirtNumber": 13,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Cristian ROMERO"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "ROMERO"
                            }
                        ],
                        "Position": 4,
                        "PlayerPicture": {
                            "Id": "10DB8803-6D59-4A4B-909D4CB0B297F904",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/10db8803-6d59-4a4b-909d-4cb0b297f904/1442742592"
                        },
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "389485",
                        "IdTeam": "43922",
                        "ShirtNumber": 14,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Exequiel PALACIOS"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "PALACIOS"
                            }
                        ],
                        "Position": 4,
                        "PlayerPicture": {
                            "Id": "79DB6A3D-02A5-470E-9B367828D5B2D02C",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/79db6a3d-02a5-470e-9b36-7828d5b2d02c/1442742805"
                        },
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "385343",
                        "IdTeam": "43922",
                        "ShirtNumber": 15,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Angel CORREA"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "A. CORREA"
                            }
                        ],
                        "Position": 4,
                        "PlayerPicture": {
                            "Id": "1C0D0DD2-244D-4566-972DF6BF0ACD3632",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/1c0d0dd2-244d-4566-972d-f6bf0acd3632/1442770400"
                        },
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "418975",
                        "IdTeam": "43922",
                        "ShirtNumber": 16,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Thiago ALMADA"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Thiago Almada"
                            }
                        ],
                        "Position": 4,
                        "PlayerPicture": {
                            "Id": "555F0C07-6203-4AAF-88306FD909A22193",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/555f0c07-6203-4aaf-8830-6fd909a22193/1442770365"
                        },
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "228537",
                        "IdTeam": "43922",
                        "ShirtNumber": 17,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Alejandro GOMEZ"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "GOMEZ"
                            }
                        ],
                        "Position": 4,
                        "PlayerPicture": {
                            "Id": "A2A9F2D8-B6EF-4152-87EA82314C0C106E",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/a2a9f2d8-b6ef-4152-87ea-82314c0c106e/1442741064"
                        },
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "392905",
                        "IdTeam": "43922",
                        "ShirtNumber": 21,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Paulo DYBALA"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "DYBALA"
                            }
                        ],
                        "Position": 4,
                        "PlayerPicture": {
                            "Id": "CA25A965-D216-44DC-B256CC708605D2FC",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/ca25a965-d216-44dc-b256-cc708605d2fc/1442740796"
                        },
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "448252",
                        "IdTeam": "43922",
                        "ShirtNumber": 24,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Enzo FERNANDEZ"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "FERNÁNDEZ"
                            }
                        ],
                        "Position": 4,
                        "PlayerPicture": {
                            "Id": "15281E95-C206-4287-A572C52D523E188C",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/15281e95-c206-4287-a572-c52d523e188c/1442743507"
                        },
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "402925",
                        "IdTeam": "43922",
                        "ShirtNumber": 26,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Nahuel MOLINA"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "MOLINA"
                            }
                        ],
                        "Position": 4,
                        "PlayerPicture": {
                            "Id": "0C723EC4-6132-42DE-8B21A70AD7A1C4C1",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/0c723ec4-6132-42de-8b21-a70ad7a1c4c1/1442742632"
                        },
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    }
                ],
                "Bookings": [
                    {
                        "Card": 1,
                        "Period": 3,
                        "IdEvent": null,
                        "EventNumber": null,
                        "IdPlayer": "402926",
                        "IdCoach": null,
                        "IdTeam": "43922",
                        "Minute": "43'",
                        "Reason": null
                    }
                ],
                "Goals": [
                    {
                        "Type": 2,
                        "IdPlayer": "229397",
                        "Minute": "64'",
                        "IdAssistPlayer": "215285",
                        "Period": 5,
                        "IdGoal": null,
                        "IdTeam": "43922"
                    }
                ],
                "Substitutions": [
                    {
                        "IdEvent": null,
                        "Period": 5,
                        "Reason": 0,
                        "SubstitutePosition": 4,
                        "IdPlayerOff": "392671",
                        "IdPlayerOn": "448252",
                        "PlayerOffName": [
                            {
                                "Locale": "en-GB",
                                "Description": "RODRIGUEZ GUIDO"
                            }
                        ],
                        "PlayerOnName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Enzo FERNANDEZ"
                            }
                        ],
                        "Minute": "57'",
                        "IdTeam": "43922"
                    },
                    {
                        "IdEvent": null,
                        "Period": 5,
                        "Reason": 0,
                        "SubstitutePosition": 4,
                        "IdPlayerOff": "402920",
                        "IdPlayerOn": "416081",
                        "PlayerOffName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Lautaro MARTINEZ"
                            }
                        ],
                        "PlayerOnName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Julian ALVAREZ"
                            }
                        ],
                        "Minute": "63'",
                        "IdTeam": "43922"
                    },
                    {
                        "IdEvent": null,
                        "Period": 5,
                        "Reason": 0,
                        "SubstitutePosition": 4,
                        "IdPlayerOff": "402926",
                        "IdPlayerOn": "402925",
                        "PlayerOffName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Gonzalo MONTIEL"
                            }
                        ],
                        "PlayerOnName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Nahuel MOLINA"
                            }
                        ],
                        "Minute": "63'",
                        "IdTeam": "43922"
                    }
                ],
                "FootballType": 0,
                "Gender": 1,
                "IdAssociation": "ARG",
                "ShortClubName": "Argentina"
            },
            "AwayTeam": {
                "Score": 0,
                "Side": null,
                "IdTeam": "43911",
                "PictureUrl": "https://api.fifa.com/api/v3/picture/flags-{format}-{size}/MEX",
                "IdCountry": "MEX",
                "TeamType": 1,
                "AgeType": 7,
                "Tactics": "5-3-2",
                "TeamName": [
                    {
                        "Locale": "en-GB",
                        "Description": "Mexico"
                    }
                ],
                "Abbreviation": "MEX",
                "Coaches": [
                    {
                        "IdCoach": "192445",
                        "IdCountry": "ARG",
                        "PictureUrl": "https://digitalhub.fifa.com/transform/3ac968b4-1742-4773-ac0e-c07c55859981/1442569708",
                        "Name": [
                            {
                                "Locale": "en-GB",
                                "Description": "Gerardo Daniel Martino"
                            }
                        ],
                        "Alias": [
                            {
                                "Locale": "en-GB",
                                "Description": "Gerardo MARTINO"
                            }
                        ],
                        "Role": 0,
                        "SpecialStatus": null
                    },
                    {
                        "IdCoach": "192443",
                        "IdCountry": "ITA",
                        "PictureUrl": null,
                        "Name": [
                            {
                                "Locale": "en-GB",
                                "Description": "Sergio GIOVAGNOLI"
                            }
                        ],
                        "Alias": [
                            {
                                "Locale": "en-GB",
                                "Description": "Sergio GIOVAGNOLI"
                            }
                        ],
                        "Role": 1,
                        "SpecialStatus": null
                    },
                    {
                        "IdCoach": "192906",
                        "IdCountry": "ARG",
                        "PictureUrl": null,
                        "Name": [
                            {
                                "Locale": "en-GB",
                                "Description": "Jorge Walter Theiler"
                            }
                        ],
                        "Alias": [
                            {
                                "Locale": "en-GB",
                                "Description": "Jorge THEILER"
                            }
                        ],
                        "Role": 1,
                        "SpecialStatus": null
                    },
                    {
                        "IdCoach": "448032",
                        "IdCountry": "ARG",
                        "PictureUrl": null,
                        "Name": [
                            {
                                "Locale": "en-GB",
                                "Description": "Damian SILVERO"
                            }
                        ],
                        "Alias": [
                            {
                                "Locale": "en-GB",
                                "Description": "Damian SILVERO"
                            }
                        ],
                        "Role": 1,
                        "SpecialStatus": null
                    }
                ],
                "Players": [
                    {
                        "IdPlayer": "215285",
                        "IdTeam": "43911",
                        "ShirtNumber": 13,
                        "Status": 1,
                        "SpecialStatus": 2,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Guillermo OCHOA"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "G. OCHOA"
                            }
                        ],
                        "Position": 0,
                        "PlayerPicture": {
                            "Id": "C334C4EA-5BA3-4C34-AB21219A5F9A304D",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/c334c4ea-5ba3-4c34-ab21-219a5f9a304d/1442569806"
                        },
                        "FieldStatus": 1,
                        "LineupX": 10.0,
                        "LineupY": 1.0
                    },
                    {
                        "IdPlayer": "336707",
                        "IdTeam": "43911",
                        "ShirtNumber": 2,
                        "Status": 1,
                        "SpecialStatus": 2,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Nestor ARAUJO"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "N. ARAUJO"
                            }
                        ],
                        "Position": 1,
                        "PlayerPicture": {
                            "Id": "46DB0D50-3AA7-4935-BB0B369960534DD6",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/46db0d50-3aa7-4935-bb0b-369960534dd6/1442571186"
                        },
                        "FieldStatus": 1,
                        "LineupX": 14.0,
                        "LineupY": 4.0
                    },
                    {
                        "IdPlayer": "395516",
                        "IdTeam": "43911",
                        "ShirtNumber": 3,
                        "Status": 1,
                        "SpecialStatus": 2,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Cesar  MONTES"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "C. MONTES"
                            }
                        ],
                        "Position": 1,
                        "PlayerPicture": {
                            "Id": "996B1BEC-EB20-4225-8F1DA5FA2EEAEBAB",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/996b1bec-eb20-4225-8f1d-a5fa2eeaebab/1442569877"
                        },
                        "FieldStatus": 1,
                        "LineupX": 10.0,
                        "LineupY": 4.0
                    },
                    {
                        "IdPlayer": "238112",
                        "IdTeam": "43911",
                        "ShirtNumber": 15,
                        "Status": 1,
                        "SpecialStatus": 6,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Hector MORENO"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "H. MORENO"
                            }
                        ],
                        "Position": 1,
                        "PlayerPicture": {
                            "Id": "96911B40-3A64-4831-8B291DB53CE5EE8B",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/96911b40-3a64-4831-8b29-1db53ce5ee8b/1442570919"
                        },
                        "FieldStatus": 1,
                        "LineupX": 6.0,
                        "LineupY": 4.0
                    },
                    {
                        "IdPlayer": "402772",
                        "IdTeam": "43911",
                        "ShirtNumber": 23,
                        "Status": 1,
                        "SpecialStatus": 2,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Jesus GALLARDO"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "J. GALLARDO"
                            }
                        ],
                        "Position": 1,
                        "PlayerPicture": {
                            "Id": "DC8D4F29-9ADB-4CDD-B44F841B09965003",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/dc8d4f29-9adb-4cdd-b44f-841b09965003/1442569931"
                        },
                        "FieldStatus": 1,
                        "LineupX": 2.0,
                        "LineupY": 4.0
                    },
                    {
                        "IdPlayer": "419506",
                        "IdTeam": "43911",
                        "ShirtNumber": 26,
                        "Status": 1,
                        "SpecialStatus": 2,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Kevin ALVAREZ"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "K. ÁLVAREZ"
                            }
                        ],
                        "Position": 1,
                        "PlayerPicture": {
                            "Id": "EC064316-4789-4EC2-8FA5348C977E32D3",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/ec064316-4789-4ec2-8fa5-348c977e32d3/1442571236"
                        },
                        "FieldStatus": 1,
                        "LineupX": 18.0,
                        "LineupY": 4.0
                    },
                    {
                        "IdPlayer": "329092",
                        "IdTeam": "43911",
                        "ShirtNumber": 16,
                        "Status": 1,
                        "SpecialStatus": 2,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Hector HERRERA"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "H. HERRERA"
                            }
                        ],
                        "Position": 2,
                        "PlayerPicture": {
                            "Id": "CB8A2537-555D-46AF-BC004C9A24955A0D",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/cb8a2537-555d-46af-bc00-4c9a24955a0d/1442571030"
                        },
                        "FieldStatus": 1,
                        "LineupX": 10.0,
                        "LineupY": 8.0
                    },
                    {
                        "IdPlayer": "251352",
                        "IdTeam": "43911",
                        "ShirtNumber": 18,
                        "Status": 1,
                        "SpecialStatus": 3,
                        "Captain": true,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Andres GUARDADO"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "A. GUARDADO"
                            }
                        ],
                        "Position": 2,
                        "PlayerPicture": {
                            "Id": "9AEABAB1-850A-47DD-875047165DD3D639",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/9aeabab1-850a-47dd-8750-47165dd3d639/1442571347"
                        },
                        "FieldStatus": 2,
                        "LineupX": 2.0,
                        "LineupY": 8.0
                    },
                    {
                        "IdPlayer": "448051",
                        "IdTeam": "43911",
                        "ShirtNumber": 24,
                        "Status": 1,
                        "SpecialStatus": 2,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Luis CHAVEZ"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "L.CHAVEZ"
                            }
                        ],
                        "Position": 2,
                        "PlayerPicture": {
                            "Id": "B2BB90EA-B2D0-43E4-B0C0F0AEE5E55EB0",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/b2bb90ea-b2d0-43e4-b0c0-f0aee5e55eb0/1442571210"
                        },
                        "FieldStatus": 1,
                        "LineupX": 18.0,
                        "LineupY": 8.0
                    },
                    {
                        "IdPlayer": "430766",
                        "IdTeam": "43911",
                        "ShirtNumber": 10,
                        "Status": 1,
                        "SpecialStatus": 2,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Alexis VEGA"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Alexis Vega"
                            }
                        ],
                        "Position": 3,
                        "PlayerPicture": {
                            "Id": "34526E10-28D4-40F1-9E960323E84390EF",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/34526e10-28d4-40f1-9e96-0323e84390ef/1442570859"
                        },
                        "FieldStatus": 1,
                        "LineupX": 6.0,
                        "LineupY": 12.0
                    },
                    {
                        "IdPlayer": "386337",
                        "IdTeam": "43911",
                        "ShirtNumber": 22,
                        "Status": 1,
                        "SpecialStatus": 2,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Hirving  LOZANO"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "H. LOZANO"
                            }
                        ],
                        "Position": 3,
                        "PlayerPicture": {
                            "Id": "25256872-CD39-4782-862A45AA19BA5875",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/25256872-cd39-4782-862a-45aa19ba5875/1442570768"
                        },
                        "FieldStatus": 1,
                        "LineupX": 14.0,
                        "LineupY": 12.0
                    },
                    {
                        "IdPlayer": "175546",
                        "IdTeam": "43911",
                        "ShirtNumber": 1,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Alfredo TALAVERA"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "A. TALAVERA"
                            }
                        ],
                        "Position": 0,
                        "PlayerPicture": {
                            "Id": "B5FF7388-C2C0-45E1-BA1CC1A4A13AB634",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/b5ff7388-c2c0-45e1-ba1c-c1a4a13ab634/1442569767"
                        },
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "268714",
                        "IdTeam": "43911",
                        "ShirtNumber": 12,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Rodolfo COTA"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "R. COTA"
                            }
                        ],
                        "Position": 0,
                        "PlayerPicture": {
                            "Id": "36AD6D81-AC44-4AC8-AC456629FEDDE428",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/36ad6d81-ac44-4ac8-ac45-6629fedde428/1442569779"
                        },
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "400634",
                        "IdTeam": "43911",
                        "ShirtNumber": 4,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Edson ALVAREZ"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "E. ÁLVAREZ"
                            }
                        ],
                        "Position": 4,
                        "PlayerPicture": {
                            "Id": "8BEA6C6F-05EC-4F00-8307FAF1347E2E3E",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/8bea6c6f-05ec-4f00-8307-faf1347e2e3e/1442570749"
                        },
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "425701",
                        "IdTeam": "43911",
                        "ShirtNumber": 5,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Johan VASQUEZ"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "J. VASQUEZ"
                            }
                        ],
                        "Position": 4,
                        "PlayerPicture": {
                            "Id": "139FF36E-7958-4F6D-A16FC616AFC51546",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/139ff36e-7958-4f6d-a16f-c616afc51546/1442571172"
                        },
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "403587",
                        "IdTeam": "43911",
                        "ShirtNumber": 6,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Gerardo ARTEAGA"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "G. ARTEAGA"
                            }
                        ],
                        "Position": 4,
                        "PlayerPicture": {
                            "Id": "30F75DDE-AA1B-45D8-ACE844007B3AA644",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/30f75dde-aa1b-45d8-ace8-44007b3aa644/1442571070"
                        },
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "430763",
                        "IdTeam": "43911",
                        "ShirtNumber": 7,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Luis ROMO"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Luis Romo"
                            }
                        ],
                        "Position": 4,
                        "PlayerPicture": {
                            "Id": "D93D9965-A8DF-4358-8068012C556A8743",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/d93d9965-a8df-4358-8068-012c556a8743/1442569867"
                        },
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "425702",
                        "IdTeam": "43911",
                        "ShirtNumber": 8,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Carlos RODRIGUEZ"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "C. RODRÍGUEZ"
                            }
                        ],
                        "Position": 4,
                        "PlayerPicture": {
                            "Id": "6F06B814-70FA-4563-9752AFD201F8303D",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/6f06b814-70fa-4563-9752-afd201f8303d/1442569978"
                        },
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "356731",
                        "IdTeam": "43911",
                        "ShirtNumber": 9,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Raul JIMENEZ"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "RAÚL"
                            }
                        ],
                        "Position": 4,
                        "PlayerPicture": {
                            "Id": "38F80EE1-1325-4456-8F3FB44AF2540650",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/38f80ee1-1325-4456-8f3f-b44af2540650/1442570723"
                        },
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "333004",
                        "IdTeam": "43911",
                        "ShirtNumber": 11,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Rogelio FUNES MORI"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "R. FUNES MORI"
                            }
                        ],
                        "Position": 4,
                        "PlayerPicture": {
                            "Id": "D41C450F-98C3-48F0-A10F528A08DDA11A",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/d41c450f-98c3-48f0-a10f-528a08dda11a/1442570031"
                        },
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "386332",
                        "IdTeam": "43911",
                        "ShirtNumber": 14,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Erick GUTIERREZ"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "E. GUTIÉRREZ"
                            }
                        ],
                        "Position": 4,
                        "PlayerPicture": {
                            "Id": "55D090EB-9D0C-402C-9350AE02F6629552",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/55d090eb-9d0c-402c-9350-ae02f6629552/1442569937"
                        },
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "372090",
                        "IdTeam": "43911",
                        "ShirtNumber": 17,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Orbelin PINEDA"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "O. PINEDA"
                            }
                        ],
                        "Position": 4,
                        "PlayerPicture": {
                            "Id": "13290007-0A3B-4D62-9E7FEFCB923416C5",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/13290007-0a3b-4d62-9e7f-efcb923416c5/1442571274"
                        },
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "403596",
                        "IdTeam": "43911",
                        "ShirtNumber": 19,
                        "Status": 2,
                        "SpecialStatus": 4,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Jorge SANCHEZ"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Jorge Sánchez"
                            }
                        ],
                        "Position": 4,
                        "PlayerPicture": {
                            "Id": "8595731E-ECFB-46C3-85438D111BDCE8A4",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/8595731e-ecfb-46c3-8543-8d111bdce8a4/1442571149"
                        },
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "430765",
                        "IdTeam": "43911",
                        "ShirtNumber": 20,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Henry MARTIN"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Henry Martín"
                            }
                        ],
                        "Position": 4,
                        "PlayerPicture": {
                            "Id": "9029F6F8-3D83-4B4D-B122C909DA41C589",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/9029f6f8-3d83-4b4d-b122-c909da41c589/1442570950"
                        },
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "403586",
                        "IdTeam": "43911",
                        "ShirtNumber": 21,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Uriel ANTUNA"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "U. ANTUNA"
                            }
                        ],
                        "Position": 4,
                        "PlayerPicture": {
                            "Id": "FD0FEF42-8129-4CC0-A958F44521A742BE",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/fd0fef42-8129-4cc0-a958-f44521a742be/1442571310"
                        },
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "403585",
                        "IdTeam": "43911",
                        "ShirtNumber": 25,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Roberto ALVARADO"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Roberto Alvarado"
                            }
                        ],
                        "Position": 4,
                        "PlayerPicture": {
                            "Id": "854DFFF9-1715-4EF0-913ECF6D8E08A087",
                            "PictureUrl": "https://digitalhub.fifa.com/transform/854dfff9-1715-4ef0-913e-cf6d8e08a087/1442571123"
                        },
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    }
                ],
                "Bookings": [
                    {
                        "Card": 1,
                        "Period": 3,
                        "IdEvent": null,
                        "EventNumber": null,
                        "IdPlayer": "336707",
                        "IdCoach": null,
                        "IdTeam": "43911",
                        "Minute": "22'",
                        "Reason": null
                    },
                    {
                        "Card": 1,
                        "Period": 5,
                        "IdEvent": null,
                        "EventNumber": null,
                        "IdPlayer": "386332",
                        "IdCoach": null,
                        "IdTeam": "43911",
                        "Minute": "50'",
                        "Reason": null
                    }
                ],
                "Goals": [],
                "Substitutions": [
                    {
                        "IdEvent": null,
                        "Period": 3,
                        "Reason": 0,
                        "SubstitutePosition": 4,
                        "IdPlayerOff": "251352",
                        "IdPlayerOn": "386332",
                        "PlayerOffName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Andres GUARDADO"
                            }
                        ],
                        "PlayerOnName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Erick GUTIERREZ"
                            }
                        ],
                        "Minute": "42'",
                        "IdTeam": "43911"
                    }
                ],
                "FootballType": 0,
                "Gender": 1,
                "IdAssociation": "MEX",
                "ShortClubName": "Mexico"
            },
            "BallPossession": {
                "Intervals": [],
                "LastX": [],
                "OverallHome": 65.75546,
                "OverallAway": 34.24454
            },
            "TerritorialPossesion": null,
            "TerritorialThirdPossesion": null,
            "Officials": [
                {
                    "IdCountry": "ITA",
                    "OfficialId": "315593",
                    "NameShort": [
                        {
                            "Locale": "en-GB",
                            "Description": "Daniele ORSATO"
                        }
                    ],
                    "Name": [
                        {
                            "Locale": "en-GB",
                            "Description": "Daniele Orsato"
                        }
                    ],
                    "OfficialType": 1,
                    "TypeLocalized": [
                        {
                            "Locale": "en-GB",
                            "Description": "Referee"
                        }
                    ]
                },
                {
                    "IdCountry": "ITA",
                    "OfficialId": "415967",
                    "NameShort": [
                        {
                            "Locale": "en-GB",
                            "Description": "Ciro CARBONE"
                        }
                    ],
                    "Name": [
                        {
                            "Locale": "en-GB",
                            "Description": "Ciro Carbone"
                        }
                    ],
                    "OfficialType": 2,
                    "TypeLocalized": [
                        {
                            "Locale": "en-GB",
                            "Description": "Assistant Referee 1"
                        }
                    ]
                },
                {
                    "IdCountry": "ITA",
                    "OfficialId": "360574",
                    "NameShort": [
                        {
                            "Locale": "en-GB",
                            "Description": "Alessandro GIALLATINI"
                        }
                    ],
                    "Name": [
                        {
                            "Locale": "en-GB",
                            "Description": "Alessandro Giallatini"
                        }
                    ],
                    "OfficialType": 3,
                    "TypeLocalized": [
                        {
                            "Locale": "en-GB",
                            "Description": "Assistant Referee 2"
                        }
                    ]
                },
                {
                    "IdCountry": "ROU",
                    "OfficialId": "314605",
                    "NameShort": [
                        {
                            "Locale": "en-GB",
                            "Description": "Istvan KOVACS"
                        }
                    ],
                    "Name": [
                        {
                            "Locale": "en-GB",
                            "Description": "Istvan Kovacs"
                        }
                    ],
                    "OfficialType": 4,
                    "TypeLocalized": [
                        {
                            "Locale": "en-GB",
                            "Description": "Fourth official"
                        }
                    ]
                },
                {
                    "IdCountry": "ITA",
                    "OfficialId": "399108",
                    "NameShort": [
                        {
                            "Locale": "en-GB",
                            "Description": "Massimiliano IRRATI"
                        }
                    ],
                    "Name": [
                        {
                            "Locale": "en-GB",
                            "Description": "Massimiliano Irrati"
                        }
                    ],
                    "OfficialType": 5,
                    "TypeLocalized": [
                        {
                            "Locale": "en-GB",
                            "Description": "Video Assistant Referee (VAR)"
                        }
                    ]
                },
                {
                    "IdCountry": "ESP",
                    "OfficialId": "280881",
                    "NameShort": [
                        {
                            "Locale": "en-GB",
                            "Description": "Roberto DIAZ"
                        }
                    ],
                    "Name": [
                        {
                            "Locale": "en-GB",
                            "Description": "Roberto Díaz Pérez Del Palomar"
                        }
                    ],
                    "OfficialType": 7,
                    "TypeLocalized": [
                        {
                            "Locale": "en-GB",
                            "Description": "Offside VAR"
                        }
                    ]
                },
                {
                    "IdCountry": "ITA",
                    "OfficialId": "328434",
                    "NameShort": [
                        {
                            "Locale": "en-GB",
                            "Description": "Paolo VALERI"
                        }
                    ],
                    "Name": [
                        {
                            "Locale": "en-GB",
                            "Description": "Paolo Valeri"
                        }
                    ],
                    "OfficialType": 8,
                    "TypeLocalized": [
                        {
                            "Locale": "en-GB",
                            "Description": "Assistant VAR"
                        }
                    ]
                },
                {
                    "IdCountry": "SGP",
                    "OfficialId": "347099",
                    "NameShort": [
                        {
                            "Locale": "en-GB",
                            "Description": "Muhammad BIN JAHARI"
                        }
                    ],
                    "Name": [
                        {
                            "Locale": "en-GB",
                            "Description": "Muhammad Taqi Aljaafari bin Jahari"
                        }
                    ],
                    "OfficialType": 9,
                    "TypeLocalized": [
                        {
                            "Locale": "en-GB",
                            "Description": "Support VAR"
                        }
                    ]
                },
                {
                    "IdCountry": "ROU",
                    "OfficialId": "361587",
                    "NameShort": [
                        {
                            "Locale": "en-GB",
                            "Description": "Mihai ARTENE"
                        }
                    ],
                    "Name": [
                        {
                            "Locale": "en-GB",
                            "Description": "Mihai Ovidiu Artene"
                        }
                    ],
                    "OfficialType": 10,
                    "TypeLocalized": [
                        {
                            "Locale": "en-GB",
                            "Description": "Reserve Assistant Referee"
                        }
                    ]
                }
            ],
            "MatchStatus": 3,
            "GroupName": [
                {
                    "Locale": "en-GB",
                    "Description": "Group C"
                }
            ],
            "StageName": [
                {
                    "Locale": "en-GB",
                    "Description": "First stage"
                }
            ],
            "OfficialityStatus": 0,
            "TimeDefined": true,
            "Properties": {
                "IdIFES": "133014"
            },
            "IsUpdateable": null
        },
        {
            "IdMatch": "75knysm67l3iv2jgrqini4cgk",
            "IdStage": "9pwuiojh0jtyf5ifb3nzx4sgk",
            "IdGroup": "9q7c6vw7r4z1kchyu9ebbdlp0",
            "IdSeason": "9pfmrnznfsm69226miwy41dec",
            "IdCompetition": "bqvy41un7sf86rbse9tv810x7",
            "CompetitionName": [
                {
                    "Locale": "en-gb",
                    "Description": "Taça da Liga"
                }
            ],
            "SeasonName": [
                {
                    "Locale": "en-gb",
                    "Description": "Taça da Liga 2022/2023"
                }
            ],
            "SeasonShortName": [],
            "Stadium": {
                "IdStadium": null,
                "Name": [
                    {
                        "Locale": "en-gb",
                        "Description": "Estádio Municipal de Braga"
                    }
                ],
                "Capacity": 30286,
                "WebAddress": null,
                "Built": "2003-01-01T00:00:00Z",
                "Roof": false,
                "Turf": null,
                "IdCity": "400015080",
                "CityName": [
                    {
                        "Locale": "en-GB",
                        "Description": "Braga"
                    }
                ],
                "IdCountry": "POR",
                "PostalCode": null,
                "Street": "Avenida do Estádio",
                "Email": null,
                "Fax": null,
                "Phone": null,
                "AffiliationCountry": null,
                "AffiliationRegion": null,
                "Latitude": 41.562553702,
                "Longitude": -8.429903984,
                "Length": null,
                "Width": null,
                "Properties": {},
                "IsUpdateable": null
            },
            "ResultType": 1,
            "MatchDay": "1",
            "HomeTeamPenaltyScore": null,
            "AwayTeamPenaltyScore": null,
            "AggregateHomeTeamScore": null,
            "AggregateAwayTeamScore": null,
            "Weather": null,
            "Attendance": null,
            "Date": "2022-11-26T18:45:00Z",
            "LocalDate": "2022-11-26T18:45:00Z",
            "MatchTime": "75'",
            "SecondHalfTime": null,
            "FirstHalfTime": null,
            "FirstHalfExtraTime": 5,
            "SecondHalfExtraTime": 0,
            "Winner": null,
            "Period": 5,
            "HomeTeam": {
                "Score": 2,
                "Side": null,
                "IdTeam": "1884019",
                "PictureUrl": "https://api.fifa.com/api/v3/picture/teams-{format}-{size}/1884019",
                "IdCountry": "POR",
                "TeamType": 0,
                "AgeType": 7,
                "Tactics": "4-4-2",
                "TeamName": [
                    {
                        "Locale": "en-gb",
                        "Description": "Sporting Braga"
                    }
                ],
                "Abbreviation": "BRA",
                "Coaches": [
                    {
                        "IdCoach": "bvh4r1ysizjvelbzsn5rdk639",
                        "IdCountry": "POR",
                        "PictureUrl": null,
                        "Name": [
                            {
                                "Locale": "en-gb",
                                "Description": "Artur Jorge Torres Gomes Araújo Amorim"
                            }
                        ],
                        "Alias": [
                            {
                                "Locale": "en-gb",
                                "Description": "Artur Jorge"
                            }
                        ],
                        "Role": 0,
                        "SpecialStatus": null
                    }
                ],
                "Players": [
                    {
                        "IdPlayer": "384761",
                        "IdTeam": "1884019",
                        "ShirtNumber": 12,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": true,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "TIAGO SA"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "TIAGO SÁ"
                            }
                        ],
                        "Position": 0,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "405574",
                        "IdTeam": "1884019",
                        "ShirtNumber": 2,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Victor PEREA"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "VICTOR"
                            }
                        ],
                        "Position": 1,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "409236",
                        "IdTeam": "1884019",
                        "ShirtNumber": 4,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Sikou NIAKATE"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "S. Niakaté"
                            }
                        ],
                        "Position": 1,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "1l93qv71yu7ohs9ida8cx2zqi",
                        "IdTeam": "1884019",
                        "ShirtNumber": 24,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Bruno Rodrigues"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Bruno Rodrigues"
                            }
                        ],
                        "Position": 1,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "394336",
                        "IdTeam": "1884019",
                        "ShirtNumber": 26,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Cristian BORJA"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Cristian BORJA"
                            }
                        ],
                        "Position": 1,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "396013",
                        "IdTeam": "1884019",
                        "ShirtNumber": 8,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Mohamed ALMOATASEMBELLAH"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Almoatasembellah AL MUSRATI"
                            }
                        ],
                        "Position": 2,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "33qmom9yrc1ft1cte9qju1djp",
                        "IdTeam": "1884019",
                        "ShirtNumber": 10,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "André Horta"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "André Horta"
                            }
                        ],
                        "Position": 2,
                        "PlayerPicture": null,
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "406138",
                        "IdTeam": "1884019",
                        "ShirtNumber": 18,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Diego LAINEZ"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "D. LAINEZ"
                            }
                        ],
                        "Position": 2,
                        "PlayerPicture": null,
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "395209",
                        "IdTeam": "1884019",
                        "ShirtNumber": 45,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Iuri MEDEIROS"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Iuri Medeiros"
                            }
                        ],
                        "Position": 2,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "405569",
                        "IdTeam": "1884019",
                        "ShirtNumber": 9,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Abel RUIZ"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Abel Ruiz"
                            }
                        ],
                        "Position": 3,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "43xutn1unln5o3gwx3cl4lo0l",
                        "IdTeam": "1884019",
                        "ShirtNumber": 23,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Simon Banza"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "S. Banza"
                            }
                        ],
                        "Position": 3,
                        "PlayerPicture": null,
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "5nlzolv5bg03eiopba58rs251",
                        "IdTeam": "1884019",
                        "ShirtNumber": 1,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Matheus"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Matheus"
                            }
                        ],
                        "Position": 0,
                        "PlayerPicture": null,
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "88hqy3ztjmtahcv0sz6rduv6y",
                        "IdTeam": "1884019",
                        "ShirtNumber": 5,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Serdar Saatçi"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "S. Saatçı"
                            }
                        ],
                        "Position": 1,
                        "PlayerPicture": null,
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "dq41kziilg4jki9erxel9d3ze",
                        "IdTeam": "1884019",
                        "ShirtNumber": 70,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Fabiano"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Fabiano"
                            }
                        ],
                        "Position": 1,
                        "PlayerPicture": null,
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "5p4q1zolt4e4h07jzknksyzv8",
                        "IdTeam": "1884019",
                        "ShirtNumber": 29,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Jean-Baptiste Gorby"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "J. Gorby"
                            }
                        ],
                        "Position": 2,
                        "PlayerPicture": null,
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "5bo84b7398nyv87ztwd8wxh3p",
                        "IdTeam": "1884019",
                        "ShirtNumber": 88,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Castro"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Castro"
                            }
                        ],
                        "Position": 2,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "2apqk2zobhs8ypnaw0bzkch9m",
                        "IdTeam": "1884019",
                        "ShirtNumber": 7,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Rodrigo Gomes"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Rodrigo Gomes"
                            }
                        ],
                        "Position": 3,
                        "PlayerPicture": null,
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "23d0t7gqlra3jkskn18bh752i",
                        "IdTeam": "1884019",
                        "ShirtNumber": 14,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Álvaro Djaló"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Álvaro Djaló"
                            }
                        ],
                        "Position": 3,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "4ocmdia4dhnz2jdxr3n8sf8r8",
                        "IdTeam": "1884019",
                        "ShirtNumber": 71,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Hernâni Infande"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Hernâni Infande"
                            }
                        ],
                        "Position": 3,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "cnriqwh17pqgz7p2gt74uar6c",
                        "IdTeam": "1884019",
                        "ShirtNumber": 99,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Vitinha"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Vítinha"
                            }
                        ],
                        "Position": 3,
                        "PlayerPicture": null,
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    }
                ],
                "Bookings": [
                    {
                        "Card": 1,
                        "Period": 3,
                        "IdEvent": null,
                        "EventNumber": null,
                        "IdPlayer": "409236",
                        "IdCoach": null,
                        "IdTeam": "1884019",
                        "Minute": "39:00",
                        "Reason": "Foul"
                    },
                    {
                        "Card": 1,
                        "Period": 5,
                        "IdEvent": null,
                        "EventNumber": null,
                        "IdPlayer": "406138",
                        "IdCoach": null,
                        "IdTeam": "1884019",
                        "Minute": "64:00",
                        "Reason": "Foul"
                    }
                ],
                "Goals": [
                    {
                        "Type": 2,
                        "IdPlayer": "43xutn1unln5o3gwx3cl4lo0l",
                        "Minute": "16:00",
                        "IdAssistPlayer": "405569",
                        "Period": 3,
                        "IdGoal": null,
                        "IdTeam": "1884019"
                    },
                    {
                        "Type": 2,
                        "IdPlayer": "396013",
                        "Minute": "42:00",
                        "IdAssistPlayer": "33qmom9yrc1ft1cte9qju1djp",
                        "Period": 3,
                        "IdGoal": null,
                        "IdTeam": "1884019"
                    }
                ],
                "Substitutions": [
                    {
                        "IdEvent": null,
                        "Period": 5,
                        "Reason": 2,
                        "SubstitutePosition": 3,
                        "IdPlayerOff": "406138",
                        "IdPlayerOn": "23d0t7gqlra3jkskn18bh752i",
                        "PlayerOffName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Diego LAINEZ"
                            }
                        ],
                        "PlayerOnName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Álvaro Djaló"
                            }
                        ],
                        "Minute": "67:00",
                        "IdTeam": "1884019"
                    },
                    {
                        "IdEvent": null,
                        "Period": 5,
                        "Reason": 2,
                        "SubstitutePosition": 2,
                        "IdPlayerOff": "43xutn1unln5o3gwx3cl4lo0l",
                        "IdPlayerOn": "5bo84b7398nyv87ztwd8wxh3p",
                        "PlayerOffName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Simon Banza"
                            }
                        ],
                        "PlayerOnName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Castro"
                            }
                        ],
                        "Minute": "67:00",
                        "IdTeam": "1884019"
                    },
                    {
                        "IdEvent": null,
                        "Period": 5,
                        "Reason": 2,
                        "SubstitutePosition": 3,
                        "IdPlayerOff": "33qmom9yrc1ft1cte9qju1djp",
                        "IdPlayerOn": "4ocmdia4dhnz2jdxr3n8sf8r8",
                        "PlayerOffName": [
                            {
                                "Locale": "en-gb",
                                "Description": "André Horta"
                            }
                        ],
                        "PlayerOnName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Hernâni Infande"
                            }
                        ],
                        "Minute": "67:00",
                        "IdTeam": "1884019"
                    }
                ],
                "FootballType": 0,
                "Gender": 1,
                "IdAssociation": "POR",
                "ShortClubName": "Braga"
            },
            "AwayTeam": {
                "Score": 0,
                "Side": null,
                "IdTeam": "cud4wllxnak5zcd7oei6xrj25",
                "PictureUrl": "https://api.fifa.com/api/v3/picture/teams-{format}-{size}/cud4wllxnak5zcd7oei6xrj25",
                "IdCountry": "POR",
                "TeamType": 0,
                "AgeType": 7,
                "Tactics": "5-4-1",
                "TeamName": [
                    {
                        "Locale": "en-gb",
                        "Description": "Trofense"
                    }
                ],
                "Abbreviation": "TRO",
                "Coaches": [],
                "Players": [
                    {
                        "IdPlayer": "13icgqmte1fd2k9izty14pca2",
                        "IdTeam": "cud4wllxnak5zcd7oei6xrj25",
                        "ShirtNumber": 1,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Tiago Silva"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Tiago Silva"
                            }
                        ],
                        "Position": 0,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "bieo7sqzdulo07jngz6mux5nt",
                        "IdTeam": "cud4wllxnak5zcd7oei6xrj25",
                        "ShirtNumber": 14,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Simão Martins"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Simão Martins"
                            }
                        ],
                        "Position": 1,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "9o40x4xi2n34wperqfbpor6c5",
                        "IdTeam": "cud4wllxnak5zcd7oei6xrj25",
                        "ShirtNumber": 21,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Marcos Valente"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Marcos Valente"
                            }
                        ],
                        "Position": 1,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "9epqc9jcj8ooissi177o2qc8a",
                        "IdTeam": "cud4wllxnak5zcd7oei6xrj25",
                        "ShirtNumber": 44,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Rúben Pereira"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Rúben Pereira"
                            }
                        ],
                        "Position": 1,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "armkm0zfocqw8oaq3daarppwq",
                        "IdTeam": "cud4wllxnak5zcd7oei6xrj25",
                        "ShirtNumber": 5,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Abdoul Bandaogo"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "A. Bandaogo"
                            }
                        ],
                        "Position": 2,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "416825",
                        "IdTeam": "cud4wllxnak5zcd7oei6xrj25",
                        "ShirtNumber": 10,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "Vanilson Tita Zéu"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Vanilson"
                            }
                        ],
                        "Position": 2,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "8q4vyfjrbrx2q6nc0t03h72dx",
                        "IdTeam": "cud4wllxnak5zcd7oei6xrj25",
                        "ShirtNumber": 29,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": true,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Vasco Rocha"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Vasco Rocha"
                            }
                        ],
                        "Position": 2,
                        "PlayerPicture": null,
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "320313",
                        "IdTeam": "cud4wllxnak5zcd7oei6xrj25",
                        "ShirtNumber": 30,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "DJALMA"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-GB",
                                "Description": "DJALMA"
                            }
                        ],
                        "Position": 2,
                        "PlayerPicture": null,
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "df4radomvkn1ucra7sohjq28l",
                        "IdTeam": "cud4wllxnak5zcd7oei6xrj25",
                        "ShirtNumber": 19,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Stevy Okitokandjo"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "S. Okitokandjo"
                            }
                        ],
                        "Position": 3,
                        "PlayerPicture": null,
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "d29pmst1t7125beoyq16ab2ay",
                        "IdTeam": "cud4wllxnak5zcd7oei6xrj25",
                        "ShirtNumber": 11,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Daniel Liberal"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Daniel Liberal"
                            }
                        ],
                        "Position": 5,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "7mjfbohe0dd4xnhkazbi8htrd",
                        "IdTeam": "cud4wllxnak5zcd7oei6xrj25",
                        "ShirtNumber": 55,
                        "Status": 1,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Tiago André"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Tiago André"
                            }
                        ],
                        "Position": 5,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "3y3q3novm3ddi8qqwo6prts2y",
                        "IdTeam": "cud4wllxnak5zcd7oei6xrj25",
                        "ShirtNumber": 36,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Caio Marcelo"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Caio Marcelo"
                            }
                        ],
                        "Position": 1,
                        "PlayerPicture": null,
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "210xmaquqj9ouxnxm0bwbiwyy",
                        "IdTeam": "cud4wllxnak5zcd7oei6xrj25",
                        "ShirtNumber": 48,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Tiago Manso"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Tiago Manso"
                            }
                        ],
                        "Position": 1,
                        "PlayerPicture": null,
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "9xafo4gst8wkrjehmbjqzntyi",
                        "IdTeam": "cud4wllxnak5zcd7oei6xrj25",
                        "ShirtNumber": 6,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Martim Maia"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Martim Maia"
                            }
                        ],
                        "Position": 2,
                        "PlayerPicture": null,
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "a7hxr0rue48y788y2uiptmhud",
                        "IdTeam": "cud4wllxnak5zcd7oei6xrj25",
                        "ShirtNumber": 8,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Andrézinho"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Andrézinho"
                            }
                        ],
                        "Position": 2,
                        "PlayerPicture": null,
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "424652",
                        "IdTeam": "cud4wllxnak5zcd7oei6xrj25",
                        "ShirtNumber": 16,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-GB",
                                "Description": "BENI"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Beni"
                            }
                        ],
                        "Position": 2,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "wnwhajljya6h4tn8spxg9k6i",
                        "IdTeam": "cud4wllxnak5zcd7oei6xrj25",
                        "ShirtNumber": 18,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Welves"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Welves"
                            }
                        ],
                        "Position": 2,
                        "PlayerPicture": null,
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "1021va5gmssosnew1t7ph8odm",
                        "IdTeam": "cud4wllxnak5zcd7oei6xrj25",
                        "ShirtNumber": 97,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Youcef Bechou"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Y. Bechou"
                            }
                        ],
                        "Position": 2,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "1ajurpd3v8f3a3oojmq1hravp",
                        "IdTeam": "cud4wllxnak5zcd7oei6xrj25",
                        "ShirtNumber": 7,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Traquina"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Traquina"
                            }
                        ],
                        "Position": 3,
                        "PlayerPicture": null,
                        "FieldStatus": 2,
                        "LineupX": null,
                        "LineupY": null
                    },
                    {
                        "IdPlayer": "8ucn0gmb0o32s5z0e86v3eui1",
                        "IdTeam": "cud4wllxnak5zcd7oei6xrj25",
                        "ShirtNumber": 33,
                        "Status": 2,
                        "SpecialStatus": null,
                        "Captain": false,
                        "PlayerName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Pachu"
                            }
                        ],
                        "ShortName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Pachu"
                            }
                        ],
                        "Position": 3,
                        "PlayerPicture": null,
                        "FieldStatus": 1,
                        "LineupX": null,
                        "LineupY": null
                    }
                ],
                "Bookings": [
                    {
                        "Card": 1,
                        "Period": 3,
                        "IdEvent": null,
                        "EventNumber": null,
                        "IdPlayer": "df4radomvkn1ucra7sohjq28l",
                        "IdCoach": null,
                        "IdTeam": "cud4wllxnak5zcd7oei6xrj25",
                        "Minute": "47:00",
                        "Reason": "Foul"
                    }
                ],
                "Goals": [],
                "Substitutions": [
                    {
                        "IdEvent": null,
                        "Period": 5,
                        "Reason": 2,
                        "SubstitutePosition": 3,
                        "IdPlayerOff": "df4radomvkn1ucra7sohjq28l",
                        "IdPlayerOn": "8ucn0gmb0o32s5z0e86v3eui1",
                        "PlayerOffName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Stevy Okitokandjo"
                            }
                        ],
                        "PlayerOnName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Pachu"
                            }
                        ],
                        "Minute": "45:00",
                        "IdTeam": "cud4wllxnak5zcd7oei6xrj25"
                    },
                    {
                        "IdEvent": null,
                        "Period": 5,
                        "Reason": 2,
                        "SubstitutePosition": 2,
                        "IdPlayerOff": "8q4vyfjrbrx2q6nc0t03h72dx",
                        "IdPlayerOn": "1021va5gmssosnew1t7ph8odm",
                        "PlayerOffName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Vasco Rocha"
                            }
                        ],
                        "PlayerOnName": [
                            {
                                "Locale": "en-gb",
                                "Description": "Youcef Bechou"
                            }
                        ],
                        "Minute": "65:00",
                        "IdTeam": "cud4wllxnak5zcd7oei6xrj25"
                    },
                    {
                        "IdEvent": null,
                        "Period": 5,
                        "Reason": 2,
                        "SubstitutePosition": 2,
                        "IdPlayerOff": "320313",
                        "IdPlayerOn": "424652",
                        "PlayerOffName": [
                            {
                                "Locale": "en-GB",
                                "Description": "DJALMA"
                            }
                        ],
                        "PlayerOnName": [
                            {
                                "Locale": "en-GB",
                                "Description": "BENI"
                            }
                        ],
                        "Minute": "65:00",
                        "IdTeam": "cud4wllxnak5zcd7oei6xrj25"
                    }
                ],
                "FootballType": 0,
                "Gender": 1,
                "IdAssociation": "POR",
                "ShortClubName": "Trofense"
            },
            "BallPossession": null,
            "TerritorialPossesion": null,
            "TerritorialThirdPossesion": null,
            "Officials": [
                {
                    "IdCountry": "POR",
                    "OfficialId": "327547",
                    "NameShort": [
                        {
                            "Locale": "en-gb",
                            "Description": "Ricardo Baixinho"
                        }
                    ],
                    "Name": [
                        {
                            "Locale": "en-GB",
                            "Description": "Ricardo Jorge Antunes Baixinho"
                        }
                    ],
                    "OfficialType": 1,
                    "TypeLocalized": [
                        {
                            "Locale": "en-GB",
                            "Description": "Referee"
                        }
                    ]
                },
                {
                    "IdCountry": "POR",
                    "OfficialId": "432893",
                    "NameShort": [
                        {
                            "Locale": "en-gb",
                            "Description": "Pedro Martins"
                        }
                    ],
                    "Name": [
                        {
                            "Locale": "en-GB",
                            "Description": "Pedro Nuno de Sá Martins"
                        }
                    ],
                    "OfficialType": 2,
                    "TypeLocalized": [
                        {
                            "Locale": "en-GB",
                            "Description": "Assistant Referee 1"
                        }
                    ]
                },
                {
                    "IdCountry": "POR",
                    "OfficialId": "3qno8pu5sd4qbvmqd6zjr58q2",
                    "NameShort": [
                        {
                            "Locale": "en-gb",
                            "Description": "Hugo Marques"
                        }
                    ],
                    "Name": [
                        {
                            "Locale": "en-gb",
                            "Description": "Hugo Manuel Freire Marques"
                        }
                    ],
                    "OfficialType": 3,
                    "TypeLocalized": [
                        {
                            "Locale": "en-GB",
                            "Description": "Assistant Referee 2"
                        }
                    ]
                },
                {
                    "IdCountry": "POR",
                    "OfficialId": "7aqwmak995g6fnlubjnuyqu39",
                    "NameShort": [
                        {
                            "Locale": "en-gb",
                            "Description": "Rui Marcelo Rodrigues"
                        }
                    ],
                    "Name": [
                        {
                            "Locale": "en-gb",
                            "Description": "Rui Marcelo Rodrigues"
                        }
                    ],
                    "OfficialType": 4,
                    "TypeLocalized": [
                        {
                            "Locale": "en-GB",
                            "Description": "Fourth official"
                        }
                    ]
                }
            ],
            "MatchStatus": 3,
            "GroupName": [
                {
                    "Locale": "en-gb",
                    "Description": "Group D"
                }
            ],
            "StageName": [
                {
                    "Locale": "en-gb",
                    "Description": "Group Stage"
                }
            ],
            "OfficialityStatus": 0,
            "TimeDefined": true,
            "Properties": {
                "IdStatsPerform": "75knysm67l3iv2jgrqini4cgk"
            },
            "IsUpdateable": null
        }
    ]
}
`

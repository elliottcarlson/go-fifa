package go_fifa_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	fifa "github.com/imdevinc/go-fifa"
	"github.com/stretchr/testify/assert"
)

var expTimelineResp = fifa.TimelineResponse{
	StageId:       "285063",
	MatchId:       "400128082",
	CompetitionId: "17",
	SeasonId:      "255711",
	GroupId:       "285065",
	Events: []fifa.TimelineEvent{
		{
			Id:          "18185300000030",
			TeamId:      "43834",
			Timestamp:   time.Date(2022, 11, 20, 16, 01, 30, 0, time.UTC),
			MatchMinute: "1'",
			Period:      fifa.NotStarted,
			HomeGoals:   0,
			AwayGoals:   0,
			Type:        fifa.CoinToss,
			TypeLocalized: []fifa.LocaleDescription{{
				Locale:      "en-GB",
				Description: "Coin Toss",
			}},
			HomePenaltyGoals: 0,
			AwayPenaltyGoals: 0,
			Description:      []fifa.LocaleDescription{},
			Qualifiers:       []any{},
		},
		{
			Id:          "18185300000031",
			Timestamp:   time.Date(2022, 11, 20, 16, 02, 13, 0, time.UTC),
			MatchMinute: "1'",
			Period:      fifa.FirstPeriod,
			HomeGoals:   0,
			AwayGoals:   0,
			Type:        fifa.MatchStart,
			TypeLocalized: []fifa.LocaleDescription{{
				Locale:      "en-GB",
				Description: "Start Time",
			}},
			HomePenaltyGoals: 0,
			AwayPenaltyGoals: 0,
			Description: []fifa.LocaleDescription{{
				Locale:      "en-GB",
				Description: "The referee signals the start of the first period.",
			}},
			Qualifiers: []any{},
		},
		{
			Id:          "18185300001525",
			Timestamp:   time.Date(2022, 11, 20, 17, 59, 31, 0, time.UTC),
			MatchMinute: "90'+6'",
			Period:      fifa.SecondPeriod,
			HomeGoals:   0,
			AwayGoals:   2,
			Type:        fifa.HalfEnd,
			TypeLocalized: []fifa.LocaleDescription{{
				Locale:      "en-GB",
				Description: "End Time",
			}},
			HomePenaltyGoals: 0,
			AwayPenaltyGoals: 0,
			Description: []fifa.LocaleDescription{{
				Locale:      "en-GB",
				Description: "The referee brings the second period to an end.",
			}},
			Qualifiers: []any{},
		},
		{
			Id:          "18185300001526",
			Timestamp:   time.Date(2022, 11, 20, 17, 59, 34, 0, time.UTC),
			MatchMinute: "90'+6'",
			Period:      fifa.MatchOver,
			HomeGoals:   0,
			AwayGoals:   2,
			Type:        fifa.MatchEnd,
			TypeLocalized: []fifa.LocaleDescription{{
				Locale:      "en-GB",
				Description: "Match end",
			}},
			HomePenaltyGoals: 0,
			AwayPenaltyGoals: 0,
			Description: []fifa.LocaleDescription{{
				Locale:      "en-GB",
				Description: "The final whistle sounds.",
			}},
			Qualifiers: []any{},
		},
	},
}

func TestGetMatchEvents(t *testing.T) {
	t.Parallel()
	client := fifa.Client{
		Client: &TestClient{},
	}
	resp, err := client.GetMatchEvents(&fifa.GetMatchEventOptions{
		CompetitionId: "17",
		SeasonId:      "255711",
		StageId:       "285063",
		MatchId:       "400128082",
	})
	if ok := assert.Nil(t, err, "expected no error with GetMatchEvents, got: %s", err); !ok {
		t.FailNow()
	}
	diff := cmp.Diff(resp, expTimelineResp)
	if ok := assert.Equal(t, diff, "", diff); !ok {
		t.Fail()
	}
}

// This tests against a real match to make sure our marshalling is correct
func TestLiveEvents(t *testing.T) {
	t.Parallel()
	client := fifa.Client{}
	_, err := client.GetMatchEvents(&fifa.GetMatchEventOptions{
		CompetitionId: "17",
		SeasonId:      "255711",
		StageId:       "285063",
		MatchId:       "400235487",
	})
	if ok := assert.Nil(t, err, "expected no error with GetMatchEvents, got: %s", err); !ok {
		t.FailNow()
	}
}

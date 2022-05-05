package go_fifa

import "time"

type Gender int

const (
	MALE   Gender = 1
	FEMALE Gender = 2
)

type PeriodEnum int

const (
	FIRST        PeriodEnum = 3
	SECOND       PeriodEnum = 5
	FIRST_EXTRA  PeriodEnum = 7
	SECOND_EXTRA PeriodEnum = 9
	SHOOTOUT     PeriodEnum = 11
)

type MatchEvent int

const (
	GOAL_SCORED        MatchEvent = 0
	YELLOW_CARD        MatchEvent = 2
	RED_CARD           MatchEvent = 3
	DOUBLE_YELLOW      MatchEvent = 4
	SUBSTITUTION       MatchEvent = 5
	IGNORE             MatchEvent = 6
	MATCH_START        MatchEvent = 7
	HALF_END           MatchEvent = 8
	BLOCKED_SHOT       MatchEvent = 12
	FOUL_UNKNOWN       MatchEvent = 14
	OFFSIdE            MatchEvent = 15
	CORNER_KICK        MatchEvent = 16
	BLOCKED_SHOT_2     MatchEvent = 17
	FOUL               MatchEvent = 18
	MATCH_END          MatchEvent = 26
	CROSSBAR           MatchEvent = 32
	CROSSBAR_2         MatchEvent = 33
	OWN_GOAL           MatchEvent = 34
	HAND_BALL          MatchEvent = 37
	FREE_KICK_GOAL     MatchEvent = 39
	PENALTY_GOAL       MatchEvent = 41
	FREE_KICK_CROSSBAR MatchEvent = 44
	PENALTY_MISSED     MatchEvent = 60
	PENALTY_MISSED_2   MatchEvent = 65
	VAR_PENALTY        MatchEvent = 72
)

type PaginatedResponse struct {
	ContinuationToken string `json:"ContinuationToken"`
	ContinuationHash  string `json:"ContinuationHash"`
}

type MatchResponse struct {
	Id                        string                       `json:"IdMatch"`
	StageId                   string                       `json:"IdStage"`
	GroupId                   string                       `json:"IdGroup"`
	SeasonId                  string                       `json:"IdSeason"`
	CompetitionId             string                       `json:"IdCompetition"`
	Competition               []DefaultDescriptionResponse `json:"CompetitionName"`
	Season                    []DefaultDescriptionResponse `json:"SeasonName"`
	SeasonShortName           []interface{}                `json:"SeasonShortName"`
	Stadium                   StadiumResponse              `json:"Stadium"`
	ResultType                int                          `json:"ResultType"`
	MatchDay                  string                       `json:"MatchDay"`
	HomeTeamPenaltyScore      int                          `json:"HomeTeamPenaltyScore"`
	AwayTeamPenaltyScore      int                          `json:"AwayTeamPenaltyScore"`
	AggregateHomeTeamScore    int                          `json:"AggregateHomeTeamScore"`
	AggregateAwayTeamScore    int                          `json:"AggregateAwayTeamScore"`
	Weather                   WeatherResponse              `json:"Weather"`
	Date                      time.Time                    `json:"Date"`
	LocalDate                 time.Time                    `json:"LocalDate"`
	MatchTime                 string                       `json:"MatchTime"`
	SecondHalfTime            string                       `json:"SecondHalfTime"`
	FirstHalfTime             string                       `json:"FirstHalfTime"`
	FirstHalfExtraTime        int                          `json:"FirstHalfExtraTime"`
	SecondHalfExtraTime       int                          `json:"SecondHalfExtraTime"`
	WinnerId                  string                       `json:"Winner"`
	Period                    PeriodEnum                   `json:"Period"`
	HomeTeam                  TeamResponse                 `json:"HomeTeam"`
	AwayTeam                  TeamResponse                 `json:"AwayTeam"`
	BallPossession            BallPossessionResponse       `json:"BallPossession"`
	TerritorialPossesion      string                       `json:"TerritorialPossesion"`
	TerritorialThirdPossesion string                       `json:"TerritorialThirdPossesion"`
	Officials                 []OfficialResponse           `json:"Officials"`
	Status                    int                          `json:"MatchStatus"`
	GroupName                 []DefaultDescriptionResponse `json:"GroupName"`
	StageName                 []DefaultDescriptionResponse `json:"StageName"`
	OfficialityStatus         int                          `json:"OfficialityStatus"`
	TimeDefined               bool                         `json:"TimeDefined"`
	Properties                interface{}                  `json:"Properties"`
	IsUpdateable              bool                         `json:"IsUpdateable"`
}

type DefaultDescriptionResponse struct {
	Locale      string `json:"Locale"`
	Description string `json:"Description"`
}

type StadiumResponse struct {
	Id                 string                       `json:"IdStadium"`
	Name               []DefaultDescriptionResponse `json:"Name"`
	Capacity           string                       `json:"Capacity"`
	WebAddress         string                       `json:"WebAddress"`
	Built              string                       `json:"Built"`
	HasRoof            bool                         `json:"Roof"`
	Turf               string                       `json:"Turf"`
	CityId             string                       `json:"IdCity"`
	City               []DefaultDescriptionResponse `json:"CityName"`
	CountryId          string                       `json:"IdCountry"`
	PostalCode         string                       `json:"PostcalCode"`
	Street             string                       `json:"Street"`
	Email              string                       `json:"Email"`
	Fax                string                       `json:"Fax"`
	Phone              string                       `json:"Phone"`
	AffiliationCountry string                       `json:"AffiliationCountry"`
	AffiliationRegion  string                       `json:"AffiliationRegion"`
	Latitude           string                       `json:"Latitude"`
	Longitude          string                       `json:"Longitude"`
	Length             string                       `json:"Length"`
	WIdth              string                       `json:"WIdth"`
	Properties         interface{}                  `json:"Properties"`
	IsUpdateable       bool                         `json:"IsUpdateable"`
}

type TeamResponse struct {
	Score         int                          `json:"Score"`
	SIde          string                       `json:"SIde"`
	Id            string                       `json:"TeamId"`
	PictureURL    string                       `json:"PictureURL"`
	CountryId     string                       `json:"IdCountry"`
	Type          int                          `json:"TeamType"`
	AgeType       int                          `json:"AgeType"`
	Tactics       string                       `json:"Tactics"`
	Name          []DefaultDescriptionResponse `json:"TeamName"`
	Abbreviation  string                       `json:"Abbreviation"`
	Coaches       []CoachResponse              `json:"Coaches"`
	Players       []MatchPlayerResponse        `json:"Players"`
	Bookings      []BookingResponse            `json:"Bookings"`
	Goals         []GoalResponse               `json:"Goals"`
	Substitutions []SubstitutionResponse       `json:"Substitutions"`
	FootballType  int                          `json:"FootballType"`
	Gender        Gender                       `json:"Gender"`
	AssociationId string                       `json:"IdAssociation"`
}

type WeatherResponse struct {
	HumIdity      string                       `json:"HumIdity"`
	Temperature   string                       `json:"Temperature"`
	WindSpeed     string                       `json:"WindSpeed"`
	Type          int                          `json:"Type"`
	TypeLocalized []DefaultDescriptionResponse `json:"TypeLocalized"`
}

type BallPossessionResponse struct {
	Intervals   []interface{} `json:"Intervals"`
	LastX       []interface{} `json:"LastX"`
	OverallHome float32       `json:"OverallHome"`
	OverallAway float32       `json:"OverallAway"`
}

type MatchPlayerResponse struct {
	Id            string                       `json:"IdPlayer"`
	TeamId        string                       `json:"IdTeam"`
	ShirtNumber   int                          `json:"ShirtNumber"`
	Status        int                          `json:"Status"` // TODO: Enum
	SpecialStatus string                       `json:"SpecialStatus"`
	IsCaptain     bool                         `json:"Captain"`
	Name          []DefaultDescriptionResponse `json:"PlayerName"`
	ShortName     []DefaultDescriptionResponse `json:"ShortName"`
	Position      int                          `json:"Position"`    // TODO: Enum
	FieldStatus   int                          `json:"FieldStatus"` // TODO: Enum
	LineupX       string                       `json:"LineupX"`
	LineupY       string                       `json:"LineupY"`
}

type CoachResponse struct {
	Id            string                       `json:"IdCoach"`
	CountryId     string                       `json:"IdCountry"`
	Name          []DefaultDescriptionResponse `json:"Name"`
	Alias         []DefaultDescriptionResponse `json:"Alias"`
	Role          int                          `json:"Role"` // TODO: Enum
	SpecialStatus string                       `json:"SpecialStatus"`
}

type BookingResponse struct {
	Card        int        `json:"Card"`
	Period      PeriodEnum `json:"Period"`
	EventId     string     `json:"IdEvent"`
	EventNumber string     `json:"EventNumber"`
	PlayerId    string     `json:"IdPlayer"`
	CoachId     string     `json:"IdCoach"`
	TeamId      string     `json:"IdTeam"`
	Minute      string     `json:"Minute"`
	Reason      string     `json:"Reason"`
}

type SubstitutionResponse struct {
	EventId       string                       `json:"IdEvent"`
	Period        PeriodEnum                   `json:"Period"`
	Reason        int                          `json:"Reason"`            // TODO: Enum
	Position      int                          `json:"SubstituePosition"` // TODO: Enum
	PlayerOffId   string                       `json:"IdPlayerOff"`
	PlayerOnId    string                       `json:"IdPlayerOn"`
	PlayerOffName []DefaultDescriptionResponse `json:"PlayerOffName"`
	PlayerOnName  []DefaultDescriptionResponse `json:"PlayerOnName"`
	Minute        string                       `json:"Minute"`
	TeamId        string                       `json:"TeamId"`
}

type OfficialResponse struct {
	Id            string                       `json:"OfficialsId"`
	CountryId     string                       `json:"IdCountry"`
	Name          []DefaultDescriptionResponse `json:"Name"`
	ShortName     []DefaultDescriptionResponse `json:"ShortName"`
	Type          int                          `json:"OfficialType"` // TODO: Enum
	TypeLocalized []DefaultDescriptionResponse `json:"TypeLocalized"`
}

type GoalResponse struct {
	Id             string     `json:"IdGoal"`
	TeamId         string     `json:"IdTeam"`
	Type           int        `json:"Type"` // TODO: Enum
	PlayerId       string     `json:"IdPlayer"`
	Minute         string     `json:"Minute"`
	AssistPlayerId string     `json:"IdAssistPlayer"`
	Period         PeriodEnum `json:"Period"`
}

type CompetitionResponse struct {
	CompetitionId       string                       `json:"IdCompetition"`
	Name                []DefaultDescriptionResponse `json:"Name"`
	ConfederationId     []string                     `json:"IdConfederation"`
	MemberAssociationId []string                     `json:"IdMemberAssociation"`
	OwnerId             string                       `json:"IdOwner"`
	Gender              Gender                       `json:"Gender"`
	FootballType        int                          `json:"FootballType"`    // TODO: Enum
	TeamType            int                          `json:"TeamType"`        // TODO: Enum
	Type                int                          `json:"CompetitionType"` // TODO: Enum
	Properties          interface{}                  `json:"Properties"`
	IsUpdateable        bool                         `json:"IsUpdateable"`
}

type PlayerResponse struct {
	Id                       string                       `json:"IdPlayer"`
	Name                     []DefaultDescriptionResponse `json:"Name"`
	Alias                    []DefaultDescriptionResponse `json:"Alias"`
	Birthdate                time.Time                    `json:"Birthdate"`
	Weight                   float32                      `json:"Weight"`
	Height                   float32                      `json:"Height"`
	Birthplace               string                       `json:"BirthPlace"`
	CountryId                string                       `json:"IdCountry"`
	InternationalCaps        int                          `json:"InternationalCaps"`
	InternationalDebut       int                          `json:"InternationalDebut"`
	TopCompetitionDebut      int                          `json:"TopCompetitionDebut"`
	PictureURL               string                       `json:"PictureUrl"`
	ThumbnailURL             string                       `json:"ThumbnailUrl"`
	TwitterAccount           string                       `json:"TwitterAccount"`
	PreferredFoot            string                       `json:"PreferredFoot"`
	MediaContent             []string                     `json:"MediaContent"`
	LocalizedTwitterAccounts []DefaultDescriptionResponse `json:"LocalizedTwitterAccounts"`
	Goals                    int                          `json:"Goals"`
	Properties               interface{}                  `json:"Properties"`
	IsUpdateable             bool                         `json:"IsUpdateable"`
}

type SeasonResponse struct {
	Id                  string                       `json:"IdSeason"`
	Name                []DefaultDescriptionResponse `json:"name"`
	ShortName           []DefaultDescriptionResponse `json:"ShortName"`
	Abbreviation        string                       `json:"abbreviation"`
	MemberAssociations  []string                     `json:"IdMemberAssocation"`
	Confederations      []string                     `json:"IdConfederation"`
	CompetitionId       string                       `json:"IdCompetition"`
	StartDate           time.Time                    `json:"StartDate"`
	EndDate             time.Time                    `json:"EndDate"`
	PictureURL          string                       `json:"PictureUrl"`
	MascotPictureURL    string                       `json:"MascotPictureUrl"`
	MatchBallPictureURL string                       `json:"MatchBallPictureUrl"`
	HostTeams           interface{}                  `json:"HostTeams"`
	SportType           int                          `json:"SportType"`
	Properties          SeasonProperties             `json:"Properties"`
	IsUpdateable        bool                         `json:"IsUpdateable"`
}

type SeasonProperties struct {
	InfostradaId string `json:"IdInfostrada"`
	Providers    string `json:"ProvIders"`
}

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
	GoalScore        MatchEvent = 0
	YellowCard       MatchEvent = 2
	RedCard          MatchEvent = 3
	DoubleYellow     MatchEvent = 4
	Substitution     MatchEvent = 5
	PenaltyAwarded   MatchEvent = 6
	MatchStart       MatchEvent = 7
	HalfEnd          MatchEvent = 8
	MatchPaused      MatchEvent = 9
	MatchResumed     MatchEvent = 10
	GoalAttempt      MatchEvent = 12
	FoulUnknown      MatchEvent = 14
	Offside          MatchEvent = 15
	CornerKick       MatchEvent = 16
	ShotBlocked      MatchEvent = 17
	Foul             MatchEvent = 18
	CoinToss         MatchEvent = 19
	Unknown3         MatchEvent = 20
	ThrowIn          MatchEvent = 24
	Clearance        MatchEvent = 25
	MatchEnd         MatchEvent = 26
	Unknown2         MatchEvent = 27
	Crossbar         MatchEvent = 32
	Crossbar2        MatchEvent = 33
	OwnGoal          MatchEvent = 34
	HandBall         MatchEvent = 37
	FreeKickGoal     MatchEvent = 39
	PenaltyGoal      MatchEvent = 41
	FreeKickCrossbar MatchEvent = 44
	FreeKickPost     MatchEvent = 49
	PenaltyMissed    MatchEvent = 60
	PenaltyMissed2   MatchEvent = 65
	GoalieSaved      MatchEvent = 57
	VARPenalty       MatchEvent = 72
	Unknown          MatchEvent = 9999
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
	Capacity           int                          `json:"Capacity"`
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
	Latitude           float64                      `json:"Latitude"`
	Longitude          float64                      `json:"Longitude"`
	Length             string                       `json:"Length"`
	Width              string                       `json:"Width"`
	Properties         interface{}                  `json:"Properties"`
	IsUpdateable       bool                         `json:"IsUpdateable"`
}

type TeamResponse struct {
	Score         int                          `json:"Score"`
	Side          string                       `json:"Side"`
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
	Humidity      string                       `json:"Humidity"`
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
	Status        int                          `json:"Status"`        // TODO: Enum
	SpecialStatus *int                         `json:"SpecialStatus"` // TODO: Evaluate
	IsCaptain     bool                         `json:"Captain"`
	Name          []DefaultDescriptionResponse `json:"PlayerName"`
	ShortName     []DefaultDescriptionResponse `json:"ShortName"`
	Position      int                          `json:"Position"`    // TODO: Enum
	FieldStatus   int                          `json:"FieldStatus"` // TODO: Enum
	LineupX       *float64                     `json:"LineupX"`     // TODO: Evaluate
	LineupY       *float64                     `json:"LineupY"`     // TODO: Evaluate
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

type StandingsMatchResult struct {
	MatchId   string    `json:"IdMatch"`
	StartTime time.Time `json:"StartTime"`
	Result    int       `json:"Result"`
	GroupId   string    `json:"IdGroup"`
	StageId   string    `json:"IdStage"`
}

type StandingsProperties struct {
	InfostradaId string `json:"IdInfostrada"`
}

type StandingResponse struct {
	Results []StandingsResult `json:"Results"`
}

type StandingsResult struct {
	MatchDay            int                    `json:"MatchDay"`
	CompetitionId       string                 `json:"IdCompetition"`
	SeasonId            string                 `json:"IdSeason"`
	GroupId             string                 `json:"IdGroup"`
	Date                time.Time              `json:"Date"`
	Group               interface{}            `json:"Group"`
	Won                 int                    `json:"Won"`
	Lost                int                    `json:"Lost"`
	Drawn               int                    `json:"Drawn"`
	Played              int                    `json:"Played"`
	HomeWon             int                    `json:"HomeWon"`
	HomeLost            int                    `json:"HomeLost"`
	HomeDrawn           int                    `json:"HomeDrawn"`
	HomePlayed          int                    `json:"HomePlayed"`
	AwayWon             int                    `json:"AwayWon"`
	AwayLost            int                    `json:"AwayLost"`
	AwayDrawn           int                    `json:"AwayDrawn"`
	AwayPlayed          int                    `json:"AwayPlayed"`
	Against             int                    `json:"Against"`
	For                 int                    `json:"For"`
	HomeAgainst         int                    `json:"HomeAgainst"`
	HomeFor             int                    `json:"HomeFor"`
	AwayAgainst         int                    `json:"AwayAgainst"`
	AwayFor             int                    `json:"AwayFor"`
	Position            int                    `json:"Position"`
	HomePosition        int                    `json:"HomePosition"`
	AwayPosition        int                    `json:"AwayPosition"`
	Points              int                    `json:"Points"`
	HomePoints          int                    `json:"HomePoints"`
	AwayPoints          int                    `json:"AwayPoints"`
	PreviousPosition    int                    `json:"PreviousPosition"`
	GoalsDifference     int                    `json:"GoalsDiference"` // Typo is expected
	Team                TeamResponse           `json:"Team"`
	StartDate           time.Time              `json:"StartDate"`
	EndDate             time.Time              `json:"EndDate"`
	FairPlayCoefficient float32                `json:"FairPlayCoefficient"`
	WinByExtraTime      int                    `json:"WinByExtraTime"`
	WinByPenalty        int                    `json:"WinByPenalty"`
	MatchResults        []StandingsMatchResult `json:"MatchResults"`
	Properties          StandingsProperties    `json:"Properties"`
	IsUpdateable        bool                   `json:"IsUpdateable"`
}

type MatchDataResponse struct {
	MatchId                   string                       `json:"IdMatch"`
	StageId                   string                       `json:"IdStage"`
	GroupId                   string                       `json:"IdGroup"`
	SeasonId                  string                       `json:"IdSeason"`
	CompetitionId             string                       `json:"IdCompetition"`
	CompetitionName           []DefaultDescriptionResponse `json:"CompetitionName"`
	SeasonName                []DefaultDescriptionResponse `json:"SeasonName"`
	SeasonShortName           []DefaultDescriptionResponse `json:"SeasonShortName"`
	Stadium                   StadiumResponse              `json:"Stadium"`
	ResultType                int                          `json:"ResultType"`
	MatchDay                  interface{}                  `json:"MatchDay"`
	HomeTeamPenaltyScore      int                          `json:"HomeTeamPenaltyScore"`
	AwayTeamPenaltyScore      int                          `json:"AwayTeamPenaltyScore"`
	AggregateHomeTeamScore    int                          `json:"AggregateHomeTeamScore"`
	AggregateAwayTeamScore    int                          `json:"AggregateAwayTeamScore"`
	Weather                   interface{}                  `json:"Weather"`
	Attendance                string                       `json:"Attendance"`
	Date                      time.Time                    `json:"Date"`
	LocalDate                 time.Time                    `json:"LocalDate"`
	MatchTime                 string                       `json:"MatchTime"`
	SecondHalfTime            interface{}                  `json:"SecondHalfTime"`
	FirstHalfTime             interface{}                  `json:"FirstHalfTime"`
	FirstHalfExtraTime        int                          `json:"FirstHalfExtraTime"`
	SecondHalfExtraTime       int                          `json:"SecondHalfExtraTime"`
	Winner                    interface{}                  `json:"Winner"`
	Period                    int                          `json:"Period"`
	HomeTeam                  TeamResponse                 `json:"HomeTeam"`
	AwayTeam                  TeamResponse                 `json:"AwayTeam"`
	BallPossession            BallPossessionResponse       `json:"BallPossession"`
	TerritorialPossesion      interface{}                  `json:"TerritorialPossesion"`
	TerritorialThirdPossesion interface{}                  `json:"TerritorialThirdPossesion"`
	Officials                 []OfficialResponse           `json:"Officials"`
	MatchStatus               int                          `json:"MatchStatus"`
	GroupName                 []interface{}                `json:"GroupName"`
	StageName                 []DefaultDescriptionResponse `json:"StageName"`
	OfficialityStatus         int                          `json:"OfficialityStatus"`
	TimeDefined               bool                         `json:"TimeDefined"`
	Properties                StatsResponse                `json:"Properties"`
	IsUpdateable              interface{}                  `json:"IsUpdateable"`
}

type StatsResponse struct {
	StatsPerformId string `json:"IdStatsPerform"`
}

type VarNotificationDataResponse struct {
	Incident int `json:"Incident"`
	Reason   int `json:"Reason"`
	Status   int `json:"Status"`
	Result   int `json:"Result"`
}

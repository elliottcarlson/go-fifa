package go_fifa

import (
	"time"

	"github.com/google/uuid"
)

type Gender int // Represents the gender of the player

const (
	MaleGender   Gender = 1
	FemaleGender Gender = 2
)

type PeriodEnum int // Represents the period of play

const (
	FirstPeriod       PeriodEnum = 3
	SecondPeriod      PeriodEnum = 5
	FirstExtraPeriod  PeriodEnum = 7
	SecondExtraPeriod PeriodEnum = 9
	ShootoutPeriod    PeriodEnum = 11
)

type MatchEvent int // Represents the type of event that occured

const (
	GoalScore        MatchEvent = 0    // A player scored a goal
	Assist           MatchEvent = 1    // A player assisted a goal
	YellowCard       MatchEvent = 2    // A player was given a yellow card
	RedCard          MatchEvent = 3    // A player was given a red card
	DoubleYellow     MatchEvent = 4    // A player was given a second yellow card
	Substitution     MatchEvent = 5    // A player was substituded out
	PenaltyAwarded   MatchEvent = 6    // A penalty was awarded
	MatchStart       MatchEvent = 7    // The match has started
	HalfEnd          MatchEvent = 8    // The half is over
	MatchPaused      MatchEvent = 9    // The referee paused the match
	MatchResumed     MatchEvent = 10   // The match has resumed
	GoalAttempt      MatchEvent = 12   // A player made an attempt on goal
	FoulUnknown      MatchEvent = 14   // Unknown event type
	Offside          MatchEvent = 15   // A player was deemed offside
	CornerKick       MatchEvent = 16   // Referee awarded a corner kick
	ShotBlocked      MatchEvent = 17   // The goalie blocked a shot
	Foul             MatchEvent = 18   // A player commited a foul
	CoinToss         MatchEvent = 19   // The coin toss happened
	Unknown3         MatchEvent = 20   // Unknown event type
	DroppedBall      MatchEvent = 23   // The referee called a drop ball
	ThrowIn          MatchEvent = 24   // A player took a throw-in
	Clearance        MatchEvent = 25   // A player cleared the ball
	MatchEnd         MatchEvent = 26   // The match is over
	Unknown2         MatchEvent = 27   // Unknown event type
	HitCrossbar      MatchEvent = 32   // The player hit the crossbar
	HitPost          MatchEvent = 33   // The player hit the post
	OwnGoal          MatchEvent = 34   // A player scored against their own team
	HandBall         MatchEvent = 37   // The referee called a handball
	FreeKickGoal     MatchEvent = 39   // A player scored a goal from a free-kick
	PenaltyGoal      MatchEvent = 41   // A player scored a goal from a penalty
	FreeKickCrossbar MatchEvent = 44   // A player hit the crossbar from a free-kick
	FreeKickPost     MatchEvent = 49   // A player hit the post from a free-kick
	PenaltyMissed    MatchEvent = 60   // A player missed their penalty
	PenaltyMissed2   MatchEvent = 65   // A player missed their penalty??
	GoalieSaved      MatchEvent = 57   // The goalie stopped the shot
	VARPenalty       MatchEvent = 72   // A penalty was awarded after VAR
	Pending          MatchEvent = 9999 // This event is still pending a final event type, query again for the new event
)

// PagindatedResponse represents the options for pagination results from the API
type PaginatedResponse struct {
	ContinuationToken string `json:"ContinuationToken"`
	ContinuationHash  string `json:"ContinuationHash"`
}

type GetMatchesResponse struct {
	PaginatedResponse
	Results []MatchResponse
}

// MatchResponse represents the response from calling an API to get match data
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
	HomeTeamScore             int                          `json:"HomeTeamScore"`
	AwayTeamScore             int                          `json:"AwayTeamScore"`
	HomeTeamPenaltyScore      int                          `json:"HomeTeamPenaltyScore"`
	AwayTeamPenaltyScore      int                          `json:"AwayTeamPenaltyScore"`
	AggregateHomeTeamScore    int                          `json:"AggregateHomeTeamScore"`
	AggregateAwayTeamScore    int                          `json:"AggregateAwayTeamScore"`
	LastPeriodUpdate          interface{}                  `json:"LastPeriodUpdate"`
	Leg                       interface{}                  `json:"Leg"`
	IsHomeMatch               bool                         `json:"IsHomeMatch"`
	Weather                   WeatherResponse              `json:"Weather"`
	Date                      time.Time                    `json:"Date"`
	LocalDate                 time.Time                    `json:"LocalDate"`
	IsTicketSalesAllowed      bool                         `json:"IsTicketSalesAllowed"`
	MatchTime                 string                       `json:"MatchTime"`
	SecondHalfTime            string                       `json:"SecondHalfTime"`
	FirstHalfTime             string                       `json:"FirstHalfTime"`
	FirstHalfExtraTime        int                          `json:"FirstHalfExtraTime"`
	SecondHalfExtraTime       int                          `json:"SecondHalfExtraTime"`
	WinnerId                  string                       `json:"Winner"`
	MatchReportUrl            string                       `json:"MatchReportUrl"`
	PlaceHolderA              string                       `json:"PlaceHolderA"`
	PlaceHolderB              string                       `json:"PlaceHolderB"`
	Period                    PeriodEnum                   `json:"Period"`
	HomeTeam                  MatchTeamResponse            `json:"Home"`
	AwayTeam                  MatchTeamResponse            `json:"Away"`
	BallPossession            BallPossessionResponse       `json:"BallPossession"`
	TerritorialPossesion      string                       `json:"TerritorialPossesion"`
	TerritorialThirdPossesion string                       `json:"TerritorialThirdPossesion"`
	Officials                 []OfficialResponse           `json:"Officials"`
	Status                    int                          `json:"MatchStatus"`
	GroupName                 []DefaultDescriptionResponse `json:"GroupName"`
	StageName                 []DefaultDescriptionResponse `json:"StageName"`
	OfficialityStatus         int                          `json:"OfficialityStatus"`
	MatchNumber               int                          `json:"MatchNumber"`
	MatchLegInfo              interface{}                  `json:"MatchLegInfo"`
	TimeDefined               bool                         `json:"TimeDefined"`
	Properties                interface{}                  `json:"Properties"`
	IsUpdateable              bool                         `json:"IsUpdateable"`
}

// DefaultDescriptionResponse represents a common usage for specifying
// locale information for different nested structures
type DefaultDescriptionResponse struct {
	Locale      string `json:"Locale"`
	Description string `json:"Description"`
}

// StadiumResponse represents the information about a stadium
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

// MatchTeamResponse represents the information about a team during a match
type MatchTeamResponse struct {
	Score         int                          `json:"Score"`
	Confederation string                       `json:"IdConfederation"`
	Side          string                       `json:"Side"`
	Id            string                       `json:"IdTeam"`
	PictureURL    string                       `json:"PictureUrl"`
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
	ShortClubName string                       `json:"ShortClubName"`
}

// TeamResponse represents the structure of a team
type TeamResponse struct {
	Id             string                       `json:"IdTeam"`
	Confederation  string                       `json:"IdConfederation"`
	Type           int                          `json:"Type"`
	AgeType        int                          `json:"AgeType"`
	FootballType   int                          `json:"FootballType"`
	Gender         Gender                       `json:"Gender"`
	Name           []DefaultDescriptionResponse `json:"Name"`
	AssociationId  string                       `json:"IdAssociation"`
	CityId         any                          `json:"city"`
	Headquarters   any                          `json:"Headquarters"`
	TrainingCenter any                          `json:"TrainingCentre"`
	OfficialSite   any                          `json:"OfficialSite"`
	City           string                       `json:"City"`
	CountryId      string                       `json:"IdCountry"`
	PostalCode     string                       `json:"PostalCode"`
	RegionName     any                          `json:"RegionName"`
	ShortClubName  string                       `json:"ShortClubName"`
	Abbreviation   string                       `json:"Abbreviation"`
	Street         string                       `json:"Street"`
	FoundationYear any                          `json:"FoundationYear"`
	Stadium        any                          `json:"Stadium"`
	PictureUrl     string                       `json:"PictureUrl"`
	ThumbnailUrl   string                       `json:"ThumbnailUrl"`
	DisplayName    []DefaultDescriptionResponse `json:"DisplayName"`
	Content        any                          `json:"Content"`
	Properties     Properties                   `json:"Properties"`
	IsUpdateable   bool                         `json:"IsUpdateable"`
}

// WeatherRespones represents the information about weather
type WeatherResponse struct {
	Humidity      string                       `json:"Humidity"`
	Temperature   string                       `json:"Temperature"`
	WindSpeed     string                       `json:"WindSpeed"`
	Type          int                          `json:"Type"`
	TypeLocalized []DefaultDescriptionResponse `json:"TypeLocalized"`
}

// BallPossessionResponse represents the information about ball possession
type BallPossessionResponse struct {
	Intervals   []interface{} `json:"Intervals"`
	LastX       []interface{} `json:"LastX"`
	OverallHome float32       `json:"OverallHome"`
	OverallAway float32       `json:"OverallAway"`
}

// MatchPlayerResponse represents the information about a player
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

// CoachResponse represents the information about a coach
type CoachResponse struct {
	Id            string                       `json:"IdCoach"`
	CountryId     string                       `json:"IdCountry"`
	Name          []DefaultDescriptionResponse `json:"Name"`
	Alias         []DefaultDescriptionResponse `json:"Alias"`
	Role          int                          `json:"Role"` // TODO: Enum
	SpecialStatus string                       `json:"SpecialStatus"`
}

// BookingResponse represents the information about a booking
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

// SubstitutionResponse represents the information about a substitution
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

// OfficialResponse represents the information about an offical
type OfficialResponse struct {
	Id            string                       `json:"OfficialId"`
	CountryId     string                       `json:"IdCountry"`
	Name          []DefaultDescriptionResponse `json:"Name"`
	ShortName     []DefaultDescriptionResponse `json:"NameShort"`
	Type          int                          `json:"OfficialType"` // TODO: Enum
	TypeLocalized []DefaultDescriptionResponse `json:"TypeLocalized"`
}

// GoalResponse represents the information about a goal
type GoalResponse struct {
	Id             string     `json:"IdGoal"`
	TeamId         string     `json:"IdTeam"`
	Type           int        `json:"Type"` // TODO: Enum
	PlayerId       string     `json:"IdPlayer"`
	Minute         string     `json:"Minute"`
	AssistPlayerId string     `json:"IdAssistPlayer"`
	Period         PeriodEnum `json:"Period"`
}

// CompetitionResponse represents the information about a competition
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

// PlayerResponse represents the information about a player
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
	Picture                  interface{}                  `json:"PlayerPicture"`
}

// SeasonResponse represents the information about a season
type SeasonResponse struct {
	Id                  string                       `json:"IdSeason"`
	Name                []DefaultDescriptionResponse `json:"name"`
	ShortName           []DefaultDescriptionResponse `json:"ShortName"`
	Abbreviation        string                       `json:"abbreviation"`
	MemberAssociations  []string                     `json:"IdMemberAssociation"`
	Confederations      []string                     `json:"IdConfederation"`
	CompetitionId       string                       `json:"IdCompetition"`
	StartDate           time.Time                    `json:"StartDate"`
	EndDate             time.Time                    `json:"EndDate"`
	PictureURL          string                       `json:"PictureUrl"`
	MascotPictureURL    string                       `json:"MascotPictureUrl"`
	MatchBallPictureURL string                       `json:"MatchBallPictureUrl"`
	HostTeams           interface{}                  `json:"HostTeams"`
	SportType           int                          `json:"SportType"`
	Properties          Properties                   `json:"Properties"`
	IsUpdateable        bool                         `json:"IsUpdateable"`
}

// StandingsMatchResult represents the information about standings
type StandingsMatchResult struct {
	MatchId   string    `json:"IdMatch"`
	StartTime time.Time `json:"StartTime"`
	Result    int       `json:"Result"`
	GroupId   string    `json:"IdGroup"`
	StageId   string    `json:"IdStage"`
}

// StandingResponse represents the API response about standings
type StandingResponse struct {
	Results []StandingsResult `json:"Results"`
}

// StandingsSesult represents the information about standings
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
	Properties          Properties             `json:"Properties"`
	IsUpdateable        bool                   `json:"IsUpdateable"`
}

// MatchDataResponse represents the response about match events from the API
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
	HomeTeam                  MatchTeamResponse            `json:"HomeTeam"`
	AwayTeam                  MatchTeamResponse            `json:"AwayTeam"`
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

// StatsResponse represents the stats response from the API
type StatsResponse struct {
	StatsPerformId string `json:"IdStatsPerform"`
}

// VarNotificationDataResponse struct represents information about VAR
type VarNotificationDataResponse struct {
	Incident int `json:"Incident"`
	Reason   int `json:"Reason"`
	Status   int `json:"Status"`
	Result   int `json:"Result"`
}

// CurrentMatchesResponse represents the response from getting current matches
type CurrentMatchesResponse struct {
	PaginatedResponse
	Results []MatchDataResponse `json:"Results"`
}

// GetMatchEventsResponse represents the response returned when querying event data
type GetMatchEventsResponse struct {
	StageId       string          `json:"IdStage"`
	MatchId       string          `json:"IdMatch"`
	CompetitionId string          `json:"IdCompetition"`
	SeasonId      string          `json:"IdSeason"`
	GroupId       string          `json:"IdGroup"`
	Events        []EventResponse `json:"Event"`
	Properties    interface{}     `json:"Properties"`
	IsUpdateable  bool            `json:"IsUpdateable"`
}

// EventResponse represents the structure of an event that happened during a match
type EventResponse struct {
	Context             string                       `json:"Context"`
	Id                  string                       `json:"EventId"`
	GuId                uuid.UUID                    `json:"GuId"`
	TeamId              string                       `json:"IdTeam"`
	PlayerId            string                       `json:"IdPlayer"`
	SubPlayerId         string                       `json:"IdSubPlayer"`
	PersonId            string                       `json:"IdPerson"`
	SubTeamId           string                       `json:"IdSubTeam"`
	Timestamp           time.Time                    `json:"Timestamp"`
	DateTimeUTC         time.Time                    `json:"DateTimeUTC"`
	MatchMinute         string                       `json:"MatchMinute"`
	Period              PeriodEnum                   `json:"Period"`
	HomeGoals           int                          `json:"HomeGoals"`
	AwayGoals           int                          `json:"AwayGoals"`
	Type                MatchEvent                   `json:"Type"`
	TypeLocalized       []DefaultDescriptionResponse `json:"TypeLocalized"`
	PositionX           float32                      `json:"PositionX"`
	PositionY           float32                      `json:"PositionY"`
	GoalGatePositionX   float32                      `json:"GoalGatePositionX"`
	GoalGatePositionY   float32                      `json:"GoalGatePositionY"`
	GoalGatePositionZ   float32                      `json:"GoalGatePositionZ"`
	VarDetail           string                       `json:"VarDetail"`
	VarNotificationData VarNotificationDataResponse  `json:"VarNotificationData"`
	HomePenaltyGoals    int                          `json:"HomePenaltyGoals"`
	AwayPenaltyGoals    int                          `json:"AwayPenaltyGoals"`
	EventDescription    []DefaultDescriptionResponse `json:"EventDescription"`
}

// GetCompetitionsResponse represents the object returned when querying a competition
type GetCompetitionsResponse struct {
	PaginatedResponse
	Results []CompetitionResponse `json:"Results"`
}

type Properties struct {
	IdIFES             string `json:"IdIFES,omitempty"`
	StatsPerformIfesId string `json:"StatsPerformIfesId,omitempty"`
	IdStatsPerform     string `json:"IdStatsPerform,omitempty"`
	Providers          string `json:"Providers,omitempty"`
}

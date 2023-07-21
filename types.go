package go_fifa

import "time"

// Represents the gender of the player
type Gender int

const (
	MaleGender   Gender = 1 // Male
	FemaleGender Gender = 2 // Female
)

// Represents the period of play
type Period int

const (
	NotStarted        Period = 0  // The first period hasn't started yet
	FirstPeriod       Period = 3  // First period
	HalfTime          Period = 4  // Half Time
	SecondPeriod      Period = 5  // Second period
	FirstExtraPeriod  Period = 7  // First extra period
	SecondExtraPeriod Period = 9  // Second extra period
	MatchOver         Period = 10 // The match is over
	ShootoutPeriod    Period = 11 // Shootout
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

type OfficialType int

const (
	Referee                 OfficialType = 1
	AssistantReferee1       OfficialType = 2
	AssistantReferee2       OfficialType = 3
	FourthOfficial          OfficialType = 4
	VideoAssistantReferee   OfficialType = 5
	ReserveReferee          OfficialType = 6
	OffsideVAR              OfficialType = 7
	AssistantVAR            OfficialType = 8
	SupportVAR              OfficialType = 9
	ReserveAssistantReferee OfficialType = 10
)

type CoachRole int

const (
	ManagerRole      CoachRole = 0
	AssistantManager CoachRole = 1
)

type PaginationResponse struct {
	ContinuationToken string `json:"ContinuationToken"`
	ContinuationHash  string `json:"ContinuationHash"`
}

// CompetitionsResponse represents the response from the competitions API
type CompetitionsResponse struct {
	PaginationResponse
	Results []Competition `json:"Results"`
}

// Competition represents the structure of a competition
type Competition struct {
	Id                   string              `json:"IdCompetition"`
	Name                 []LocaleDescription `json:"Name"`
	ConfederationIds     []string            `json:"IdConfederation"`
	MemberAssociationIds []string            `json:"IdMemberAssociation"`
	OwnerId              string              `json:"IdOwner"`
	Gender               Gender              `json:"Gender"`
	FootballType         int                 `json:"FootballType"`    // TODO: Enum
	TeamType             int                 `json:"TeamType"`        // TODO: Enum
	Type                 int                 `json:"CompetitionType"` // TODO: Enum
	AgeType              int                 `json:"AgeType"`         // TODO: Enum
	Properties           Properties          `json:"Properties"`
	IsUpdateable         bool                `json:"IsUpdateable"`
}

// Properties represents the properties that can be attached to many objects
type Properties struct {
	IdIFES             string `json:"IdIFES,omitempty"`
	StatsPerformIfesId string `json:"StatsPerformIfesId,omitempty"`
	IdStatsPerform     string `json:"IdStatsPerform,omitempty"`
	Providers          string `json:"Providers,omitempty"`
}

// LocaleDescription represents a common usage for specifying
// locale information for different nested structures
type LocaleDescription struct {
	Locale      string `json:"Locale"`
	Description string `json:"Description"`
}

// TeamResponse represents the structure returned from the teams API
type TeamResponse struct {
	Id              string              `json:"IdTeam"`
	ConfederationId string              `json:"IdConfederation"`
	Type            int                 `json:"Type"`         // TODO: Enum
	AgeType         int                 `json:"AgeType"`      // TODO: Enum
	FootballType    int                 `json:"FootballType"` // TODO: Enum
	Gender          Gender              `json:"Gender"`       // TODO: Enum
	Name            []LocaleDescription `json:"Name"`
	AssociationId   string              `json:"IdAssociation"`
	CityId          string              `json:"IdCity"`
	Headquarters    any                 `json:"Headquarters"`
	TrainingCenter  any                 `json:"TrainingCentre"`
	OfficialSite    any                 `json:"OfficialSite"`
	City            string              `json:"City"`
	CountryId       string              `json:"IdCountry"`
	PostalCode      string              `json:"PostalCode"`
	RegionName      any                 `json:"RegionName"`
	ShortClubName   string              `json:"ShortClubName"`
	Abbreviation    string              `json:"Abbreviation"`
	Street          string              `json:"Street"`
	FoundationYear  int                 `json:"FoundationYear"`
	Stadium         any                 `json:"Stadium"`
	PictureUrl      string              `json:"PictureUrl"`
	DisplayName     []LocaleDescription `json:"DisplayName"`
	Content         any                 `json:"Content"`
	Properties      Properties          `json:"Properties"`
	IsUpdateable    bool                `json:"IsUpdateable"`
}

// SeasonResponse represents the structure returned from the Seasons API
type SeasonResponse struct {
	Id                   string              `json:"IdSeason"`
	Name                 []LocaleDescription `json:"Name"`
	ShortName            []string            `json:"ShortName"`
	Abbreviation         string              `json:"Abbreviation"`
	MemberAssociationIds []string            `json:"IdMemberAssociation"`
	ConfederationIds     []string            `json:"IdConfederation"`
	CompetitionId        string              `json:"IdCompetition"`
	StartDate            time.Time           `json:"StartDate"`
	EndDate              time.Time           `json:"EndDate"`
	PictureUrl           string              `json:"PictureUrl"`
	MascotPictureUrl     string              `json:"MascotPictureUrl"`
	MatchBallPictureUrl  string              `json:"MatchBallPictureUrl"`
	HostTeams            []SeasonHostTeam    `json:"HostTeams"`
	SportType            int                 `json:"SportType"`
	Properties           Properties          `json:"Properties"`
	IsUpdateable         bool                `json:"IsUpdateable"`
}

// HostTeam represents the structure of a host team
type SeasonHostTeam struct {
	TeamId string `json:"IdTeam"`
}

// PlayerResponse represents the structure of a player from the API
type PlayerResponse struct {
	Id                       string              `json:"IdPlayer"`
	Name                     []LocaleDescription `json:"Name"`
	Alias                    []LocaleDescription `json:"Alias"`
	BirthDate                time.Time           `json:"BirthDate"`
	Weight                   int                 `json:"Weight"`
	Height                   int                 `json:"Height"`
	Birthplace               string              `json:"BirthPlace"`
	CountryId                string              `json:"IdCountry"`
	InternationalCaps        int                 `json:"InternationalCaps"`
	InternationalDebut       any                 `json:"InternationalDebut"`
	TopCompetitionDebut      any                 `json:"TopCompetitionDebut"`
	PictureUrl               string              `json:"PictureUrl"`
	ThumbnailUrl             string              `json:"ThumbnailUrl"`
	TwitterAccount           any                 `json:"TwitterAccount"`
	PreferredFoot            any                 `json:"PreferredFoot"`
	MediaContent             any                 `json:"MediaContent"`
	LocalizedTwitterAccounts any                 `json:"LocalizedTwitterAccounts"`
	Goals                    int                 `json:"Goals"`
	PlayerPicture            any                 `json:"PlayerPicture"`
	Properties               Properties          `json:"Properties"`
	IsUpdateable             bool                `json:"IsUpdateable"`
}

// LiveResponse represents the structure of a response from the "now" api
type LiveResponse struct {
	PaginationResponse
	Results []LiveMatch `json:"Results"`
}

// LiveMatch represents the structure of a live match from the "now" api
type LiveMatch struct {
	MatchId                   string              `json:"IdMatch"`
	StageId                   string              `json:"IdStage"`
	GroupId                   string              `json:"IdGroup"`
	SeasonId                  string              `json:"IdSeason"`
	CompetitionId             string              `json:"IdCompetition"`
	CompetitionName           []LocaleDescription `json:"CompetitionName"`
	SeasonName                []LocaleDescription `json:"SeasonName"`
	SeasonShortName           any                 `json:"SeasonShortName"`
	Stadium                   Stadium             `json:"Stadium"`
	ResultType                int                 `json:"ResultType"` // TODO: Enum
	MatchDay                  string              `json:"MatchDay"`
	HomeTeamPenaltyScore      int                 `json:"HomeTeamPenaltyScore"`
	AwayTeamPenaltyScore      int                 `json:"AwayTeamPenaltyScore"`
	AggregateHomeTeamScore    any                 `json:"AggregateHomeTeamScore"`
	AggregateAwayTeamScore    any                 `json:"AggregateAwayTeamScore"`
	Weather                   Weather             `json:"Weather"`
	Attendance                any                 `json:"Attendance"`
	Date                      time.Time           `json:"Date"`
	LocalDate                 time.Time           `json:"LocalDate"`
	MatchTime                 string              `json:"MatchTime"`
	SecondHalfTime            any                 `json:"SecondHalfTime"`
	FirstHalfTime             any                 `json:"FirstHalfTime"`
	FirstHalfExtraTime        int                 `json:"FirstHalfExtraTime"`
	SecondHalfExtraTime       int                 `json:"SecondHalfExtraTime"`
	Winner                    any                 `json:"Winner"`
	Period                    Period              `json:"Period"`
	HomeTeam                  Team                `json:"HomeTeam"`
	AwayTeam                  Team                `json:"AwayTeam"`
	BallPossession            BallPossession      `json:"BallPossession"`
	TerritorialPossesion      any                 `json:"TerritorialPossesion"`
	TerritorialThirdPossesion any                 `json:"TerritorialThirdPossesion"`
	Officials                 []Official          `json:"Officials"`
	MatchStatus               int                 `json:"MatchStatus"` // TODO: Enum
	GroupName                 []LocaleDescription `json:"GroupName"`
	StageName                 []LocaleDescription `json:"StageName"`
	OfficialityStatus         int                 `json:"OfficialityStatus"` // TODO: Enum
	TimeDefined               bool                `json:"TimeDefined"`
	Properties                Properties          `json:"Properties"`
	IsUpdateable              bool                `json:"IsUpdateable"`
}

// Stadium represents the structure of a stadium
type Stadium struct {
	Id                 string              `json:"IdStadium"`
	Name               []LocaleDescription `json:"Name"`
	Capacity           int                 `json:"Capacity"`
	WebAddress         string              `json:"WebAddress"`
	Built              time.Time           `json:"Built"`
	HasRoof            bool                `json:"Roof"`
	Turf               any                 `json:"Turf"`
	CityId             string              `json:"IdCity"`
	CityName           []LocaleDescription `json:"CityName"`
	CountryId          string              `json:"IdCountry"`
	PostalCode         string              `json:"PostalCode"`
	Street             string              `json:"Street"`
	Email              any                 `json:"Email"`
	Fax                any                 `json:"Fax"`
	Phone              any                 `json:"Phone"`
	AffiliationCountry any                 `json:"AffiliationCountry"`
	AffiliationRegion  any                 `json:"AffiliationRegion"`
	Latitude           float64             `json:"Latitude"`
	Longitude          float64             `json:"Longitude"`
	Length             any                 `json:"Length"`
	Width              any                 `json:"Width"`
	Properties         Properties          `json:"Properties"`
	IsUpdateable       bool                `json:"IsUpdateable"`
}

// Weather represents the information about weather
type Weather struct {
	Humidity      string              `json:"Humidity"`
	Temperature   string              `json:"Temperature"`
	WindSpeed     string              `json:"WindSpeed"`
	Type          int                 `json:"Type"`
	TypeLocalized []LocaleDescription `json:"TypeLocalized"`
}

// Team represents the structure of a team object
type Team struct {
	Score         int                 `json:"Score"`
	Side          any                 `json:"Side"`
	Id            string              `json:"IdTeam"`
	PictureUrl    string              `json:"PictureUrl"`
	CountryId     string              `json:"IdCountry"`
	Type          int                 `json:"TeamType"` // TODO: Enum
	AgeType       int                 `json:"AgeType"`  // TODO: Enum
	Tactics       string              `json:"Tactics"`
	Name          []LocaleDescription `json:"TeamName"`
	Abbreviation  string              `json:"Abbreviation"`
	Coaches       []TeamCoach         `json:"Coaches"`
	Players       []TeamPlayer        `json:"Players"`
	Bookings      []Booking           `json:"Bookings"`
	Goals         []Goal              `json:"Goals"`
	Substitutions []TeamSubstitution  `json:"Substitutions"`
	FootballType  int                 `json:"FootballType"` // TODO: Enum
	Gender        Gender              `json:"Gender"`
	AssociationId string              `json:"IdAssociation"`
	ShortClubName string              `json:"ShortClubName"`
}

// TeamCoach represents the structure of a coach object
type TeamCoach struct {
	Id            string              `json:"IdCoach"`
	CountryId     string              `json:"IdCountry"`
	PictureUrl    string              `json:"PictureUrl"`
	Name          []LocaleDescription `json:"Name"`
	Alias         []LocaleDescription `json:"Alias"`
	Role          CoachRole           `json:"Role"`
	SpecialStatus any                 `json:"SpecialStatus"`
}

// TeamPlayer represents the structure of a player from the teams API
type TeamPlayer struct {
	Id            string              `json:"IdPlayer"`
	TeamId        string              `json:"IdTeam"`
	ShirtNumber   int                 `json:"ShirtNumber"`
	Status        int                 `json:"Status"`        // TODO: Enum
	SpecialStatus int                 `json:"SpecialStatus"` // TODO: Enum
	Captain       bool                `json:"Captain"`
	Name          []LocaleDescription `json:"PlayerName"`
	ShortName     []LocaleDescription `json:"ShortName"`
	Position      int                 `json:"Position"` // TODO: Enum
	PlayerPicture TeamPlayerPicture   `json:"PlayerPicture"`
	FieldStatus   int                 `json:"FieldStatus"` // TODO: Enum
	LineupX       int                 `json:"LineupX"`
	LineupY       int                 `json:"LineupY"`
}

// TeamPlayerPicture represents a picture object from the "now" API
type TeamPlayerPicture struct {
	Id         string `json:"Id"`
	PictureUrl string `json:"PictureUrl"`
}

// Booking represents the structure of a booking
type Booking struct {
	Card        int    `json:"Card"`
	Period      Period `json:"Period"`
	EventId     string `json:"IdEvent"`
	EventNumber string `json:"EventNumber"`
	PlayerId    string `json:"IdPlayer"`
	CoachId     string `json:"IdCoach"`
	TeamId      string `json:"IdTeam"`
	Minute      string `json:"Minute"`
	Reason      string `json:"Reason"`
}

// Goal represents the structure of a goal
type Goal struct {
	Id             string `json:"IdGoal"`
	TeamId         string `json:"IdTeam"`
	Type           int    `json:"Type"` // TODO: Enum
	PlayerId       string `json:"IdPlayer"`
	Minute         string `json:"Minute"`
	AssistPlayerId string `json:"IdAssistPlayer"`
	Period         Period `json:"Period"`
}

// TeamSubstitution represents the information about a substitution
type TeamSubstitution struct {
	EventId       string              `json:"IdEvent"`
	Period        Period              `json:"Period"`
	Reason        int                 `json:"Reason"`             // TODO: Enum
	Position      int                 `json:"SubstitutePosition"` // TODO: Enum
	PlayerOffId   string              `json:"IdPlayerOff"`
	PlayerOnId    string              `json:"IdPlayerOn"`
	PlayerOffName []LocaleDescription `json:"PlayerOffName"`
	PlayerOnName  []LocaleDescription `json:"PlayerOnName"`
	Minute        string              `json:"Minute"`
	TeamId        string              `json:"IdTeam"`
}

// BallPossession represents the information about ball possession
type BallPossession struct {
	Intervals   any     `json:"Intervals"`
	LastX       any     `json:"LastX"`
	OverallHome float32 `json:"OverallHome"`
	OverallAway float32 `json:"OverallAway"`
}

// Official represents the information about an offical
type Official struct {
	Id            string              `json:"OfficialId"`
	CountryId     string              `json:"IdCountry"`
	Name          []LocaleDescription `json:"Name"`
	ShortName     []LocaleDescription `json:"NameShort"`
	Type          OfficialType        `json:"OfficialType"`
	TypeLocalized []LocaleDescription `json:"TypeLocalized"`
}

// MatchesResponse represents the information returned from the calendar API
type MatchesResponse struct {
	PaginationResponse
	Results []MatchData `json:"Results"`
}

// MatchData represents the structure of a match object
type MatchData struct {
	CompetitionId          string              `json:"IdCompetition"`
	SeasonId               string              `json:"IdSeason"`
	StageId                string              `json:"IdStage"`
	GroupId                string              `json:"IdGroup"`
	Weather                Weather             `json:"Weather"`
	Attendance             string              `json:"Attendance"`
	Id                     string              `json:"IdMatch"`
	MatchDay               string              `json:"MatchDay"`
	StageName              []LocaleDescription `json:"StageName"`
	GroupName              []LocaleDescription `json:"GroupName"`
	CompetitionName        []LocaleDescription `json:"CompetitionName"`
	SeasonName             []LocaleDescription `json:"SeasonName"`
	SeasonShortName        any                 `json:"SeasonShortName"`
	Date                   time.Time           `json:"Date"`
	LocalDate              time.Time           `json:"LocalDate"`
	Home                   Team                `json:"Home"`
	Away                   Team                `json:"Away"`
	HomeTeamScore          int                 `json:"HomeTeamScore"`
	AwayTeamScore          int                 `json:"AwayTeamScore"`
	AggregateHomeTeamScore any                 `json:"AggregateHomeTeamScore"`
	AggregateAwayTeamScore any                 `json:"AggregateAwayTeamScore"`
	HomeTeamPenaltyScore   int                 `json:"HomeTeamPenaltyScore"`
	AwayTeamPenaltyScore   int                 `json:"AwayTeamPenaltyScore"`
	LastPeriodicUpdate     any                 `json:"LastPeriodicUpdate"`
	Leg                    any                 `json:"Leg"`
	IsHomeMatch            bool                `json:"IsHomeMatch"`
	Stadium                Stadium             `json:"Stadium"`
	IsTicketSalesAllowed   bool                `json:"IsTicketSalesAllowed"`
	MatchTime              string              `json:"MatchTime"`
	SecondHalfTime         any                 `json:"SecondHalfTime"`
	FirstHalfTime          any                 `json:"FirstHalfTime"`
	FirstHalfExtraTime     int                 `json:"FirstHalfExtraTime"`
	SecondHalfExtraTime    int                 `json:"SecondHalfExtraTime"`
	Winner                 string              `json:"Winner"`
	MatchReportUrl         any                 `json:"MatchReportUrl"`
	PlaceholderA           string              `json:"PlaceHolderA"`
	PlaceholderB           string              `json:"PlaceHolderB"`
	BallPossession         BallPossession      `json:"BallPossession"`
	Officials              []Official          `json:"Officials"`
	MatchStatus            int                 `json:"MatchStatus"` // TODO: Enum
	ResultType             int                 `json:"ResultType"`  // TODO: Enum
	MatchNumber            int                 `json:"MatchNumber"`
	TimeDefined            bool                `json:"TimeDefined"`
	OfficialityStatus      int                 `json:"OfficialityStatus"` // TODO: Enum
	MatchLegInfo           any                 `json:"MatchLegInfo"`
	Properties             Properties          `json:"Properties"`
	IsUpdateable           bool                `json:"IsUpdateable"`
}

// TimelineResponse represents the response from the timeline API
type TimelineResponse struct {
	StageId       string          `json:"IdStage"`
	MatchId       string          `json:"IdMatch"`
	CompetitionId string          `json:"IdCompetition"`
	SeasonId      string          `json:"IdSeason"`
	GroupId       string          `json:"IdGroup"`
	Events        []TimelineEvent `json:"Event"`
}

// TimelineEvent represent the structure of a timeline event
type TimelineEvent struct {
	Id                  string              `json:"EventId"`
	TeamId              string              `json:"IdTeam"`
	Timestamp           time.Time           `json:"Timestamp"`
	MatchMinute         string              `json:"MatchMinute"`
	Period              Period              `json:"Period"`
	HomeGoals           int                 `json:"HomeGoals"`
	AwayGoals           int                 `json:"AwayGoals"`
	Type                MatchEvent          `json:"Type"`
	Qualifiers          any                 `json:"Qualifiers"`
	TypeLocalized       []LocaleDescription `json:"TypeLocalized"`
	HomePenaltyGoals    int                 `json:"HomePenaltyGoals"`
	AwayPenaltyGoals    int                 `json:"AwayPenaltyGoals"`
	Description         []LocaleDescription `json:"EventDescription"`
	VarNotificationData VARNotification     `json:"VarNotificationData"`
}

// VARNotification struct represents information about VAR
type VARNotification struct {
	Incident int `json:"Incident"`
	Reason   int `json:"Reason"`
	Status   int `json:"Status"`
	Result   int `json:"Result"`
}

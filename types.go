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
	OFFSIDE            MatchEvent = 15
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
	ID                        string                       `json:"IdMatch"`
	StageID                   string                       `json:"IdStage"`
	GroupID                   string                       `json:"IdGroup"`
	SeasonID                  string                       `json:"IdSeason"`
	CompetitionID             string                       `json:"IdCompetition"`
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
	WinnerID                  string                       `json:"Winner"`
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
	ID                 string                       `json:"IdStadium"`
	Name               []DefaultDescriptionResponse `json:"Name"`
	Capacity           string                       `json:"Capacity"`
	WebAddress         string                       `json:"WebAddress"`
	Built              string                       `json:"Built"`
	HasRoof            bool                         `json:"Roof"`
	Turf               string                       `json:"Turf"`
	CityID             string                       `json:"IdCity"`
	City               []DefaultDescriptionResponse `json:"CityName"`
	CountryID          string                       `json:"IdCountry"`
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
	Width              string                       `json:"Width"`
	Properties         interface{}                  `json:"Properties"`
	IsUpdateable       bool                         `json:"IsUpdateable"`
}

type TeamResponse struct {
	Score         int                          `json:"Score"`
	Side          string                       `json:"Side"`
	ID            string                       `json:"TeamID"`
	PictureURL    string                       `json:"PictureURL"`
	CountryID     string                       `json:"IdCountry"`
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
	AssociationID string                       `json:"IdAssociation"`
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
	ID            string                       `json:"IdPlayer"`
	TeamID        string                       `json:"IdTeam"`
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
	ID            string                       `json:"IdCoach"`
	CountryID     string                       `json:"IdCountry"`
	Name          []DefaultDescriptionResponse `json:"Name"`
	Alias         []DefaultDescriptionResponse `json:"Alias"`
	Role          int                          `json:"Role"` // TODO: Enum
	SpecialStatus string                       `json:"SpecialStatus"`
}

type BookingResponse struct {
	Card        int        `json:"Card"`
	Period      PeriodEnum `json:"Period"`
	EventID     string     `json:"IdEvent"`
	EventNumber string     `json:"EventNumber"`
	PlayerID    string     `json:"IdPlayer"`
	CoachID     string     `json:"IdCoach"`
	TeamID      string     `json:"IdTeam"`
	Minute      string     `json:"Minute"`
	Reason      string     `json:"Reason"`
}

type SubstitutionResponse struct {
	EventID       string                       `json:"IdEvent"`
	Period        PeriodEnum                   `json:"Period"`
	Reason        int                          `json:"Reason"`            // TODO: Enum
	Position      int                          `json:"SubstituePosition"` // TODO: Enum
	PlayerOffID   string                       `json:"IdPlayerOff"`
	PlayerOnID    string                       `json:"IdPlayerOn"`
	PlayerOffName []DefaultDescriptionResponse `json:"PlayerOffName"`
	PlayerOnName  []DefaultDescriptionResponse `json:"PlayerOnName"`
	Minute        string                       `json:"Minute"`
	TeamID        string                       `json:"TeamID"`
}

type OfficialResponse struct {
	ID            string                       `json:"OfficialsId"`
	CountryID     string                       `json:"IdCountry"`
	Name          []DefaultDescriptionResponse `json:"Name"`
	ShortName     []DefaultDescriptionResponse `json:"ShortName"`
	Type          int                          `json:"OfficialType"` // TODO: Enum
	TypeLocalized []DefaultDescriptionResponse `json:"TypeLocalized"`
}

type GoalResponse struct {
	ID             string     `json:"IdGoal"`
	TeamID         string     `json:"IdTeam"`
	Type           int        `json:"Type"` // TODO: Enum
	PlayerID       string     `json:"IdPlayer"`
	Minute         string     `json:"Minute"`
	AssistPlayerID string     `json:"IdAssistPlayer"`
	Period         PeriodEnum `json:"Period"`
}

type CompetitionResponse struct {
	CompetitionID       string                       `json:"IdCompetition"`
	Name                []DefaultDescriptionResponse `json:"Name"`
	ConfederationID     []string                     `json:"IdConfederation"`
	MemberAssociationID []string                     `json:"IdMemberAssociation"`
	OwnerID             string                       `json:"IdOwner"`
	Gender              Gender                       `json:"Gender"`
	FootballType        int                          `json:"FootballType"`    // TODO: Enum
	TeamType            int                          `json:"TeamType"`        // TODO: Enum
	Type                int                          `json:"CompetitionType"` // TODO: Enum
	Properties          interface{}                  `json:"Properties"`
	IsUpdateable        bool                         `json:"IsUpdateable"`
}

type PlayerResponse struct {
	ID                       string                       `json:"IdPlayer"`
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

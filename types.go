package goball

import "time"

//Status describes the state of a match, e.g. scheduled, or finished
type Status string

//Venue is the name of the location of the match
type Venue string

//LeagureCode is the constant code ofa a league
//type LeagueCode string

type PastNext string

const (
	Past PastNext = "p"
	Next PastNext = "n"
)

const (
	Scheduled Status = "SCHEDULED"
	Timed     Status = "TIMED"
	Postponed Status = "POSTPONED"
	InPlay    Status = "IN_PLAY"
	Finished  Status = "FINISHED"
)

const (
	VenueHome Venue = "home"
	VenueAway Venue = "away"
)

const (
	Germany1Bundesliga    string = "BL1"
	Germany2Bundesliga    string = "BL2"
	Germany3Bundesliga    string = "BL2"
	GermanyDfbPokal       string = "DFB"
	EnglandPremierLeague  string = "PL"
	EnglandLeagueOne      string = "EL1"
	EnglandChampionship   string = "ELC"
	EnglandFACup          string = "FAC"
	ItalySerieA           string = "SA"
	ItalySerieB           string = "SB"
	SpainPrimeraDivision  string = "PD"
	SpainSegundaDivision  string = "SD"
	SpainCopaDelRey       string = "CDR"
	FrancueLigue1         string = "FL!"
	FranceLigue2          string = "FL2"
	NetherlandsEredivise  string = "DED"
	PortugalPrimeiraLiga  string = "PPL"
	GreeceSuperLeague     string = "GSL"
	EuropeChampionsLeague string = "CL"
	EuropeUefaLeague      string = "EL"
	EuropeCupOfNations    string = "EC"
	WorldCup              string = "WC"
)

type CompetitionResult struct {
	Count        int            `json:"count"`
	Competitions []*Competition `json:"competitions"`
}

//Competition desribes a competition, like a tournament or a league
type Competition struct {
	ID          uint32    `json:"id"`
	Name        string    `json:"name"`
	Type        string    `json:"type"`
	LastUpdated time.Time `json:"lastUpdated"`
}

//Score describes the current number of goals for each team
type Score struct {
	Home uint8
	Away uint8
}

//Result describes several Scores of a fixture splitup into actual score, halftime, extratime and penalty
type Result struct {
	Winner   string
	Duration string
	//HalfTime        Score
	//ExtraTime       Score
	FullTime        Score
	PenaltyShootout Score
}

//Head2Head describes statistics for this match, e.g. how often the two teams have played against each other
type Head2Head struct {
	Count               uint16
	TimeFrameStart      time.Time
	TimeFrameEnd        time.Time
	HomeTeamWins        uint16
	AwayTeamWins        uint16
	Draws               uint16
	LastHomeWinHomeTeam *Fixture
	LastWinHomeTeam     *Fixture
	LastAwayWinAwayTeam *Fixture
	LastWinAwayTeam     *Fixture
	Fixtures            []Fixture
}

//Fixture describes a match including date, status...
type Fixture struct {
	ID            uint32    `json:"id"`
	CompetitionID uint16    `json:"competitionId"`
	Date          time.Time `json:"utcDate"`
	Status        Status
	Matchday      uint16
	HomeTeam      *Team
	AwayTeam      *Team
	Result        Result `json:"score"`
}

//FixtureListResult is the result of a requst of a fixtureList for a competition
type FixtureListResult struct {
	Competition *Competition `json:"competition"`
	Matches     []Fixture    `json:"matches"`
}

//TimeFramedFixtureListResult is the result of a requst of a fixtureListRequest
type TimeFramedFixtureListResult struct {
	TimeFrameStart time.Time
	TimeFrameEnd   time.Time
	FixtureListResult
}

//CompetitionTeamsResult describes the result of the request of teams for a competition
type CompetitionTeamsResult struct {
	Count uint8
	Teams []Team
}

//Team describes a team
type Team struct {
	Name      string
	Shortname string
	CrestURL  string
}

//Player describes a football player, e.g. position, name...
type Player struct {
	Name         string
	Position     string
	JerseyNumber uint8
	DateOfBirth  time.Time
	Nationality  string
	ContractUnti time.Time
	MarketValue  string
}

//PlayerList result of a list of players
type PlayerList struct {
	Count   uint8
	Players []Player
}

//TeamLeagueStanding describes the current standing of a team in a league
type LeagueTeamStanding struct {
	Rank            uint8
	Team            string
	TeamID          uint16
	PlayedGames     uint8
	CrestURI        string
	Points          uint8
	Goals           uint8
	GoalsAgainst    uint8
	GoalsDifference uint8
}

type GroupTeamStanding struct {
	Group string
	LeagueTeamStanding
}

//LeagueTable describes the current table and the team standings
type GroupCompetitionTable struct {
	LeagueCaption string
	Matchday      uint8
	Standings     GroupStanding
}

//GroupStanding reflects the stnading within a group of a competition, e.g. for chamions league or worldcup
type GroupStanding map[string][]GroupTeamStanding

//LeagueTable reflects the current table in a league, e.g. Bundesliga without groups
type LeagueTable struct {
	LeagueCaption string
	Matchday      uint8
	Standing      []LeagueTeamStanding
}

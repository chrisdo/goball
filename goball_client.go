package goball

import "net/http"
import "time"
import "fmt"

//Client is a client for the football-data.org
type Client struct {
	httpClient *http.Client
	authToken  string
}

func NewGoballClient(token string) Client {
	httpClient := new(http.Client)
	httpClient.Timeout = time.Duration(10 * time.Second)
	return Client{httpClient: httpClient, authToken: token}
}

func (c Client) NewCompetitionRequest() CompetitionRequest {
	req := CompetitionRequest{request{"/competitions", nil}}
	return req
}

func (c Client) NewCompetitionTeamsRequest(competitionId uint32) CompetitionTeamsRequest {
	req := CompetitionTeamsRequest{request{fmt.Sprintf("/competitions/%d/teams", competitionId), nil}}
	return req
}

func (c Client) NewCompetitionLeagueTableRequest(competitionId uint32) CompetitionLeagueTableRequest {
	req := CompetitionLeagueTableRequest{request{fmt.Sprintf("/competitions/%d/leagueTable", competitionId), nil}}
	return req
}

func (c Client) NewCompetitionFixturesRequest(competitionId uint32) CompetitionFixturesRequest {
	req := CompetitionFixturesRequest{request{fmt.Sprintf("/competitions/%d/fixtures", competitionId), nil}}
	return req
}

func (c Client) NewFixtureListRequest() FixtureListRequest {
	req := FixtureListRequest{request{"/fixtures", nil}}
	return req
}

func (c Client) NewFixtureRequest(fixtureId uint16) FixtureRequest {
	req := FixtureRequest{request{fmt.Sprintf("/fixtures/%d", fixtureId), nil}}
	return req
}

func (c Client) NewTeamFixturesRequest(teamId uint16) TeamFixturesRequest {
	req := TeamFixturesRequest{request{fmt.Sprintf("teams/%d/fixtures", teamId), nil}}
	return req
}

func (c Client) NewTeamRequest(teamId uint16) TeamRequest {
	req := TeamRequest{request{fmt.Sprintf("teams/%d", teamId), nil}}
	return req
}

func (c Client) NewTeamPlayersRequest(teamId uint16) TeamPlayersRequest {
	req := TeamPlayersRequest{request{fmt.Sprintf("teams/%d/players", teamId), nil}}
	return req
}

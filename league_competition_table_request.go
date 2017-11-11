package goball

import (
	"fmt"
)

//CompetitionLeagueTableRequest can be used to request a list of all competitions.
//The request can be modified with a filter to set a specific season
type CompetitionLeagueTableRequest struct {
	request
}

//GetLeagueTable is to be used if the competition is just a league based competition, e.g. Premier League
func (req *CompetitionLeagueTableRequest) GetLeagueTable(c *Client) (leagueResult *LeagueTable) {

	j, err := req.getData(c)
	if err != nil {
		fmt.Printf("something wnet wrong: %s\n", err)
	}

	j.Decode(&leagueResult)
	return
}

//GetGroupTable is to be used if the competitions is organised in groups, e.g. CL or World Cup
func (req *CompetitionLeagueTableRequest) GetGroupTable(c *Client) (groupResult *GroupCompetitionTable) {

	j, err := req.getData(c)
	if err != nil {
		fmt.Printf("something wnet wrong: %s\n", err)
	}

	j.Decode(&groupResult)
	return
}

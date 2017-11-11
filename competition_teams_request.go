package goball

import (
	"fmt"
)

//CompetitionRequest can be used to request a list of all competitions.
//The request can be modified with a filter to set a specific season
type CompetitionTeamsRequest struct {
	request
}

func (req *CompetitionTeamsRequest) Send(c *Client) (teamsResult *CompetitionTeamsResult) {

	j, err := req.getData(c)
	if err != nil {
		fmt.Printf("something wnet wrong: %s\n", err)
	}

	j.Decode(&teamsResult)
	return
}

package goball

import (
	"fmt"
)

//TeamPlayersRequest can be used to request a list of all competitions.
type TeamPlayersRequest struct {
	request
}

//Send sends the request
func (req *TeamPlayersRequest) Send(c *Client) (teamPlayersResult *PlayerList) {

	j, err := req.getData(c)
	if err != nil {
		fmt.Printf("something wnet wrong: %s\n", err)
	}

	j.Decode(&teamPlayersResult)
	return
}

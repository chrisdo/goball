package goball

import (
	"fmt"
)

//CompetitionRequest can be used to request a list of all competitions.
//The request can be modified with a filter to set a specific season
type TeamRequest struct {
	request
}

//Send sends the request
func (req *TeamRequest) Send(c *Client) (teamResult *Team) {

	j, err := req.getData(c)
	if err != nil {
		fmt.Printf("something wnet wrong: %s\n", err)
	}

	j.Decode(&teamResult)
	return
}

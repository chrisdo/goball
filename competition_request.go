package goball

import (
	"fmt"
	"net/url"
)

//CompetitionRequest can be used to request a list of all competitions.
//The request can be modified with a filter to set a specific season
type CompetitionRequest struct {
	request
}

//Season modifies the request filter and asets the specified season
func (c *CompetitionRequest) Season(year uint16) {
	if c.parameter == nil {
		c.parameter = url.Values{}
	}
	c.parameter.Set(ParamSeason, fmt.Sprintf("%d", year))
}

//Send sends the request
func (req *CompetitionRequest) Send(c *Client) (competitions []*Competition) {

	j, err := req.getData(c)
	if err != nil {
		fmt.Printf("something wnet wrong: %s\n", err)
	}

	j.Decode(&competitions)
	return
}

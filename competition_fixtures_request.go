package goball

import (
	"fmt"
	"net/url"
)

//CompetitionFixturesRequest can be used to request a list of all competitions.
//The request can be modified with a filter to set a specific season
type CompetitionFixturesRequest struct {
	request
}

func (c *CompetitionRequest) Matchday(matchday uint16) {
	if c.parameter == nil {
		c.parameter = url.Values{}
	}
	c.parameter.Set(ParamMatchday, fmt.Sprintf("%d", matchday))
}

func (req *CompetitionFixturesRequest) Send(c *Client) (fixtureListResult *FixtureListResult) {

	j, err := req.getData(c)
	if err != nil {
		fmt.Printf("something wnet wrong: %s\n", err)
	}

	j.Decode(&fixtureListResult)
	return
}

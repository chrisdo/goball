package goball

import (
	"fmt"
	"net/url"
	"strings"
)

//FixturesRequest can be used to request a list of all competitions.
//The request can be modified with a filter to set a specific season
type FixtureListRequest struct {
	request
}

func (c *FixtureListRequest) League(leagues ...string) {
	if c.parameter == nil {
		c.parameter = url.Values{}
	}
	c.parameter.Set(ParamLeague, fmt.Sprintf("%s", strings.Join(leagues, ",")))
}

func (c *FixtureListRequest) Timeframe(pastNext PastNext, days uint8) {
	if c.parameter == nil {
		c.parameter = url.Values{}
	}
	c.parameter.Set(ParamTimeframe, fmt.Sprintf("%s%d", pastNext, days))
}

func (req *FixtureListRequest) Send(c *Client) (fixtureListResult *TimeFramedFixtureListResult) {

	j, err := req.getData(c)
	if err != nil {
		fmt.Printf("something wnet wrong: %s\n", err)
	}

	j.Decode(&fixtureListResult)
	return
}

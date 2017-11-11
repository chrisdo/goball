package goball

import (
	"fmt"
	"net/url"
)

//FixturesRequest can be used to request a list of all competitions.
//The request can be modified with a filter to set a specific season
type TeamFixturesRequest struct {
	request
}

func (c *TeamFixturesRequest) Season(season uint16) {
	if c.parameter == nil {
		c.parameter = url.Values{}
	}
	c.parameter.Set(ParamSeason, fmt.Sprintf("%d", season))
}

func (c *TeamFixturesRequest) Timeframe(pastNext PastNext, days uint8) {
	if c.parameter == nil {
		c.parameter = url.Values{}
	}
	c.parameter.Set(ParamTimeframe, fmt.Sprintf("%s%d", pastNext, days))
}

func (c *TeamFixturesRequest) Venue(venue Venue) {
	if c.parameter == nil {
		c.parameter = url.Values{}
	}
	c.parameter.Set(ParamVenue, fmt.Sprintf("%s", venue))
}

func (req *TeamFixturesRequest) Send(c *Client) (fixtureListResult *TimeFramedFixtureListResult) {

	j, err := req.getData(c)
	if err != nil {
		fmt.Printf("something wnet wrong: %s\n", err)
	}

	j.Decode(&fixtureListResult)
	return
}

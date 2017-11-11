package goball

import (
	"fmt"
	"net/url"
)

//FixtureRequest can be used to request a list of all competitions.
//The request can be modified with a filter to set a specific season
type FixtureRequest struct {
	request
}

func (c *FixtureRequest) Head2Head(count uint8) {
	if c.parameter == nil {
		c.parameter = url.Values{}
	}
	c.parameter.Set(ParamHead2Head, fmt.Sprintf("%d", count))
}

func (req *FixtureRequest) Send(c *Client) (fixtureResult *Fixture) {

	j, err := req.getData(c)
	if err != nil {
		fmt.Printf("something wnet wrong: %s\n", err)
	}

	j.Decode(&fixtureResult)
	return
}

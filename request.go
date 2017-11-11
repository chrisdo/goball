package goball

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	ADDRESS string = "http://football-data.org/"
	VERSION string = "v1"

	ParamLeague    string = "league"
	ParamSeason    string = "season"
	ParamTimeframe string = "timeFrame"
	ParamVenue     string = "venue"

	ParamHead2Head string = "head2head"
	ParamMatchday  string = "matchday"
)

//Request describes the URI endpoint and contains possible parameters, TODO: add a possible client for football-data.org
type request struct {
	uri       string //the uri of the resource to be requested, e.g. /v1/competitions
	parameter url.Values
}

//run the execute to the request
func (r request) getData(c *Client) (j *json.Decoder, err error) {
	//1. generate the url for this request based on the url to football-data.org, the uri and the parameters
	//parse the adress
	url, err := url.Parse(ADDRESS)
	if err != nil {
		return nil, err
	}

	url.Path = VERSION + r.uri
	url.RawQuery = r.parameter.Encode() // the raw execute parameters without ?

	//set header if auth thoken is set
	req, err := http.NewRequest("GET", url.String(), nil)
	req.Header = http.Header{}
	if c.authToken != "" {
		//set the stuff
		req.Header.Set("X-Auth-Token", c.authToken)
	}
	req.Header.Set("X-Response-Control", "minified")
	fmt.Printf("requesting %s\n", url.String())

	fmt.Printf("Header %s\n", req.Header)
	response, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	fmt.Printf("ResponseCode: %s\n", response.Status)
	body, err := ioutil.ReadAll(response.Body)
	j = json.NewDecoder(bytes.NewReader(body))
	return j, nil
}

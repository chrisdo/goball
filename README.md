# goball
This library is intended to provide access to the football-data.org API in Go.

All features from https://api.football-data.org/documentation are supported.

## Attention
This library is still in development state and not adivsed to be used in production right now. I just started to convert from v1 to v4, still need to update a few requests and response types.
Also ,the naming is not correct any more, e.g. in v1 a Match was a Fixture. So partly in code it is still a Fixture.

### Planned Actions
* cleanup code
* cleanup and improve documentation
* add tests for all datasets

## How to use 

``` Go
package main

import (
	"fmt"
	"github.com/chrisdo/goball"
)

func main() {

    goballClient := goball.NewGoballClient("<your api key>")

    competitionRequest := goballClient.NewCompetitionRequest()
    comps := competitionRequest.Send(&goballClient)
    for _, c := range comps.Competitions {
	fmt.Printf("%s: %d\n", c.Caption, c.ID)
    }


    //ATTENTION: THE FOLLOWING CODE SAMPLE IS NOT UPDATED YET TO v4

    //the resulting standings of a league competitions, e.g. premier league
    leagueTableRequest := goballClient.NewCompetitionLeagueTableRequest(id)
    leagueTable := leagueTableRequest.GetLeagueTable(&goballClient)
    
    for _, s := range leagueTable.Standing {
        fmt.Printf("\t%d %s %d\n", s.Rank, s.Team, s.Points)	
    }
    
    //the resulting standings of a competition having groups, e.g. Champions league
    groupTableRequest := goballClient.NewCompetitionLeagueTableRequest(id)
    groupTable := groupTableRequest.GetGroupTable(&goballClient)
    for k, s := range groupTable.Standings {
	    fmt.Printf("\nGroup %s", k)
        for _, g := range s {
            fmt.Printf("\t%d %s %d\n", g.Rank, g.Team, g.Points)
        }
    }
}
```



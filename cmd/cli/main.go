package main

import (
	"fmt"
	"os"

	"github.com/chrisdo/goball"
)

func main() {

	if len(os.Args) != 2 {
		panic("Must provide api key as parameter")
	}
	footballDataApiKey := os.Args[1]

	goballClient := goball.NewGoballClient(footballDataApiKey)
	competitionsReqyest := goballClient.NewCompetitionRequest()
	comps := competitionsReqyest.Send(&goballClient)

	fmt.Printf("Comps: %d\n", comps.Count)
	for _, f := range comps.Competitions {

		fmt.Printf("%v\n", f)

	}

	wmFixturesRequest := goballClient.NewCompetitionFixturesRequest(2000)
	//TODO: API needs to be updated to v4 ... :D
	fixtures := wmFixturesRequest.Send(&goballClient)

	for _, f := range fixtures.Matches {

		fmt.Printf("%v\n", f)

	}
}

package home

import (
	"fmt"
	"log"
	"testing"
)

var ProdcastsTest = []Prodcasts{
	{ProdcastAuthor: "AuthorOne", Comments: "comment one, comment two", StreamedAt: 1221, Liked: 21, Plays: 20},
}

func prodcasttests(t *testing.T, err error) {

	fmt.Println(ProdcastsTest)

	if err != nil {
		log.Fatal(err)
	}

}

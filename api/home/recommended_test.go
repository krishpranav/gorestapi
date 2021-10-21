package home

import (
	"fmt"
	"log"
	"testing"
)

var RecommendTracks = []Recommended {
	{TopTracks: "TrackOne", ArtistComposed: "Artist One", Plays: 1000, DateReleased: 2020},
}

func recommendedtest(t *testing.T, err error) {

	fmt.Println(RecommendTracks)

	if err != nil {
		log.Fatal(err)
	}

}
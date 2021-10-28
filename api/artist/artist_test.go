package artist

import (
	"fmt"
	"testing"
)

var ArtistVar = []Artist{
	{Name: "Some", Location: "Some where in the world", JoinedAt: 100},
}

func ArtistTest(t *testing.T, err error) {
	fmt.Println(ArtistVar)

	if err != nil {
		panic(err)
	}
}

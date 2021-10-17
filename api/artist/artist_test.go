package artist

import (
	"fmt"
	"testing"
)

var ArtistTest = []Artist {
	{Name: "Some", Location: "Some where in the world", JoinedAt: 100},
}

func artisttest(t *testing.T) {
	fmt.Println(ArtistTest)
}

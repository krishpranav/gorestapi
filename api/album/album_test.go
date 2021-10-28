package album

import (
	"fmt"
	"testing"
)

var AlbumVar = []Album{
	{ID: "1", Title: "Hello"},
}

func AlbumTest(t *testing.T, err error) {
	fmt.Println(AlbumVar)

	if err != nil {
		panic(err)
	}
}

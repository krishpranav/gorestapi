package album

import (
	"fmt"
	"testing"
)

var AlbumVar = []Album {
	{ID: "1", Title: "Hello"},
}

func AlbumTest(t *testing.T) {
	fmt.Println(AlbumVar)
}

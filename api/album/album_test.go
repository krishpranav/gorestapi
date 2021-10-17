package album

import (
	"fmt"
	"testing"
)

var album = []Album {
	{ID: "1", Title: "Hello"},
}

func albumtest(t *testing.T) {
	fmt.Println(album)
}

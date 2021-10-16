package about

import (
	"fmt"
	"testing"
)

var aboutapi = []Aboutapi {
	{App: "Test"},
}

func about(t *testing.T) {
	fmt.Println(aboutapi)
}
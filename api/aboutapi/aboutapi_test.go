package about

import (
	"fmt"
	"testing"
)

var AboutVar = []Aboutapi {
	{App: "Test"},
}

func AboutTest(t *testing.T) {
	fmt.Println(AboutVar)
}
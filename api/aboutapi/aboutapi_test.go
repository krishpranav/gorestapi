package about

import (
	"fmt"
	"testing"
)

var AboutVar = []Aboutapi {
	{App: "Test"},
}

func abouttest(t *testing.T) {
	fmt.Println(AboutVar)
}
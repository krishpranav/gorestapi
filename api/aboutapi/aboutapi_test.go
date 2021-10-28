package about

import (
	"fmt"
	"testing"
)

var AboutVar = []Aboutapi{
	{App: "Test", Github: "https://github.com/krishpranav/gorestapi", Author: "Krishpranav", Version: 4},
}

func AboutTest(t *testing.T, err error) {
	fmt.Println(AboutVar)

	if err != nil {
		panic(err)
	}
}

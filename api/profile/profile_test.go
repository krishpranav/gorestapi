package aboutme

import (
	"fmt"
	"testing"
)

var Profile = []Aboutme {
	{User: "Some", JoinedAt: 100, Toptracks: "Trackone", Following: "SomeOne"},
}

func profiletest(t *testing.T, err error) {
	fmt.Println(Profile)

	if err != nil {
		panic(err)
	}

}
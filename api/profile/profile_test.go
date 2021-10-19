package profile

import (
	"fmt"
	"testing"
)

var ProfileTest = []Profile {
	{User: "Some", JoinedAt: 100, Toptracks: "Trackone", Following: "SomeOne"},
}

func profiletest(t *testing.T, err error) {
	fmt.Println(ProfileTest)

	if err != nil {
		panic(err)
	}

}
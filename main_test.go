package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

/* perform post request */
func postRequest(t *testing.T, err error) {
	r := gin.Default()

	posttest := r.Group("/posttest")
	{
		posttest.POST("/post", func(c *gin.Context) {

		})
	}

	if err != nil {
		panic(err)
	}

	r.Run()
}

func sqlconnection(err error) {
	file, err := os.Open("db.sql")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(file.Name() + "db sql")
	}

}

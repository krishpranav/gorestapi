package main

import (
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

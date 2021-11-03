package main

import (
	"testing"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

/* frontend test class */
func frontendtest(t *testing.T, err error) {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./views/test", true)))

	testapi := r.Group("/test")
	{
		testapi.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "test",
			})
		})
	}

	if err != nil {
		panic(err)
	}

	r.Run()
}

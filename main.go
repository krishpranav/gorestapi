package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

type aboutapi struct {
    App string `json:"app"`
    Version float64 `json:"version"`
}

var aboutapp = []aboutapi {
    {App: "SuperRestApi", Version: 1},
}

var albums = []album{
    {ID: "1", Title: "ExampleOne", Artist: "User1", Price: 39.99},
    {ID: "2", Title: "ExampleTwo", Artist: "User2", Price: 39.99},
    {ID: "3", Title: "ExampleThree", Artist: "User3", Price: 39.99},
}

var mainpage = "Hello World"

func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.GET("/albums/:id", getAlbumByID)
    router.POST("/albums", postAlbums)
    router.GET("/about", getAbout)
    router.GET("/", getMain)

    router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

func getAbout(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, aboutapp)
}

func postAlbums(c *gin.Context) {
    var newAlbum album

    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func getMain(c *gin.Context) {
    c.IndentedJSON(http.StatusCreated, mainpage)
}


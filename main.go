package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	about "github.com/krishpranav/gorestapi/api/aboutapi"
	"github.com/krishpranav/gorestapi/api/album"
	"github.com/krishpranav/gorestapi/api/artist"
	"github.com/krishpranav/gorestapi/api/home"
	"github.com/krishpranav/gorestapi/api/profile"
	_ "github.com/mattn/go-sqlite3"
)

var aboutapp = []about.Aboutapi{
	{App: "gorestapi", Author: "krishpranav", Github: "https://github.com/krishpranav/gorestapi", Version: 4},
}

var albums = []album.Album{
	{ID: "1", Title: "AlbumOne", Artist: "ArtistOne", Price: 39.99},
}

var artists = []artist.Artist{
	{Name: "ArtistOne", Location: "IN", JoinedAt: 2017},
	{Name: "ArtistTwo", Location: "USA", JoinedAt: 2018},
	{Name: "ArtistThree", Location: "UK", JoinedAt: 2019},
	{Name: "ArtistFour", Location: "IN", JoinedAt: 2020},
	{Name: "ArtistFive", Location: "USA", JoinedAt: 2021},
}

var profiles = []profile.Profile{
	{User: "UserOne", Password: "", JoinedAt: 2020, Toptracks: "TrackOne, TrackTwo, TrackThree", Following: "ArtistOne", Followers: 21},
	{User: "UserTwo", Password: "", JoinedAt: 2021, Toptracks: "TrackFour, TrackFive, TrackSix", Following: "ArtistTwo, ArtistThree", Followers: 0},
}

var Top = []home.Recommended{
	{TopTracks: "TrackOne", ArtistComposed: "Artist One", Plays: 1000, DateReleased: 2020},
	{TopTracks: "TrackTwo", ArtistComposed: "Artist Two", Plays: 10000, DateReleased: 2021},
	{TopTracks: "TrackThree", ArtistComposed: "Artist Three", Plays: 10000, DateReleased: 2019},
}

var Prodcast = []home.Prodcasts{
	{ProdcastAuthor: "AuthorOne", Comments: "comment one, comment two", StreamedAt: 1.22, Liked: 100, Plays: 200},
	{ProdcastAuthor: "AuthorTwo", Comments: "comment three, comment four", StreamedAt: 4.11, Liked: 200, Plays: 300},
	{ProdcastAuthor: "AuthorThree", Comments: "comment five, comment six", StreamedAt: 2.22, Liked: 300, Plays: 400},
	{ProdcastAuthor: "AuthorFour", Comments: "this is good", StreamedAt: 1.11, Liked: 400, Plays: 500},
	{ProdcastAuthor: "AuthorFive", Comments: "awsome", StreamedAt: 12.12, Liked: 500, Plays: 600},
}

var mainpage = "Hello World"

func main() {

	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./views/dist/", true)))

	api := router.Group("/api")
	{
		/* albums api */
		router.GET("/albums", getAlbums)
		router.GET("/albums/:id", getAlbumByID)
		api.POST("/albums", postAlbums)

		/* about api */
		router.GET("/about", getAbout)

		/* artist api */
		router.GET("/artist", getArtistData)
		api.POST("/artist", postArtist)

		/* main index api */
		router.GET("/", getMain)

		/* profile api */
		router.GET("/profile", getProfile)
		api.POST("/profile/:id", postProfile)

		/* top tracks api */
		router.GET("/toptracks", getRecommendedTracks)
		api.POST("/toptracks", postNewTopTracks)

		/* prodcast api */
		router.GET("/prodcast", getProdcast)
	}

	router.Run("localhost:8080")

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server Run Failed:")
		log.Panic(err)
	}
}

/**
 * TODO: database connection is under development soon this will be integerated :)
 */
func databaseconnection() {
	database, _ := sql.Open("sqlite3", "./database.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS user (uid INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	statement.Exec()
	statement, _ = database.Prepare("INSERT INTO user (firstname, lastname) VALUES (?, ?)")
	statement.Exec("Krish", "Pranav")
	rows, _ := database.Query("SELECT uid, firstname, lastname FROM user")

	var uid int
	var firstname string
	var lastname string

	for rows.Next() {
		rows.Scan(&uid, &firstname, &lastname)
		fmt.Println(strconv.Itoa(uid) + ": " + firstname + " " + lastname)
	}
}

/* start of get functions */
func getProdcast(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Prodcast)
}

func getRecommendedTracks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Top)
}

func getProfile(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, profiles)
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAbout(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, aboutapp)
}

func getArtistData(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, artists)
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

/* end of get functions */

/* start of post functions */
func postArtist(c *gin.Context) {
	var newArtist artist.Artist

	if err := c.BindJSON(&newArtist); err != nil {
		return
	}

	artists = append(artists, newArtist)
	c.IndentedJSON(http.StatusCreated, newArtist)
}

func postProfile(c *gin.Context) {
	var newProfile profile.Profile

	if err := c.BindJSON(&newProfile); err != nil {
		return
	}

	profiles = append(profiles, newProfile)
	c.IndentedJSON(http.StatusCreated, newProfile)
}

func postNewTopTracks(c *gin.Context) {
	var newTopTracks home.Recommended

	if err := c.BindJSON(&newTopTracks); err != nil {
		return
	}

	Top = append(Top, newTopTracks)
	c.IndentedJSON(http.StatusAccepted, Top)
}

func postAlbums(c *gin.Context) {
	var newAlbum album.Album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

/* end of post functions */

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

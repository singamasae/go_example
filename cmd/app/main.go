package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Album represents data about a record album.
type Album struct {
	Id     string  `json:"Id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []Album{
	{Id: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{Id: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{Id: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	fmt.Print("hello world")
	router.GET("/hello", findAlbums)
	router.Run("localhost:8000")
}

func findAlbums(c *gin.Context) {
	fmt.Println(albums)
	c.IndentedJSON(http.StatusOK, albums)
}

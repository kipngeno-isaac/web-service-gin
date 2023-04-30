package main

import "net/http"
import "github.com/gin-gonic/gin"

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{{ID: "1", Title: "Hurt2bHuman", Artist: "Pink", Price: 17.40}, {ID: "2", Title: "Pink Friday", Artist: "Nicki Minaj", Price: 18.64}, {ID: "3", Title: "Royalty", Artist: "Chris Breezy", Price: 20.04}}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}

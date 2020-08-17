package main

import (
	"gofirebase/api"
	"gofirebase/config"

	"github.com/gin-gonic/gin"
)

func main() {

	// initialize new gin engine (for server)
	r := gin.Default()

	// configure database
	db := config.CreateDatabase()

	// set db to gin context with a middleware to all incoming request
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	// routes definition for finding and creating artists
	r.GET("/artist", api.FindArtists)
	r.POST("/artist", api.CreateArtist)

	// start the server
	r.Run(":5000")

}

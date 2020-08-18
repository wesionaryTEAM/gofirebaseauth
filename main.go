package main

import (
	"gofirebase/api"
	"gofirebase/config"
	"gofirebase/middleware"

	"github.com/gin-gonic/gin"
)

func main() {

	// initialize new gin engine (for server)
	r := gin.Default()

	// configure database
	db := config.CreateDatabase()

	// configure firebase
	firebaseAuth := config.SetupFirebase()

	// set db & firebase auth to gin context with a middleware to all incoming request
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Set("firebaseAuth", firebaseAuth)
	})

	// using the auth middle ware to validate api requests
	r.Use(middleware.AuthMiddleware)

	// routes definition for finding and creating artists
	r.GET("/artist", api.FindArtists)
	r.POST("/artist", api.CreateArtist)

	// start the server
	r.Run(":5000")

}

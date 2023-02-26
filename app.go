package main

import (
	"gdsc/configs"
	"gdsc/routes"

	"github.com/gin-gonic/gin"
)

func main() {
  configs.ConnectDB()
	router := gin.Default()
	configs.ConnectDB()

	v1 := router.Group("/v1")

	routes.BlogRoute(v1)
	// routes.PointRoute(v1)

	router.Run("localhost:8080")
}

//FACEBOOK EASY TO USE HARD TO INSTALL
//GOOGLE = HARD TO USE EASY TO INSTALL
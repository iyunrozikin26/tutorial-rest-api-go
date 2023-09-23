package main

import (
	"github.com/gin-gonic/gin"
	"github.com/iyunrozikin26/tutorial-rest-api-go.git/src/config"
	route "github.com/iyunrozikin26/tutorial-rest-api-go.git/src/routes"
)

func main() {
	router := gin.Default()
	db := config.DB()

	route.Api(router, db)

	router.Run()
}

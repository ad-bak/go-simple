package main

import (
	"github.com/ad-bak/go-react/config"
	"github.com/ad-bak/go-react/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	config.Connect()

	routes.UserRoute(router)

	router.Run(":4000")
}

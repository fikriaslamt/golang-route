package main

import (
	"inventaris/config"
	"inventaris/route"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()
	route.SetupRoutes(r)
	r.Run()
}

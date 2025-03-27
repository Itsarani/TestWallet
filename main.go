package main

import (
	"TestWallet/config"
	"TestWallet/routes"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	config.ConnectDatabase()
	fmt.Println("Database Connected OK")
	r.Use(gin.Logger(), gin.Recovery())
	routes.RegisterRoutes(r)
	r.Run(":8080")
}

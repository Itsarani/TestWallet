package routes

import (
	"TestWallet/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/users", controllers.GetUsers)
	router.POST("/wallet/verify", controllers.VerifyTopUp)
	router.POST("/wallet/confirm", controllers.ConfirmTopUp)
}

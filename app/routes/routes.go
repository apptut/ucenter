package routes

import (
	"github.com/gin-gonic/gin"
	"ucenter/app/controllers/sms"
	"ucenter/app/controllers/user"
	"ucenter/app/middleware"
)

func New() *gin.Engine {
	router := gin.New()
	router.GET("/v1/sms/send", sms.Send)
	router.POST("/v1/user/reg", user.Reg)
	router.POST("/v1/user/login", user.Login)
	router.GET("/v1/user/profile", middleware.Auth(), user.Profile)

	return router
}

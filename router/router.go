package router

import (
	"myapp/controller"

	"github.com/gin-gonic/gin"
)

func ApiRouter(router *gin.Engine) {
	router.POST("/post", controller.PostCreate)
	router.GET("/posts", controller.PostGetAll)
}

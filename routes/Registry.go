package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.Engine) {
	router.Static("/static", "./static/")

	router.GET("/", index)
	router.GET("/blog", blog)
	router.GET("/post/:name", post)
}

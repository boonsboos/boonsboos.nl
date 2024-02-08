package main

import (
	"boonsboos/cms"
	"boonsboos/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	cms.PeriodicallyRefresh(6)

	r := gin.Default()

	routes.RegisterRoutes(r)

	r.Use(cachedMiddleware)

	r.LoadHTMLGlob("./html/**/*.html")

	r.Run(":8088")
}

func cachedMiddleware(c *gin.Context) {
	c.Writer.Header().Add("Cache-Control", "max-age=86400") // cache everything for one day
	c.Next()
}

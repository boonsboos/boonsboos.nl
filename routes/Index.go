package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func index(context *gin.Context) {
	context.HTML(http.StatusOK, "page/Index", gin.H{
		"PageTitle": "boonsboos",
	})
}

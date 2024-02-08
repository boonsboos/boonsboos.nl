package routes

import (
	"boonsboos/cms"
	"net/http"

	"github.com/gin-gonic/gin"
)

func blog(context *gin.Context) {
	context.HTML(http.StatusOK, "page/Blog", gin.H{
		"PageTitle": "boonsboos' blog",
		"Posts":     cms.PostInfoCache,
	})
}

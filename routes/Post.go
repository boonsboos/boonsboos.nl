package routes

import (
	"boonsboos/cms"
	"net/http"

	"github.com/gin-gonic/gin"
)

func post(context *gin.Context) {

	postName := context.Param("name")

	post, err := cms.GetPost(postName)
	if err != nil {
		context.String(http.StatusOK, err.Error())
		return
	}

	context.HTML(http.StatusOK, "page/Post", gin.H{
		"PageTitle": post.Info.Title + " | boonsboos",
		"Post":      post,
	})
}

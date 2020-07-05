package topic

import (
	"github.com/gin-gonic/gin"
	"github.com/shiminjia/myblog/controller"
	"net/http"
)

func ShowCreate(c *gin.Context) {
	categories, err := controller.GetAllCategories()
	if err != nil {
		c.HTML(http.StatusNotFound, "404", gin.H{
			"Error": "系统发生错误，请联系管理员。",
		})
		return
	}

	c.HTML(http.StatusOK, "topic_create", gin.H{
		"IsLogin":    true,
		"Categories": categories,
	})
}

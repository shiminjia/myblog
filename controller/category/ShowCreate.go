package category

import (
	"github.com/gin-gonic/gin"
	"github.com/shiminjia/myblog/dbops"
	"log"
	"net/http"
)

func ShowCreate(c *gin.Context) {
	category := dbops.Category{}
	categories, err := category.GetAllCategories()
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusFound, "404", gin.H{
			"Error": "该页面不存在。",
		})
		return
	}

	c.HTML(http.StatusOK, "category_create", gin.H{
		"IsLogin": true,
		"Categories": categories,
	})
}

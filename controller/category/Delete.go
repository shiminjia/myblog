package category

import (
	"github.com/gin-gonic/gin"
	"github.com/shiminjia/myblog/dbops"
	"log"
	"net/http"
	"strconv"
)

func Delete(c *gin.Context){
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println(err)
		c.Redirect(http.StatusBadRequest, "/")
		return
	}

	category := dbops.Category{
		Id: id,
	}
	rows, err := category.DeleteCategoryById()
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusNotFound, "404", gin.H{
			"Error": "该分类不存在",
		})
		return
	}
	if rows <= 0 {
		log.Println(rows)
		c.HTML(http.StatusNotFound, "404", gin.H{
			"Error": "该分类不存在",
		})
		return
	}

	c.Redirect(http.StatusFound,"/category/create")
	return
}
package topic

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

	topic := dbops.Topic{
		Id: id,
	}
	rows, err := topic.DeleteTopicById()
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusNotFound, "404", gin.H{
			"Error": "该文章不存在",
		})
		return
	}
	if rows <= 0 {
		log.Println(rows)
		c.HTML(http.StatusNotFound, "404", gin.H{
			"Error": "该文章不存在",
		})
		return
	}

	c.Redirect(http.StatusFound,"/")
	return
}
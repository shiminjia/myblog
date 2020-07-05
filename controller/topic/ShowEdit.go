package topic

import (
	"github.com/gin-gonic/gin"
	"github.com/shiminjia/myblog/controller"
	"github.com/shiminjia/myblog/dbops"
	"log"
	"net/http"
	"strconv"
)

func ShowEdit(c *gin.Context){
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		log.Println(err)
		c.Redirect(http.StatusFound, "/")
		return
	}

	topic := &dbops.Topic{
		Id : id,
	}
	topic, err = topic.GetTopicById()
	if err != nil {
		log.Println(err)
		c.Redirect(http.StatusFound, "/")
		return
	}

	categories, err := controller.GetAllCategories()

	c.HTML(http.StatusOK, "topic_edit", gin.H{
		"IsLogin": true,
		"Id": topic.Id,
		"Title": topic.Title,
		"Label": topic.Label,
		"Categories": categories,
		"Content": topic.Content,
	})
}
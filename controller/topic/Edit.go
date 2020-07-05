package topic

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/shiminjia/myblog/controller"
	"github.com/shiminjia/myblog/dbops"
	"github.com/shiminjia/myblog/dto"
	"log"
	"net/http"
)

func Edit(c *gin.Context){
	topic := dto.Topic{}
	err := c.ShouldBind(&topic)
	if err != nil {
		log.Println(err)
		c.Redirect(http.StatusFound,"/")
		return
	}

	categories, err := controller.GetAllCategories()
	if err != nil {
		log.Println(err)
		c.Redirect(http.StatusFound,"/")
		return
	}

	validate := validator.New()
	err = validate.Struct(topic)
	if err != nil {
		c.HTML(http.StatusBadRequest, "topic_edit", gin.H{
			"IsLogin": true,
			"Error": err.Error(),
			"Id": topic.Id,
			"Title": topic.Title,
			"Label": topic.Label,
			"Categories": categories,
			"Content": topic.Content,
		})
	}

	dao := dbops.Topic{
		Id: topic.Id,
		Title: topic.Title,
		Label: topic.Label,
		CategoryId: topic.CategoryId,
		Content: topic.Content,
	}
	id, err := dao.UpdateTopicById()
	if err != nil {
		c.HTML(http.StatusBadRequest, "topic_edit", gin.H{
			"IsLogin": true,
			"Error": err.Error(),
			"Id": topic.Id,
			"Title": topic.Title,
			"Label": topic.Label,
			"Categories": categories,
			"Content": topic.Content,
		})
	}
	log.Printf("create new topic: %d",id)
	c.Redirect(http.StatusFound, "/")
	return
}
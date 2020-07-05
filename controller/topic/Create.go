package topic

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/shiminjia/myblog/dbops"
	"github.com/shiminjia/myblog/dto"
	"log"
	"net/http"
)

func Create(c *gin.Context){
	form := dto.Topic{}
	if err := c.ShouldBind(&form); err != nil {
		log.Println(err);
		c.HTML(http.StatusInternalServerError, "404", gin.H{
			"Error": "系统发生错误，请联系管理员。",
		})
		return
	}

	category := dbops.Category{}
	allCategories, err := category.GetAllCategories()
	if err != nil {
		log.Println(err);
		c.HTML(http.StatusInternalServerError, "404", gin.H{
			"Error": "系统发生错误，请联系管理员。",
		})
		return
	}

	categories := []dto.Category{}
	for i,k := range allCategories {
		categories = append(categories, dto.Category{})
		categories[i].Id = k.Id
		categories[i].Name = k.Name
		categories[i].IsSelected = false
		if form.CategoryId == k.Id {
			categories[i].IsSelected = true
		}
	}

	validate := validator.New()
	err = validate.Struct(form)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusOK, "topic_create", gin.H{
			"IsLogin": true,
			"Error": "内容不符合要求，请检查后重新输入。",
			"Title": form.Title,
			"Label": form.Label,
			"Categories": categories,
			"Content": form.Content,
		})
		return
	}

	dao := dbops.Topic{
		Title: form.Title,
		Label: form.Label,
		CategoryId: form.CategoryId,
		Content: form.Content,
	}
	id, err := dao.CreateTopic()
	if err != nil {
		log.Println(err);
		c.HTML(http.StatusInternalServerError, "404", gin.H{
			"Error": "系统发生错误，请联系管理员。",
		})
		return
	}

	log.Printf("create new topic: %d",id)

	c.Redirect(http.StatusFound,"/")
}
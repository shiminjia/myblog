package category

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/shiminjia/myblog/dbops"
	"github.com/shiminjia/myblog/dto"
	"log"
	"net/http"
)

func Create(c *gin.Context) {
	category := dto.Category{}
	err := c.ShouldBind(&category)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusInternalServerError, "404", gin.H{
			"Error": "发生错误，请稍后重试。",
		})
		return
	}

	dao := dbops.Category{
		Name: category.Name,
	}

	// 校验名称
	validate := validator.New()
	err = validate.Struct(category)
	if err != nil {
		log.Println(err)
		categories, err := dao.GetAllCategories()
		if err != nil {
			log.Println(err)
			c.HTML(http.StatusInternalServerError, "404", gin.H{
				"Error": "发生错误，请稍后重试。",
			})
			return
		}
		c.HTML(http.StatusBadRequest, "category_create", gin.H{
			"IsLogin":    true,
			"Error":      "输入的分类不符合要求，请重新输入。",
			"Name":       category.Name,
			"Categories": categories,
		})
		return
	}

	// 校验总数不超过10
	amount, err := dao.GetCategoryAmount()
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusInternalServerError, "404", gin.H{
			"Error": "发生错误，请稍后重试。",
		})
		return
	}
	if amount > 10 {
		categories, err := dao.GetAllCategories()
		if err != nil {
			log.Println(err)
			c.HTML(http.StatusInternalServerError, "404", gin.H{
				"Error": "发生错误，请稍后重试。",
			})
			return
		}
		c.HTML(http.StatusForbidden, "category_create", gin.H{
			"IsLogin":    true,
			"Error":      "分类最多创建10条。",
			"Name":       category.Name,
			"Categories": categories,
		})
		return
	}

	// 校验没有相同的分类
	count, err := dao.GetCategoryCountByName()
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusInternalServerError, "404", gin.H{
			"Error": "发生错误，请稍后重试。",
		})
		return
	}
	if count > 0 {
		categories, err := dao.GetAllCategories()
		if err != nil {
			log.Println(err)
			c.HTML(http.StatusInternalServerError, "404", gin.H{
				"Error": "发生错误，请稍后重试。",
			})
			return
		}
		c.HTML(http.StatusForbidden, "category_create", gin.H{
			"IsLogin":    true,
			"Error":      "该分类已经存在。",
			"Name":       category.Name,
			"Categories": categories,
		})
		return
	}

	// 创建
	id, err := dao.Create()
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusInternalServerError, "404", gin.H{
			"Error": "发生错误，请稍后重试。",
		})
		return
	}
	log.Printf("create new category: %d", id)

	c.Redirect(http.StatusFound, "/category/create")
	return
}

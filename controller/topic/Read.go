package topic

import (
	"github.com/gin-gonic/gin"
	"github.com/shiminjia/myblog/controller"
	"log"
	"net/http"
	"strconv"
)

func Read(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusNotFound, "404", gin.H{
			"Error": "该文章不存在。",
		})
		return
	}

	//判断是否登录
	_, b := controller.CheckSession(c)
	isLogin := false
	if b {
		isLogin = true
	}

	//获取topic
	dto, err := controller.GetOneTopicById(id)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusNotFound, "404", gin.H{
			"Error": "该文章不存在。",
		})
		return
	}

	c.HTML(http.StatusOK, "topic_read", gin.H{
		"IsLogin":      isLogin,
		"Topic":        dto,
	})

	return
}

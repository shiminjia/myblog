package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/shiminjia/myblog/dbops"
	"github.com/shiminjia/myblog/dto"
	"log"
	"net/http"
	"strconv"
)

type pageNav struct {
	Id          int
	Link        string
	CurrentPage bool
}

func Index(c *gin.Context) {
	// 验证session
	_, isLogin := CheckSession(c)

	// 获取参数page
	page := c.Query("page")
	if page == "" {
		page = "1"
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusNotFound, "404", gin.H{
			"Error": "该页面不存在。",
		})
		return
	}

	// 获取参数categoryId
	categoryId := c.Query("categoryid")
	if categoryId == "" {
		categoryId = "0"
	}
	categoryIdInt, err := strconv.Atoi(categoryId)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusNotFound, "404", gin.H{
			"Error": "该页面不存在。",
		})
		return
	}

	// 获取topics部分内容
	var topics []dto.Topic
	// 如果categoryIdInt为0，则获取所有分类的文章
	// 如果categoryIdInt不为0，则获取该分类的文章
	topics, err = GetSomeTopicsByCategoryId(categoryIdInt, pageInt, 10)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusNotFound, "404", gin.H{
			"Error": "系统发生错误，请联系管理员。",
		})
		return
	}

	// 获取topics的总数
	var topicDao dbops.Topic
	topicAmount, err := topicDao.GetTopicAmount(categoryIdInt)
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusNotFound, "404", gin.H{
			"Error": "系统发生错误，请联系管理员。",
		})
		return
	}
	var pageAmount int64
	if topicAmount%10 == 0 {
		pageAmount = (topicAmount - topicAmount%10) / 10
	} else {
		pageAmount = (topicAmount-topicAmount%10)/10 + 1
	}

	// 获取所有分类
	categories, err := GetAllCategories()
	if err != nil {
		log.Println(err)
		c.HTML(http.StatusNotFound, "404", gin.H{
			"Error": "系统发生错误，请联系管理员。",
		})
		return
	}

	// 获取分页数据
	isFirstPage := false
	isLastPage := false
	if pageInt == 1 {
		isFirstPage = true
	}
	if pageInt == int(pageAmount) || pageAmount == 0 {
		isLastPage = true
	}

	var pageNavArray []pageNav
	for i := 1; i <= int(pageAmount); i++ {
		pageNavArray = append(pageNavArray, pageNav{
			Id:          i,
			Link:        "/index?page=" + strconv.Itoa(i) + "&categoryId=" + strconv.Itoa(categoryIdInt),
			CurrentPage: pageInt == i,
		})
	}

	previousPageLink := "/index?page=" + strconv.Itoa(pageInt-1) + "&categoryId=" + strconv.Itoa(categoryIdInt)
	nextPageLink := "/index?page=" + strconv.Itoa(pageInt+1) + "&categoryId=" + strconv.Itoa(categoryIdInt)

	// 返回结果
	c.HTML(http.StatusOK, "index", gin.H{
		"IsLogin": isLogin,

		"IsFirstPage":       isFirstPage,
		"IsLastPage":        isLastPage,
		"CurrentPageNumber": page,
		"PreviousPageLink":  previousPageLink,
		"NextPageLink":      nextPageLink,
		"PageAmount":        pageAmount,
		"PageNavArray":      pageNavArray,

		"Categories": categories,
		"Topics":     topics,
	})
}

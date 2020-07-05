package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/shiminjia/myblog/dbops"
	"github.com/shiminjia/myblog/dto"
	"log"
)

func CheckSession(c *gin.Context) (*dbops.Value,bool){
	isLogin := false

	sessionId, err := c.Cookie("myblog_cookie")
	if err != nil {
		log.Println(err)
		return nil, isLogin
	}

	session, err := dbops.GetValue(sessionId)
	if err != nil {
		log.Println(err)
		return nil, isLogin
	}

	var value dbops.Value
	err = json.Unmarshal(session, &value)
	if err != nil {
		log.Println(err)
		return nil, isLogin
	}

	if value.Id < 0  {
		log.Println(err)
		return nil, isLogin
	}

	isLogin = true
	return &value, isLogin
}

func GetOneTopicById(id int)(*dto.Topic, error){
	dao := &dbops.Topic{
		Id:   id,
		Time: &dbops.Time{},
	}

	dao, err := dao.GetOneTopicById()
	if err != nil {
		return nil, err
	}
	time := &dto.Time{
		CreatedAt: dao.Time.CreatedAt,
		UpdatedAt: dao.Time.UpdatedAt,
		DeletedAt: dao.Time.DeletedAt,
	}

	dto := &dto.Topic{
		Id:           dao.Id,
		Title:        dao.Title,
		Label:        dao.Label,
		CategoryId:   dao.CategoryId,
		CategoryName: dao.CategoryName,
		Content:      dao.Content,
		Time:         time,
	}
	return dto, nil
}

func GetSomeTopicsByCategoryId(categoryIdInt int, page int, limit int) ([]dto.Topic, error){
	dao := dbops.Topic{}
	daos, err := dao.GetSomeTopicsByCategoryId(categoryIdInt,page, limit)
	if err != nil {
		return nil, err
	}

	dtos := []dto.Topic{}
	for _, k := range daos {
		Time := &dto.Time{
			CreatedAt: k.Time.CreatedAt,
			UpdatedAt: k.Time.UpdatedAt,
			DeletedAt: k.Time.DeletedAt,
		}
		dto := dto.Topic{
			Id:         k.Id,
			Title:      k.Title,
			Label:      k.Label,
			CategoryId: k.CategoryId,
			CategoryName: k.CategoryName,
			Content:    k.Content,
			Time:       Time,
		}
		dtos = append(dtos, dto)
	}
	return dtos, nil
}

//func GetSomeTopics(page int, limit int) ([]dto.Topic, error){
//	categoryIdInt := 0
//	dtos, err := GetSomeTopicsByCategoryId(categoryIdInt, page, limit)
//	if err != nil {
//		return nil, err
//	}
//	return dtos, nil
//}

func GetAllCategories() ([]dto.Category,error) {
	category := dbops.Category{}
	allCategories, err := category.GetAllCategories()
	if err != nil {
		log.Println(err);
		return nil, err
	}

	categories := []dto.Category{}
	for i,k := range allCategories {
		categories = append(categories, dto.Category{})
		categories[i].Id = k.Id
		categories[i].Name = k.Name
		categories[i].IsSelected = false
	}
	return categories, nil
}

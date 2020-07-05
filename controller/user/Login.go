package user

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/shiminjia/myblog/dbops"
	"github.com/shiminjia/myblog/dto"
	"log"
	"net/http"
)

func Login(c *gin.Context) {
	var form dto.Login

	//bind struct
	if err := c.ShouldBind(&form); err != nil {
		log.Println(err)
		c.HTML(http.StatusOK, "login", gin.H{
			"Email": form.Email,
			"Error": "发生错误，请重新登录",
		})
		return
	}

	//Parameter check
	validate := validator.New()
	err := validate.Struct(form)
	if err != nil {
		c.HTML(http.StatusOK, "login", gin.H{
			"Email": form.Email,
			"Error": err.Error(),
		})
		return
	}

	//check email and password and get username
	dao := dbops.SelectUserByEmail(form.Email)

	valueJson, err := json.Marshal(dbops.Value{
		Id: dao.Id,
		Email: dao.Email,
		Username: dao.Username,
	})
	if err != nil {
		log.Println(err)
	}

	//auth email ane password
	if form.Password == dao.Password {
		sessionId := uuid.New().String()
		dbops.SetValue(sessionId, valueJson)
		if form.Check == "true" {
			c.SetCookie("myblog_cookie", sessionId, 60*60*24, "/", "localhost", false, true)
		} else {
			c.SetCookie("myblog_cookie", sessionId, 60*60, "/", "localhost", false, true)
		}
		c.Redirect(http.StatusFound, "/")
		return
	} else {
		c.HTML(http.StatusOK, "login", gin.H{
			"Email": form.Email,
			"Error": "用户名或密码错误",
		})
		return
	}
}

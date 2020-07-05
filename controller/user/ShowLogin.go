package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowLogin(c *gin.Context){
	c.HTML(http.StatusOK, "login", gin.H{})
}

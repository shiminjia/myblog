package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Logout(c *gin.Context) {
	c.SetCookie("myblog_cookie", "", 0, "/", "localhost", false, true)
	c.Redirect(http.StatusFound,"/")
}

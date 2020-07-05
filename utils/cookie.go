package utils

import "github.com/gin-gonic/gin"

func SetCookie(c *gin.Context, value string, time int) {
	c.SetCookie("myblog_cookie", value, time, "/", "localhost", false, true)
}

func ClearCookie(c *gin.Context) {
	c.SetCookie("myblog_cookie", "", 0, "/", "localhost", false, true)
}


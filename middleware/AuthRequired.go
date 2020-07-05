package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/shiminjia/myblog/dbops"
	"log"
	"net/http"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionId, err := c.Cookie("myblog_cookie")
		if err != nil {
			log.Println(err)
			c.HTML(http.StatusForbidden, "404", gin.H{
				"Error": "没有权限",
			})
			c.Abort()
			return
		}

		valueJson, err := dbops.GetValue(sessionId)
		if err != nil {
			log.Println(err)
			c.HTML(http.StatusForbidden, "404", gin.H{
				"Error": "Session不存在1",
			})
			c.Abort()
			return
		}

		var value dbops.Value
		err = json.Unmarshal(valueJson, &value)
		if err != nil {
			log.Println(err)
			c.HTML(http.StatusForbidden, "404", gin.H{
				"Error": "Session不存在2",
			})
			c.Abort()
			return
		}

		if value.Id == 0 {
			c.HTML(http.StatusForbidden, "404", gin.H{
				"Error": "Session不存在3",
			})
			c.Abort()
			return
		}

		c.Set("value",value)

		c.Next()
	}
}
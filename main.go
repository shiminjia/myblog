package main

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/shiminjia/myblog/controller"
	"github.com/shiminjia/myblog/controller/category"
	"github.com/shiminjia/myblog/controller/topic"
	"github.com/shiminjia/myblog/controller/user"
	"github.com/shiminjia/myblog/dbops"
	"github.com/shiminjia/myblog/middleware"
)

func createMyRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("index", "templates/index.tmpl",
		"templates/components/T.nav.tmpl", "templates/components/T.footer.tmpl","templates/components/T.msg.tmpl")
	r.AddFromFiles("login", "templates/login.tmpl","templates/components/T.msg.tmpl")
	r.AddFromFiles("topic_create", "templates/topic_create.tmpl",
		"templates/components/T.nav.tmpl", "templates/components/T.footer.tmpl","templates/components/T.msg.tmpl")
	r.AddFromFiles("topic_read", "templates/topic_read.tmpl",
		"templates/components/T.nav.tmpl", "templates/components/T.footer.tmpl","templates/components/T.msg.tmpl")
	r.AddFromFiles("topic_edit", "templates/topic_edit.tmpl",
		"templates/components/T.nav.tmpl", "templates/components/T.footer.tmpl","templates/components/T.msg.tmpl")
	r.AddFromFiles("category_create", "templates/category_create.tmpl",
		"templates/components/T.nav.tmpl", "templates/components/T.footer.tmpl","templates/components/T.msg.tmpl")
	r.AddFromFiles("404", "templates/404.tmpl")
	return r
}

func main() {
	r := gin.Default()

	//导入静态文件
	r.Static("/assets", "./assets")
	//设置模版文件夹
	//r.LoadHTMLGlob("templates/*")

	//设置多个模版文件夹
	r.HTMLRender = createMyRender()

	dbops.InitDB()
	dbops.InitRedis()

	r.GET("/", controller.Index)
	r.GET("/index", controller.Index)

	r.GET("/login", user.ShowLogin)
	r.POST("/login", user.Login)

	r.GET("/topic/read/:id", topic.Read)

	authorized := r.Group("/")
	authorized.Use(middleware.AuthRequired())
	{
		authorized.GET("/logout", user.Logout)

		authorized.GET("/category/create", category.ShowCreate)
		authorized.POST("/category/create", category.Create)
		authorized.GET("/category/delete/:id", category.Delete)

		authorized.GET("/topic/create", topic.ShowCreate)
		authorized.POST("/topic/create", topic.Create)
		authorized.GET("/topic/edit/:id", topic.ShowEdit)
		authorized.POST("/topic/edit", topic.Edit)
		authorized.GET("/topic/delete/:id", topic.Delete)
	}

	r.Run()
}

//todo list
//日志      log => zap
//参数配置   viper
//多语言

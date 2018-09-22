/*
* @Author: Newbie Coder
* @Date:   2018-09-20 10:09:43
* @Last Modified by:   Newbie Coder
* @Last Modified time: 2018-09-22 13:38:48
*/
package server

import (
	"github.com/gin-gonic/gin"
	"github.com/foolin/gin-template"
	"github.com/NewbieCoder99/go-night-framework/controllers"
)

func NewRouter() *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Static("img", "resources/assets/img")
	router.Static("css", "resources/assets/css")
	router.Static("js", "resources/assets/js")

	router.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root		: "resources/views",
		Extension 	: ".html",
		Master 		: "layouts/master",
		DisableCache: false,
	})

	// Make Route Here
	router.GET("/", new(controllers.WelComeController).Welcome)
	regist := router.Group("register")
	{
		regist.GET("", new(controllers.RegisterController).Register)
		regist.POST("", new(controllers.RegisterController).RegisterPost)
	}

	return router
}
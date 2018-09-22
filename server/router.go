/*
* @Author: Newbie Coder
* @Date:   2018-09-20 10:09:43
* @Last Modified by:   Newbie Coder
* @Last Modified time: 2018-09-22 10:09:38
*/
package server

import (
	"github.com/gin-gonic/gin"
	"github.com/NewbieCoder99/go-night-framework/controllers"
	"github.com/foolin/gin-template"
)

func NewRouter() *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Static("img", "resources/assets/img")
	router.Static("css", "resources/assets/css")
	router.Static("js", "resources/assets/js")

	//new template engine
	router.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root		: "resources/views",
		Extension 	: ".html",
		Master 		: "layouts/master",
		// Partials 	: []string{"partials/.html"},
		DisableCache: false,
	})

	// Make Route Here
	Wel := new(controllers.WelComeController)
	router.GET("/", Wel.Welcome)

	return router
}
/*
* @Author: Newbie Coder
* @Date:   2018-09-20 10:09:43
* @Last Modified by:   Newbie Coder
* @Last Modified time: 2018-09-22 17:10:08
*/
package controllers

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/NewbieCoder99/go-night-framework/config"
)

type RegisterController struct { }

type UserFields struct {

	Email     	string `form:"email" binding:"email,required"`
	Name 		string `form:"name" binding:"required"`
	Password   	string `form:"password" binding:"required"`
	Passconf 	string `form:"passconf" binding:"required"`

}

func(h RegisterController) Register(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "register", gin.H {
		"base_url" : config.BaseUrl(),
		"datenow" : time.Now().Format("2006"),
		"title": "Register Account",
	})
}

func(h RegisterController) RegisterPost(ctx *gin.Context) {

	var fields UserFields

	if err := ctx.ShouldBindWith(&fields, binding.Form); err == nil {
		ctx.JSON(200, gin.H{"message": "User founded!", "user": ctx.PostForm("name")})
	} else {
		ctx.JSON(403, gin.H{ "error": err } )
	}
}

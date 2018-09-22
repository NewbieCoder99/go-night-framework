/*
* @Author: Newbie Coder
* @Date:   2018-09-20 10:09:43
* @Last Modified by:   Newbie Coder
* @Last Modified time: 2018-09-22 13:22:08
*/
package controllers

import (
	"net/http"
	"time"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/NewbieCoder99/go-night-framework/config"
)

type RegisterController struct { }

func (h RegisterController) Register(c *gin.Context) {

	c.HTML(http.StatusOK, "register", gin.H {
		"base_url" : config.BaseUrl(),
		"datenow" : time.Now().Format("2006"),
		"title": "Register Account",
	})
}

func (h RegisterController) RegisterPost(c *gin.Context) {
	fmt.Println(c)
}

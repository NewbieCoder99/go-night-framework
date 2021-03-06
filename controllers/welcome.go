/*
* @Author: Newbie Coder
* @Date:   2018-09-20 10:09:43
* @Last Modified by:   Newbie Coder
* @Last Modified time: 2018-09-22 13:32:31
*/
package controllers

import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/NewbieCoder99/go-night-framework/config"
)

type WelComeController struct { }

func (h WelComeController) Welcome(c *gin.Context) {

	c.HTML(http.StatusOK, "index", gin.H {
		"base_url" : config.BaseUrl(),
		"datenow" : time.Now().Format("2006"),
		"title": "Welcome My Controller",
	})
}

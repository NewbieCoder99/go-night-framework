/*
* @Author: Newbie Coder
* @Date:   2018-09-20 10:09:43
* @Last Modified by:   Newbie Coder
* @Last Modified time: 2018-09-22 07:37:09
*/
package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type WelComeController struct{}

func (h WelComeController) Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "index", gin.H {
		"title": "Welcome My Controller",
	})
}

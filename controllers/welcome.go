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

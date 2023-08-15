package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func showMainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "main.html", nil)
}

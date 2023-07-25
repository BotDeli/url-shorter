package handlers

import "github.com/gin-gonic/gin"

func showMainPage(c *gin.Context) {
	c.HTML(200, "main_page.html", nil)
}

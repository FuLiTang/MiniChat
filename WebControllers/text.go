package WebControllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TextHome(c *gin.Context) {
	c.HTML(http.StatusOK, "text.html", gin.H{})
}

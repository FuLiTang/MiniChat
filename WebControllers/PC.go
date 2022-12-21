package WebControllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PcIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "pcIndex.html", gin.H{})
}
func PcHome(c *gin.Context) {
	c.HTML(http.StatusOK, "pcHome.html", gin.H{})
}

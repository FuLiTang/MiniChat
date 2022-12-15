package WebControllers

import (
	"github.com/gin-gonic/gin"
	"miniChat/models"
	"net/http"
)

func Home(c *gin.Context) {
	user := &models.UserBasic{}
	c.Bind(user)
	if !user.QueryAndIdentity() {
		c.JSON(200, gin.H{
			"code":    "0",
			"message": "系统错误！",
		})
		return
	}
	c.HTML(http.StatusOK, "home.html", gin.H{
		"user": user.FirstName(),
	})
}

func MiniChat(c *gin.Context) {
	user := &models.UserBasic{}
	c.Bind(user)
	client := &models.UserBasic{Name: c.Query("client")}
	if !user.QueryAndIdentity() {
		c.JSON(200, gin.H{
			"code":    "0",
			"message": "系统错误！",
		})
		return
	}

	c.HTML(http.StatusOK, "miniChat.html", gin.H{
		"user":   user.FirstName(),
		"client": client.FirstName(),
	})
}

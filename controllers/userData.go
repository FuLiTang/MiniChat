package controllers

import (
	"github.com/gin-gonic/gin"
	"miniChat/models"
)

// UserDataImage 更换头像
func UserDataImage(c *gin.Context) {
	//返回提示
	//if c.PostForm("base") == "" {
	//	c.JSON(200, gin.H{
	//		"code": 0,
	//		"msg":  "",
	//	})
	//	return
	//}
	user := &models.UserBasic{}
	c.Bind(user)
	base := c.PostForm("base")
	if models.UpdateUser(user.FirstName(), models.UserBasic{HeadPortrait: base}).RowsAffected > 0 {
		c.JSON(200, gin.H{
			"code": 0,
		})
	}
	c.JSON(200, gin.H{
		"code": 44,
	})

}

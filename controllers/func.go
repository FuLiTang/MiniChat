package controllers

import (
	"github.com/gin-gonic/gin"
	"miniChat/models"
	"miniChat/utils"
)

// Verify 验证真实性
func Verify(c *gin.Context) *models.UserBasic {
	user := &models.UserBasic{}
	c.Bind(user)
	if user.QueryAndIdentity() {
		return user
	}
	return nil
}
func JsonReturn(c *gin.Context, s string) {
	c.JSON(200, gin.H{
		"code": 1,
		"data": s,
	})
}

// Off 退出登录
func Off(c *gin.Context) {
	user := &models.UserBasic{}
	c.Bind(user)
	if user.QueryAndIdentity() {
		if models.UpdateUser(user, &models.UserBasic{Identity: utils.SaltTime()}).RowsAffected > 0 {
			JsonReturn(c, "已退出")
		}
	}
}

// AddFriends 搜素好友
func AddFriends(c *gin.Context) {
	//A,B,ID
	user := Verify(c)
	if user == nil {
		c.JSON(404, gin.H{})
		return
	}
	client := &models.UserBasic{Name: c.Query("client")}
	if client = client.FirstName(); client == nil {
		JsonReturn(c, "未找到此联系人")
		return
	}
	r := &models.Relation{
		A: user.Name,
		B: client.Name,
	}
	r.Query()
	//0：还不是好友 1：好友状态，2，3：B删了A，4：A删了B，5：B拉黑了A，6：A拉黑了B，7：A发起好友申请，8：B发起了好友申请
	switch r.Type {
	case 0:
		JsonReturn(c, "未成为好友")
	case 1:
		JsonReturn(c, "已是好友")
	case 3:
		JsonReturn(c, "已被删除")
	case 4:
		JsonReturn(c, "已删除了它")
	case 5:
		JsonReturn(c, "已被拉黑")
	case 6:
		JsonReturn(c, "已拉黑了它")
	case 7:
		JsonReturn(c, "我发起了好友申请")
	case 8:
		JsonReturn(c, "对方发起了好友申请")
	}

}

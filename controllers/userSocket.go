package controllers

import (
	"github.com/gin-gonic/gin"
	"miniChat/models"
)

func OffLineMessage(c *gin.Context) {
	user := &models.UserBasic{}
	c.Bind(user)
	if !user.QueryAndIdentity() {
		return
	}
	//offMsUser := &models.NewUser{
	//	SenderName:    user.Name,
	//	RecipientName: "",
	//	Context:       "",
	//	Time:          "",
	//	Type:          "",
	//	Result:        nil,
	//	CBack:         nil,
	//	Mutex:         nil,
	//	Sum:           nil,
	//	Wg:            nil,
	//}

}

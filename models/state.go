package models

import "miniChat/utils"

// OnOffLogin 登录状态
func (t *NewUser) OnOffLogin() bool {
	v, _ := utils.RDb.PubSubChannels(t.CBack, t.RecipientName+"&"+t.SenderName).Result()
	if len(v) == 0 {
		return false
	}
	return true
}

package models

import (
	"context"
	"log"
	"miniChat/utils"
	"sync"
)

type NewUser struct {
	SenderName    string `json:"senderName"`
	RecipientName string `json:"recipientName"`
	Context       string `json:"context"`
	Time          string `json:"time"`
	Type          string `json:"type"`
	Result        chan map[string]string
	CBack         context.Context
	Mutex         *sync.Mutex
	Sum           []string
	Wg            *sync.WaitGroup
}

func (this *NewUser) RedisSAdd(s string) {
	//err := utils.RDb.RPush(context.Background(), this.SenderName+"&"+this.RecipientName, s).Err()
	err := utils.RDb.SAdd(this.CBack, this.SenderName+"&"+this.RecipientName, s).Err()
	if err != nil {
		log.Println(err)
	}
}
func (this *NewUser) RedisSMembers() {
	this.Sum, _ = utils.RDb.Keys(this.CBack, "*&"+this.SenderName).Result()
	vLen := len(this.Sum)
	if vLen == 0 {
		return
	}
	this.Wg.Add(vLen)
	for _, v := range this.Sum {
		go this.redisSMembersMini(v)
	}
	this.Wg.Wait()
}
func (this *NewUser) redisSMembersMini(s string) {
	log.Println("进入")
	vData, _ := utils.RDb.LRange(this.CBack, s, 0, -1).Result()
	for _, v := range vData {
		this.Mutex.Lock()
		this.Result <- map[string]string{s: v}
		this.Mutex.Unlock()
	}
	this.Wg.Done()
}

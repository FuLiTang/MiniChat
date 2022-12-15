package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"miniChat/models"
	"miniChat/utils"
	"net/http"
)

type PodAll struct {
	SendContent   chan *models.NewUser //存web端发来的信息
	AcceptContent chan *models.NewUser //发送给web端的信息
}

var P = PodAll{
	SendContent: make(chan *models.NewUser, 1024),
}

func init() {
	go P.redisSend()
}

//接收web发送的信息
func (this *PodAll) redisSend() {
	for {
		select {
		case ms := <-this.SendContent:
			if !ms.OnOffLogin() {
				//处理离线信息
				return
			}
			go textRedisSend(ms)
		}
	}
}
func textRedisSend(ms *models.NewUser) {
	switch ms.Type {
	case "1":
		{
			utils.RDb.Publish(context.Background(), ms.SenderName+"&"+ms.RecipientName, ms.Time+"&"+ms.Context)
		}
	case "2":
		log.Println("正在存储图片")

	}
}
func ClientSocket(c *gin.Context) {
	message := &models.Message{}
	c.Bind(message)
	_, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}

}
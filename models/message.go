package models

import (
	"github.com/gorilla/websocket"
)

type Message struct {
	Sender        string `form:"sender"`    //发送者,redis位置发送和接收
	Recipient     string `form:"recipient"` //接受者
	Type          int    `form:"type"`      //消息类型
	Ws            *websocket.Conn
	SendContent   chan []byte
	AcceptContent chan []byte
	Bool          bool
}

//func Create(c *Message) {
//	utils.RDb.Send("multi")
//	utils.RDb.Send("lpush", c.Sender+c.Recipient, time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:06")+"&"+string(c.Content))
//	utils.RDb.Do("exec")
//
//}
//func Query(c *Message) (string, error) {
//	v, err := redis.String(utils.RDb.Do("rpop", c.Recipient+c.Sender))
//	if err != nil {
//		log.Println(err)
//		return v, err
//	}
//	return v, err
//}
//func Delete(c *Message) string {
//	v, _ := redis.String(utils.RDb.Do("lpop", c.Recipient+c.Sender))
//	return v
//}

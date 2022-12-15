package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"miniChat/utils"
	"net/http"
	"sync"
	"time"
)

type AllPod struct {
}
type sendContent struct {
	Name    string `json:"name"`
	Context string `json:"context"`
	Time    string `json:"time"`
	Type    string `json:"type"`
}
type acceptContent struct {
	Name    string `json:"name"`
	Context string `json:"context"`
	Time    string `json:"time"`
	Type    string `json:"type"`
}
type user struct {
	Sender    string `form:"sender"`    //发送者,redis位置发送和接收
	Recipient string `form:"recipient"` //接受者
	Type      int    `form:"type"`      //消息类型
	Bool      bool
	Lock      *sync.RWMutex
}
type pod struct { //个人消息池
	SendContent    chan *sendContent //接收web端
	AcceptContent  chan *acceptContent
	ws             *websocket.Conn
	onMessageCount map[string]int
	*user
}
type onMessage struct {
	NoPod chan *pod
}

var onePod struct {
	UserId   string
	Amessage []byte
	Bmessage []byte
}
var onMes = onMessage{
	NoPod: make(chan *pod, 2024),
}

func Socket(c *gin.Context) {
	ms := &user{}
	c.Bind(ms)
	var upWebS = &websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}}
	ws, err := upWebS.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}

	var p = pod{
		SendContent:   make(chan *sendContent, 1024),
		AcceptContent: make(chan *acceptContent, 1024),
		ws:            ws,
		user:          ms,
	}
	//p.ws.WriteMessage(1, []byte("连接成功！"))

	go p.send()
	go p.redisSend()
	go p.accept()
	go p.redisAccept()
	go p.QueryLink()
}
func (p *pod) send() {
	for {
		ms := &sendContent{}
		err := p.ws.ReadJSON(ms)
		if err != nil {
			p.ws.Close()
			return
		}
		//r := &models.Relation{
		//	A:    p.Sender,
		//	B:    ms.Name,
		//	Type: 1,
		//}
		//if r.Query() {
		//	continue
		//}
		ms.Time = time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:06")
		//fmt.Println("接收记录！", ms)
		p.SendContent <- ms
	}
}
func (p *pod) accept() {
	for {
		select {
		case ms := <-p.AcceptContent:
			err := p.ws.WriteJSON(ms)
			if err != nil {
				log.Println(err)
				return
			}
			log.Println("处理成功:", ms)
		}

	}
}

func (p *pod) redisSend() {
	var ctx = context.Background()
	for {
		select {
		case ms := <-p.SendContent:
			utils.RDb.Publish(ctx, p.Sender+"&"+ms.Name, ms.Time+"&"+ms.Context)
		}
	}
}
func (p *pod) redisAccept() {
	t := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:06")
	var ctx = context.Background()
	target := p.Recipient + "&" + p.Sender
	sub := utils.RDb.Subscribe(ctx, target)
	log.Println(p.Sender+"查看订阅", sub)
	for ms := range sub.Channel() {
		var wg sync.WaitGroup
		wg.Add(1)
		ac := &acceptContent{
			Name:    p.Recipient,
			Context: ms.Payload[len(t)+1:],
			Time:    ms.Payload[0:len(t)],
			Type:    "1", //临时，后期用户自动变量
		}
		log.Println("ac:", ac)
		p.AcceptContent <- ac
		wg.Done()
		wg.Wait()
	}
}
func (p *pod) redisAcceptAll() {
	t := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:06")
	var ctx = context.Background()
	target := "*" + "&" + p.Sender
	sub := utils.RDb.PSubscribe(ctx, target)
	for ms := range sub.Channel() {
		var wg sync.WaitGroup
		wg.Add(1)
		name := ms.Channel[0 : len(ms.Channel)-len(p.user.Sender)-1]
		p.AcceptContent <- &acceptContent{
			Name:    name,
			Context: ms.Payload[len(t)+1:],
			Time:    ms.Payload[0:len(t)],
		}
		p.onMessageCount = map[string]int{name: p.onMessageCount[name] + 1}
		wg.Done()
		wg.Wait()
	}
}
func SocketAll(c *gin.Context) {
	ms := &user{}
	c.Bind(ms)
	var upWebS = &websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}}
	ws, err := upWebS.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}

	var p = pod{
		AcceptContent: make(chan *acceptContent, 1024),
		ws:            ws,
		user:          ms,
	}
	go p.accept()
	go p.redisAcceptAll()
}

//	func MessageCount(c *gin.Context) {
//		clent := c.Query("client")
//
// }
func (p *pod) QueryLink() {
	onlink, err := utils.RDb.PubSubChannels(context.Background(), "*").Result()
	if err != nil {
		log.Println("错误", err)
	}
	for _, v := range onlink {
		log.Println("在线的channel::", v)
	}
}

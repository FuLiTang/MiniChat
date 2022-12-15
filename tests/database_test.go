package tests

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"log"
	"miniChat/models"
	"miniChat/utils"
	"strconv"
	"sync"
	"testing"
)

func TestCreateTable(t *testing.T) {
	if err := utils.MDb.AutoMigrate(&models.UserBasic{}, &models.State{}, &models.Relation{}); err != nil {
		log.Println(err)
	}
}
func TestName(t *testing.T) {
	r := &models.Relation{
		Model: gorm.Model{},
		A:     "小王",
		B:     "312",
	}
	fmt.Println(r.Create())
}

func TestQuery(t *testing.T) {
	r := &models.Relation{
		Model: gorm.Model{},
		A:     "123",
		B:     "312",
	}
	fmt.Println(r.Query())
}

//	func TestRedis(t *testing.T) {
//		//m := &models.Message{
//		//	Sender:    "asd",
//		//	Recipient: "das",
//		//	Type:      1,
//		//	Content:   []byte("123"),
//		//}
//		//models.Create(m)
//		log.Println(models.Query(&models.Message{
//			Sender:    "b",
//			Recipient: "a",
//			Type:      0,
//			Content:   nil,
//		}))
//	}
func TestMysqlStates(t *testing.T) {
	s := &models.State{
		UserName: "asdfgg",
		StateId:  "123",
	}

	//fmt.Println(s.Query())

	log.Println(s.Update())
}

func TestRelation(t *testing.T) {

	for i := 20; i < 400; i++ {
		r := &models.Relation{
			A:    "asdfgg",
			B:    "asd" + strconv.Itoa(i),
			Type: 1,
		}
		r.Create()
	}
	//r := &models.Relation{
	//	A:    "qwe",
	//	Type: 1,
	//}
	//v, k := r.QueryPage(3, 10)
	//fmt.Println("一共记录：", k.RowsAffected)
	//for k, v := range v {
	//	log.Println(k)
	//	log.Println(v)
	//}
}
func TestMYsql(t *testing.T) {

}

func TestUser(t *testing.T) {
	user := &models.UserBasic{
		Name:     "asdfgg",
		Password: "123",
		Identity: "F6ACB2FB80123445E2AD5C3E19AB655F",
	}

	z := &ls{
		Ws:   sync.WaitGroup{},
		Wi:   sync.Mutex{},
		User: user,
		C:    make(chan int, 102400000),
	}
	z.C <- 0

	go func() {
		for i := 1; i < 20*5000; i++ {
			z.C <- i
		}
	}()

	z.Ws.Add(1)
	for i := 0; i < 400; i++ {
		go z.rdb()
	}
	z.Ws.Wait()
	log.Println(z.E)
	log.Println("完成！", z.C)
}

type ls struct {
	Ws   sync.WaitGroup
	Wi   sync.Mutex
	User *models.UserBasic
	C    chan int
	E    error
}

func (this *ls) ll() {
	for i := 0; i < 400; i++ {
		log.Println(this.User.FirstName())
		//this.Wi.Lock()
		//this.C++
		//this.Wi.Unlock()
	}
	this.Ws.Done()
}
func (this *ls) rdb() {
	for {
		select {
		case i := <-this.C:
			v := utils.RDb.LPush(context.Background(), "asdffr14", i)
			if v.Err() != nil {
				log.Println(v.Err())
			}
			if i == 20*5000-1 {
				this.Ws.Done()
			}
		}

	}
}
func TestImageCRUD(t *testing.T) {
	user := &models.UserBasic{Name: "asd20"}
	log.Println(len(user.FirstName().HeadPortrait))
}
func TestRedisCRUD(t *testing.T) {
	user := &models.NewUser{
		SenderName:    "a",
		RecipientName: "",
		Context:       "",
		Time:          "",
		Type:          "",
		Result:        make(chan map[string]string, 1024),
		CBack:         context.Background(),
		Mutex:         &sync.Mutex{},
		Sum:           nil,
		Wg:            &sync.WaitGroup{},
	}
	user.RedisSMembers()
	for {
		select {
		case lr := <-user.Result:
			for k, v := range lr {
				log.Println(k)
				log.Println(v)
			}

		}
	}
}

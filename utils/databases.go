package utils

import (
	"context"
	"github.com/go-redis/redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"miniChat/conf"
)

var (
	MDb *gorm.DB
	RDb *redis.Client
	err error
)

func init() {
	mysqlLink()
	redisLink()
}
func mysqlLink() {
	MDb, err = gorm.Open(mysql.Open(conf.User+":"+conf.Password+"@"+"tcp("+conf.LocalHost+")/"+conf.Database+"?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("Mysql数据库启动失败！")
		return
	}
	log.Println("Mysql数据库连接成功！")
}
func redisLink() {

	RDb = redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		Password:     "", // no password set
		DB:           1,  // use default DB
		PoolFIFO:     true,
		PoolSize:     100000,
		MaxIdleConns: 500,
	})
	if RDb.Ping(context.Background()).Err() != nil {
		panic("Redis连接失败！")
		return
	}
	log.Println("Redis连接成功！")
}

package main

import (
	"miniChat/routers"
	_ "miniChat/utils"
	"runtime"
)

func main() {
	//go controllers.WebsocketMain()
	runtime.GOMAXPROCS(runtime.NumCPU()) //设置需要用到的cpu数量
	routers.Router().Run(":8080")
}

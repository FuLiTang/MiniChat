package routers

import (
	"github.com/gin-gonic/gin"
	"miniChat/WebControllers"
	"miniChat/controllers"
)

func Router() *gin.Engine {
	r := gin.Default()
	uC := r.Group("/user")
	//userCURD
	uC.POST("/create", controllers.CreateUser)
	uC.POST("/delete", controllers.DeleteUser)
	uC.POST("/update", controllers.UpdateUser)
	uC.POST("/query", controllers.RegisterNameAndPassword)
	u := r.Group("/userData")
	{
		u.POST("/headPortrait", controllers.UserDataImage)
		u.POST("/addFriends", controllers.AddFriends)
		u.POST("/offLogin", controllers.Off)
	}
	chat := r.Group("/chat")
	{
		chat.POST("/home", controllers.LinkMan)
		chat.POST("/homepage", controllers.LinkManPage)
	}
	web := r.Group("/web")
	{
		web.GET("/index", WebControllers.Index)
		web.GET("/home", WebControllers.Home)
		web.GET("/miniChat", WebControllers.MiniChat)

	}
	socket := r.Group("/socket")
	{
		//r.GET("/socket", controllers.WebSocket)
		socket.GET("/socket", controllers.Socket)
		socket.GET("/socketall", controllers.SocketAll)
		socket.GET("/miniChatCount")
	}
	r.GET("/text", WebControllers.TextHome)
	r.LoadHTMLGlob("views/*")     //目录加载
	r.Static("/static", "static") //静态文件夹
	return r
}

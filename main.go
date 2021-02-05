package main

import (
	"os"
	"singo/conf"
	"singo/server"
	"singo/socket"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	r := server.NewRouter()

	r.Static("/static","./static")  //fang 自定义添加 加载静态文件
	r.LoadHTMLGlob(os.Getenv("VIEWS_PATTERN"))  //fang 自定义添加 加载模板

	//go api.Cron() //开启定时器功能

	//开启tcp server
	//go socket.TcpServer()
	//time.Sleep(time.Second)
	//开启tcp client
	//go socket.TcpClient()

	//开启udp server
	go socket.UdpServer()
	//开启udp client
	go socket.UdpClient()

	r.Run(":3000")


}



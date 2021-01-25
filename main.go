package main

import (
	"os"
	"singo/conf"
	"singo/server"
)

func main() {
	// 从配置文件读取配置
	conf.Init()

	// 装载路由
	r := server.NewRouter()

	r.Static("/static","./static")  //fang 自定义添加 加载静态文件
	r.LoadHTMLGlob(os.Getenv("VIEWS_PATTERN"))  //fang 自定义添加 加载模板

	r.Run(":3000")
}

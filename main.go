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

/*func corn(){
	//定时器测试 如果是restfulapi开发 可屏蔽该功能
	for range time.Tick(time.Millisecond*1000*1){
		//fmt.Println("Hello 定时器: "+time.Now().Format("2006-01-02 15:04:05"))
		api.CronTest1()
	}

}
*/
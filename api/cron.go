package api

import (
	"fmt"
	"time"
)

// cron 定时器测试
func CronTest() {
	fmt.Println("进入CronTest方法1 "+time.Now().Format("2006-01-02 15:04:05"))
}

func Cron(){
	fmt.Println("进入CronCron方法")
	//定时器测试 如果是restfulapi开发 可屏蔽该功能
	for range time.Tick(time.Millisecond*1000*60){
		fmt.Println("Hello 定时器: "+time.Now().Format("2006-01-02 15:04:05"))
	}

}



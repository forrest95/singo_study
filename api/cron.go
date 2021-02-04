package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"singo/serializer"
	"time"
)

// cron 定时器测试
func CronTest(c *gin.Context) {
	fmt.Println("进入CronTest方法")

	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "我是CronTest测试方法 返回！",
	})

}

// cron 定时器测试
func CronTest1() {
	fmt.Println("进入CronTest方法1 "+time.Now().Format("2006-01-02 15:04:05"))
}



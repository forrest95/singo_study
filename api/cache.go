package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"singo/cache"
	"singo/serializer"
)

// 缓存 redis 测试接口
func CacheTest(c *gin.Context) {
	fmt.Println("进入CacheTest方法")
	//ctx := context.Background()

	//data, err :=cache.RedisClient.Get("bb").Result() //获取string
	//data, err :=cache.RedisClient.LRange("aa",0,100).Result() //获取list
	//data, err :=cache.RedisClient.HGetAll("logs_can").Result() //获取hash
	data, err :=cache.RedisClient.HGet("logs_can","11K0130HC204183").Result() //获取hash
	if err != nil {
		fmt.Printf("get score failed, err:%v\n", err)
		return
	}
	c.JSON(200, serializer.Response{
		Code: 220,
		Msg:  "我是CacheTest测试方法 返回！",
		Data:  data,
	})

}



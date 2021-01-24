package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
	"singo/serializer"
)

// FangTest 测试接口
func TestByte(c *gin.Context) {
	fmt.Println("进入TestByte方法")

	var MySecret = []byte("夏天夏天悄悄过去")
	var MyString = []string{"夏天夏天","秋香秋香"}
	fmt.Println("打印MySecret")
	fmt.Println(MySecret)
	fmt.Println(reflect.TypeOf(MySecret))
	fmt.Println(MySecret[2])
	fmt.Printf("%c,%v",MySecret[2],MySecret[2])


	fmt.Println()
	fmt.Println("----------------分割线------------")

	fmt.Println(MyString)
	fmt.Println(reflect.TypeOf(MyString))

	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "我是TestByte测试方法 返回！",
		Data:MySecret,
	})


}



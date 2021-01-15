package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"singo/serializer"
)

// FangTest 测试接口
func FangTest(c *gin.Context) {
	fmt.Println("进入FangTest方法")

	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "我是FangTest测试方法 返回！",
	})

}



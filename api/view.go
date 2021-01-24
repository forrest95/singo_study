package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FangTest 测试接口
func ViewIndex(c *gin.Context) {
	fmt.Println("进入ViewIndex方法")

	c.HTML(http.StatusOK, "index/index.html", gin.H{
		"msg": "easy gin",
	})

}



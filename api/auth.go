package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"singo/model"
	"singo/serializer"
	"singo/util"
)

// FangTest 测试接口
//注册获取token
func AuthHandler(c *gin.Context) {
	fmt.Println("进入AuthHandler方法")

	// 用户发送用户名和密码过来
	var user model.User
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "无效的参数",
		})
		return
	}
	fmt.Println("打印user ")
	fmt.Println(user)
	fmt.Println(user.UserName)
	// 校验用户名和密码是否正确
	//if user.UserName == "q1mi" && user.PasswordDigest == "q1mi123" {
	if user.UserName == "q1mi" {
		fmt.Println("生成token")
		// 生成Token
		tokenString, _ := util.GenToken(user.UserName)
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": gin.H{"token": tokenString},
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 2002,
		"msg":  "鉴权失败",
	})
	return

	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "我是AuthHandler测试方法 返回！",
	})

}

//测试token
func AuthhomeHandler(c *gin.Context) {
	fmt.Println("进入AuthhomeHandler方法")
	username := c.MustGet("username").(string)
	fmt.Println("打印username: "+username)
	c.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg":  "success",
		"data": gin.H{"username": username},
	})
}
package api

import (
	"fmt"
	"singo/model"
	"singo/serializer"
	"singo/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口
func UserRegister(c *gin.Context) {
	fmt.Println("进入用户注册方法11")

	var service service.UserRegisterService
	fmt.Println("ShoudBind之前")
	if err := c.ShouldBind(&service); err == nil {
		fmt.Println("ShoudBind之后")

		res := service.Register()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserLogin 用户登录接口
func UserLogin(c *gin.Context) {
	fmt.Println("进入用户登录方法")
	var service service.UserLoginService
	if err := c.ShouldBind(&service); err == nil {
		res := service.Login(c)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// UserMe 用户详情
func UserMe(c *gin.Context) {
	user := CurrentUser(c)
	res := serializer.BuildUserResponse(*user)
	c.JSON(200, res)
}

// UserLogout 用户登出
func UserLogout(c *gin.Context) {
	s := sessions.Default(c)
	s.Clear()
	s.Save()
	c.JSON(200, serializer.Response{
		Code: 0,
		Msg:  "登出成功",
	})
}

//校验用户密码
func UserCheckPwd(c *gin.Context){
	fmt.Println("进入UserCheckPwd方法")

	//查询user_name=hwtc的数据
	data:=model.DB.Model(model.User{}).Where("user_name = ?", "hwtc").Find(&model.User{})
	res:=data.Value //value是查出来的数据

	pwd:=res.(*model.User).PasswordDigest
	fmt.Println("打印密码密文: ")
	fmt.Println(pwd)
	check_res:=res.(*model.User).CheckPassword("hwtc@666")
	fmt.Println("查看密码校验结果")
	fmt.Println(check_res)
	c.JSON(200, serializer.Response{
		Code: 444,
		Msg:  "进入UserCheckPwd方法",
		Data:data,
	})
}

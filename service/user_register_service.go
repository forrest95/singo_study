package service

import (
	"fmt"
	"singo/model"
	"singo/serializer"
)

// UserRegisterService 管理用户注册服务
type UserRegisterService struct {
	Nickname        string `form:"nickname" json:"nickname" binding:"required,min=2,max=30"`
	UserName        string `form:"user_name" json:"user_name" binding:"required,min=4,max=30"`
	Password        string `form:"password" json:"password" binding:"required,min=8,max=40"`
	PasswordConfirm string `form:"password_confirm" json:"password_confirm" binding:"required,min=8,max=40"`
}

// valid 验证表单
func (service *UserRegisterService) valid() *serializer.Response {
	fmt.Println("进入user_register_valid方法")
	if service.PasswordConfirm != service.Password {
		return &serializer.Response{
			Code: 40001,
			Msg:  "两次输入的密码不相同",
		}
	}

	count := 0
	model.DB.Model(&model.User{}).Where("nickname = ?", service.Nickname).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40001,
			Msg:  "昵称被占用",
		}
	}

	count = 0
	model.DB.Model(&model.User{}).Where("user_name = ?", service.UserName).Count(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40001,
			Msg:  "用户名已经注册",
		}
	}

	return nil
}

// Register 用户注册
func (service *UserRegisterService) Register() serializer.Response {
	fmt.Println("进入user_register_service")
	user := model.User{
		Nickname: service.Nickname,
		UserName: service.UserName,
		Status:   model.Active,

	}
	fmt.Println("打印用户")
	fmt.Println(user)
	fmt.Println("进入表单验证阶段")
	// 表单验证
	if err := service.valid(); err != nil {
		return *err
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Err(
			serializer.CodeEncryptError,
			"密码加密失败",
			err,
		)
	}

	// 创建用户
	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.ParamErr("注册失败", err)
	}

	return serializer.BuildUserResponse(user)
}

package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"singo/model"
	"singo/serializer"
)

// FangTest 测试接口
func GormAdd(c *gin.Context) {
	fmt.Println("进入GormAdd方法")

	article :=model.Article{
		Title:   "aa",
		Content: "aa1",
		Author:  "aa2",
	}

	data :=model.DB.Create(&article)

	c.JSON(200, serializer.Response{
		Code: 111,
		Msg:  "我是GormAdd测试方法 返回！",
		Data:data,
	})

}

func GormUpdate(c *gin.Context){
	fmt.Println("进入GormUpdate方法")

	article :=model.Article{
		Model:gorm.Model{ID:4},
		Content:"content11",
		Author:"fang11",
	}

	article.Title="title11"
	//data:=model.DB.Save(&article); //更新所有字段

	data:=model.DB.Model(&article).Update("title","titie_我是更新后的内容11") //更新单个属性
	c.JSON(200, serializer.Response{
		Code: 222,
		Msg:  "我是GormUpdate测试方法 返回！",
		Data:data,
	})
}



package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"singo/model"
	"singo/serializer"
	"strconv"
)

// gorm add 测试接口
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

//gorm update
func GormUpdate(c *gin.Context){
	fmt.Println("进入GormUpdate方法")

	article :=model.Article{
		Model:gorm.Model{ID:4},
		Content:"content11",
		Author:"fang11",
	}

	article.Title="title11"
	//data:=model.DB.Save(&article); //更新所有字段

	//data:=model.DB.Model(&article).Update("title","titie_我是更新后的内容11") //更新单个属性

	//更新选定字段
	data:=model.DB.Model(&article).Select("title","content").Updates(
		map[string]interface{}{
			"title": "hello",
			"content": "content",
			"active": false,
		})

	c.JSON(200, serializer.Response{
		Code: 222,
		Msg:  "我是GormUpdate测试方法 返回！",
		Data:data,
	})
}

//gorm select demo
func GormSelect(c *gin.Context){
	fmt.Println("进入GormSelect方法")
	//article:=model.Article{}
	//data:=model.DB.Debug().Find(&article,13) //查询id 为13的

	//data:=model.DB.Debug().Where("id = ?", 11).First(&article)  //查询id=11的

	//查询所有
	//查询所有记录
	var articles []model.Article  //定义一个切片
	//articles := make([]model.Article, 3)
	//data:=model.DB.Find(&articles)
	//fmt.Printf("articles:%v len(articles):%v cap(articles):%v\n", articles, len(articles), cap(articles))

	// LIKE
	data:=model.DB.Debug().Where("title LIKE ?", "%aa%").Find(&articles)

	c.JSON(200, serializer.Response{
		Code: 333,
		Msg:  "我是GormSelect测试方法 返回！",
		Data:data,
	})
}

//gorm select 分页查询
func GormPage(c *gin.Context){
	fmt.Println("进入GormPage方法")
	//article:=model.Article{}
	//data:=model.DB.Debug().Find(&article,13) //查询id 为13的
	//data:=model.DB.Debug().Where("id = ?", 11).First(&article)  //查询id=11的

	//查询所有记录
	var articles []model.Article  //定义一个切片
	//articles := make([]model.Article, 3)
	//data:=model.DB.Find(&articles)
	//fmt.Printf("articles:%v len(articles):%v cap(articles):%v\n", articles, len(articles), cap(articles))

	// LIKE
	data:=model.DB.Where("title LIKE ?", "%aa%").Find(&articles)
aa,bb,_:=Pagination(c)
fmt.Println("***********************")
fmt.Println(aa)
fmt.Println(bb)
	fmt.Println("***********************")

	c.JSON(200, serializer.Response{
		Code: 444,
		Msg:  "我是GormPage测试方法 返回！",
		Data:data,
	})
}

// Pagination is page util
func Pagination(ctx *gin.Context) (pageStr string, num int, err error) {
	fmt.Println("进入Pagination方法")
	limit := ctx.DefaultQuery("page_size", "3")
	pageNumber := ctx.DefaultQuery("page_number", "2")
	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt < 0 {
		return "", 0, err
		}
	pageNumberInt, err := strconv.Atoi(pageNumber)
	if err != nil || pageNumberInt < 0 {
		return "", 0, err
		}
	if pageNumberInt != 0 {
		pageNumberInt--
		}
	offsetInt := limitInt * pageNumberInt
	pageStr = fmt.Sprintf(" limit %d offset %d", limitInt+1, offsetInt)
	return pageStr, limitInt, nil
	}



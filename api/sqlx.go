package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ifconfigure/gorm-paginate/paginage"
	"singo/model"
	"singo/serializer"
)

// sqlx add 测试接口
func SqlxAdd(c *gin.Context) {
	fmt.Println("进入SqlxAdd方法")


	c.JSON(200, serializer.Response{
		Code: 111,
		Msg:  "我是SqlxAdd测试方法 返回！",
		//Data:data,
	})

}

//sqlx update
func SqlxUpdate(c *gin.Context){
	fmt.Println("进入SqlxUpdate方法")


	c.JSON(200, serializer.Response{
		Code: 222,
		Msg:  "我是SqlxUpdate测试方法 返回！",
		//Data:data,
	})
}

//sqlx select demo
func SqlxSelect(c *gin.Context){
	fmt.Println("进入SqlxSelect方法")

	sqlStr := "select id, title,content, author from article where id=?"
	var article model.Article
	err := model.Sqlx_db.Get(&article, sqlStr, 1)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}

	c.JSON(200, serializer.Response{
		Code: 333,
		Msg:  "我是SqlxSelect测试方法 返回！",
		Data:article,
	})
}

//分页参考github https://github.com/ifconfigure/gorm-paginate

//gorm select 分页查询
func SqlxPage(c *gin.Context){
	fmt.Println("进入GormPage方法")
	//article:=model.Article{}
	//data:=model.DB.Debug().Find(&article,13) //查询id 为13的
	//data:=model.DB.Debug().Where("id = ?", 11).First(&article)  //查询id=11的

	//查询所有记录
	var articles []model.Article  //定义一个切片
	//articles := make([]model.Article, 3)
	//data:=model.DB.Find(&articles)
	//fmt.Printf("articles:%v len(articles):%v cap(articles):%v\n", articles, len(articles), cap(articles))

	// LIKE 查询数据
	//data:=model.DB.Where("title LIKE ?", "%aa%").Find(&articles)

	//start 分页测试
	currentPage:=2
	// 1、Chaining - 链式操作查询
	data := model.DB.Model(model.Article{}).
		Order("id desc").
		Where("title like ?", "%aa%").
		//Preload("User.Country")
		Find(&articles)

	//2、use paginate - 调用分页类
	res, err := paginage.Paginate(data, int(currentPage), &articles)

	if err != nil {
		fmt.Println(err)
		return
	}

	//3、output - 输出
	c.JSON(200, serializer.Response{
		Code: 444,
		Msg:  "我是GormPage 测试方法 返回！",
		Data:res,
	})
	//end 分页测试
}





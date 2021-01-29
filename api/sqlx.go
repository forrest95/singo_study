package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ifconfigure/gorm-paginate/paginage"
	"math"
	"reflect"
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
	fmt.Println("进入SqlxPage方法")

	var articles []model.Article  //定义一个切片 存储查询结果

	//查询所有记录
	sqlStr := "select id, title,content, author from article where id>?"
	err := model.Sqlx_db.Select(&articles, sqlStr, 1)

	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	//fmt.Printf("articles:%#v\n", articles)

	//start 分页测试
	currentPage:=2
	// 1、Chaining - 链式操作查询
	data := model.DB.Model(model.Article{}).
		Order("id desc").
		Where("title like ?", "%aa%").
		//Preload("User.Country")
		Find(&articles)
	fmt.Println(111111)
	fmt.Println(data)
	fmt.Println(111111)
fmt.Println(reflect.TypeOf(data))
	fmt.Println(111111)
	//2、use paginate - 调用分页类
	res, err := paginage.Paginate(data, int(currentPage), &articles)

	if err != nil {
		fmt.Println(err)
		return
	}

	//3、output - 输出
	c.JSON(200, serializer.Response{
		Code: 444,
		Msg:  "我是SqlxPage 测试方法 返回！",
		Data:res,
		//Data:articles,
	})
	//end 分页测试
}

//分页方法，根据传递过来的页数，每页数，总数，返回分页的内容 7个页数 前 1，2，3，4，5 后 的格式返回,小于5页返回具体页数
func Paginator1(page, prepage int, nums int64) map[string]interface{} {

	var firstpage int //前一页地址
	var lastpage int  //后一页地址
	//根据nums总数，和prepage每页数量 生成分页总数
	totalpages := int(math.Ceil(float64(nums) / float64(prepage))) //page总数
	if page > totalpages {
		page = totalpages
	}
	if page <= 0 {
		page = 1
	}
	var pages []int
	switch {
	case page >= totalpages-5 && totalpages > 5: //最后5页
		start := totalpages - 5 + 1
		firstpage = page - 1
		lastpage = int(math.Min(float64(totalpages), float64(page+1)))
		pages = make([]int, 5)
		for i, _ := range pages {
			pages[i] = start + i
		}
	case page >= 3 && totalpages > 5:
		start := page - 3 + 1
		pages = make([]int, 5)
		firstpage = page - 3
		for i, _ := range pages {
			pages[i] = start + i
		}
		firstpage = page - 1
		lastpage = page + 1
	default:
		pages = make([]int, int(math.Min(5, float64(totalpages))))
		for i, _ := range pages {
			pages[i] = i + 1
		}
		firstpage = int(math.Max(float64(1), float64(page-1)))
		lastpage = page + 1
		//fmt.Println(pages)
	}
	paginatorMap := make(map[string]interface{})
	paginatorMap["pages"] = pages
	paginatorMap["totalpages"] = totalpages
	paginatorMap["firstpage"] = firstpage
	paginatorMap["lastpage"] = lastpage
	paginatorMap["currpage"] = page
	return paginatorMap
}





package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"reflect"
	"singo/util"
	"strconv"
)

// FangTest 测试接口
func ViewIndex(c *gin.Context) {
	fmt.Println("进入ViewIndex方法")

	c.HTML(200,"index/index",gin.H{
		"msg": "easy singo",
	})

	/*c.JSON(200,"index/index",gin.H{
		"msg": "easy singo",
	})*/

}

//分页测试
func ViewPage(c *gin.Context){
	fmt.Println("进入ViewPage方法")

	//接收分页参数
	currentPage, _ :=strconv.Atoi(c.DefaultQuery("page","1"))
	pageSize, _ :=strconv.Atoi(c.DefaultQuery("pageSize","5"))

	fmt.Println(currentPage,pageSize)

	//分页一 参考文档 https://my.oschina.net/tongjh/blog/1800443
	var paging=util.CreatePaging(int64(currentPage), int64(pageSize),365);
	fmt.Println(333333333)
	fmt.Println(reflect.TypeOf(paging))
	fmt.Println(paging)
	fmt.Println(333333333)

	// 分页二  参考文档 https://blog.csdn.net/ciwei_ice/article/details/50429835
	data2:=Paginator(currentPage,pageSize,365)
	fmt.Println(11111)
	fmt.Println(reflect.TypeOf(data2))
	fmt.Println(data2)
	fmt.Println(11111)

	pageStr, a, b :=util.GetPageStr(c, c.DefaultQuery("pageSize","5"), c.DefaultQuery("page","1"))
	fmt.Println("GetPageStr",pageStr,a,b)

	c.HTML(http.StatusOK, "view/page", gin.H{
		"msg": "ViewPage 视图层分页",
		"paging":paging,
		"pageStr":pageStr,
		"paginator":data2,
	})
}

//分页方法，根据传递过来的页数，每页数，总数，返回分页的内容 7个页数 前 1，2，3，4，5 后 的格式返回,小于5页返回具体页数
func Paginator(page, prepage int, nums int64) map[string]interface{} {

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


 



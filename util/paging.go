package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"strconv"
)

func CreatePaging(page, pagesize, total int64) *Paging {
	if page < 1 {
		page = 1
	}
	if pagesize < 1 {
		pagesize = 10
	}

	page_count := math.Ceil(float64(total) / float64(pagesize))

	paging := new(Paging)
	paging.Page = page
	paging.Pagesize = pagesize
	paging.Total = total
	paging.PageCount = int64(page_count)
	paging.NumsCount = 7
	paging.setNums()
	return paging
}

type Paging struct {
	Page      int64   //当前页
	Pagesize  int64   //每页条数
	Total     int64   //总条数
	PageCount int64   //总页数
	Nums      []int64 //分页序数
	NumsCount int64   //总页序数
}

func (this *Paging) setNums() {
	this.Nums = []int64{}
	if this.PageCount == 0 {
		return
	}

	half := math.Floor(float64(this.NumsCount) / float64(2))
	begin := this.Page - int64(half)
	if begin < 1 {
		begin = 1
	}

	end := begin + this.NumsCount - 1
	if end >= this.PageCount {
		begin = this.PageCount - this.NumsCount + 1
		if begin < 1 {
			begin = 1
		}
		end = this.PageCount
	}

	for i := begin; i <= end; i++ {
		this.Nums = append(this.Nums, i)
	}
}

//获取分页sql limit %d offset %d
// Pagination is page util 参考文档 https://www.cnblogs.com/lyp0626/p/12056143.html
func GetPageStr(ctx *gin.Context,limit,pageNumber string) (pageStr string, num int, err error) {
	//limit := ctx.DefaultQuery("page_size", "5")  //修改page_size 和page_num参数获取方式
	//pageNumber := ctx.DefaultQuery("page_number", "5")
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
	pageStr = fmt.Sprintf(" limit %d offset %d", limitInt, offsetInt)
	return pageStr, limitInt, nil
}
package model

import (
	"github.com/jinzhu/gorm"
)

// Article 用户模型
type Article struct {
	gorm.Model
	Title       string
	Content string
	Author       string
}

// 将 Article 的表名设置为 `article`
func (Article) TableName() string {
	return "article"
}

// GetArticle 用ID获取article
func GetArticle(ID interface{}) (Article, error) {
	var article Article
	result := DB.First(&article, ID)
	return article, result.Error
}





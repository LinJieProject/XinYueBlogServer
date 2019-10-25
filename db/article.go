package db

import (
	"fmt"
	"xybs/models"
)

// QueryAllArticle 查询所有的文章基本信息
func QueryAllArticle() (articleList []*models.ArticleInfo, err error) {
	sqlStr := `SELECT
	id,
	title,
	view_count,
	comment_count,
	username,
	summary
FROM
	article;`
	err = db.Select(&articleList, sqlStr)
	if err != nil {
		fmt.Println("查询所有文章信息失败！")
		return
	}
	return
}

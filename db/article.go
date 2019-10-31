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

// QueryAnArticleByID 查询一篇文章
func QueryAnArticleByID(id int64) (articleDetail models.ArticleDetail,err error) {
	sqlSrt1:="SELECT id,title,view_count,comment_count,username,summary FROM article WHERE id = ?"
	sqlSrt2:="SELECT content FROM article WHERE id = ?"
	var articleInfo models.ArticleInfo
	var content string
	err=db.Get(&articleInfo,sqlSrt1,id)
	if err != nil {
		fmt.Println("查询文章失败！")
		return
	}
	err=db.Get(&content,sqlSrt2,id)
	if err != nil {
		fmt.Println("查询文章失败！")
		return
	}
	articleDetail.ArticleInfo=articleInfo
	articleDetail.Content=content
	return
}
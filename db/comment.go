package db

import (
	"fmt"
	"xybs/models"
)

// QueryCommentList 查询评论
func QueryCommentList(articleID int64) (commentList []models.Comment, err error) {
	sqlStr := "SELECT id,content,username FROM comment WHERE article_id = ?"
	err = db.Select(&commentList, sqlStr, articleID)
	if err != nil {
		fmt.Println("查询评论失败！")
		return
	}
	return
}

// InsertComment 插入一条评论
func InsertComment(content, username string, articleID int64) {
	sqlStr := `INSERT INTO COMMENT ( content, username, article_id )
VALUES
	( ?,?,? )`
	_, err := db.Exec(sqlStr, content, username, articleID)
	if err != nil {
		fmt.Println("插入评论失败！")
		return
	}
	fmt.Println("插入评论成功！")
}

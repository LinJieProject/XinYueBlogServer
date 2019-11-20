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

// UpdateCommentCount 更新评论数
func UpdateCommentCount(id int64) {
	sqlStr := "UPDATE article SET comment_count=comment_count+1 WHERE id=?"
	_, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("更新评论数失败, err:%v\n", err)
		return
	}
	fmt.Println("更新评论数成功。")
}

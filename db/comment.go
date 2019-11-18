package db

import "fmt"

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

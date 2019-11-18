package models

// Comment 评论结构体
type Comment struct {
	ID        int64  `db:"id" json:"id"`
	Content   string `db:"content" json:"content"`
	Username  string `db:"username" json:"username"`
	ArticleID int64  `db:"article_id" json:"article_id"`
}

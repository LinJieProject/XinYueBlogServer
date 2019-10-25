package models

// ArticleInfo 存放基本文章数据
type ArticleInfo struct {
	Id           int64     `db:"id" json:"id"`
	Title        string    `db:"title" json:"title"`
	ViewCount    uint32    `db:"view_count" json:"view_count"`
	CommentCount uint32    `db:"comment_count" json:"comment_count"`
	Username     string    `db:"username" json:"username"`
	Summary      string    `db:"summary" json:"summary"`
}

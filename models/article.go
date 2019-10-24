package models

import "time"
// ArticleInfo 存放基本文章数据
type ArticleInfo struct {
	Id           int64     `db:id`
	TagID        int64     `db:tag_id`
	Title        string    `db:title`
	ViewCount    uint32    `db:view_count`
	CommentCount uint32    `db:comment_count`
	Username     string    `db:username`
	Summary      string    `db:summary`
	CreateTime   time.Time `db:create_time`
}

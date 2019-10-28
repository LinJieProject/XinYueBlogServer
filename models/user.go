package models

// User 用户结构体
type User struct {
	ID       int64  `db:"id" json:"id"`
	UserID   int64  `db:"user_id" json:"user_id"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
	Gender   string `db:"gender" json:"gender"`
}

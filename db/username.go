package db

// Username 用户结构体
type Username struct{
	ID int64 `db:"id" json:"id"`
	UserID int64 `id:"user_id" json:"user_id"`
	Username string `id:"username" json:"username"`
	password string `id:"password" json:"password"`
	gender uint8 `id:"gender" json:"gender"`
}

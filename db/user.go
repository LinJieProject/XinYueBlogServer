package db

import (
	"fmt"
	"xybs/models"
)

// QueryUser
func QueryUser(username, password string) (user models.User, err error) {
	sqlStr := "SELECT * FROM `user` WHERE username=? AND `password`=?"
	err = db.Get(&user, sqlStr, username, password)
	if err != nil {
		fmt.Println("找不到该用户")
		return
	}
	return
}

func InsertUser(userID int64,username,password string){
	sqlStr:="INSERT INTO `user`(`user_id`,`username`,`password`) VALUES (?,?,?)"
	_, err := db.Exec(sqlStr,userID, username, password)
	if err != nil {
		fmt.Println("插入用户失败！")
		return
	}
	fmt.Println("插入新用户数据成功！")
}

func QueryUserByUsername(username string) (isExist bool) {
	sqlStr:="SELECT username FROM `user` WHERE username=?"
	var user models.User
	err := db.Get(&user, sqlStr, username)
	if err != nil {
		fmt.Println("不存在该用户。")
		return false
	}
	return true
}
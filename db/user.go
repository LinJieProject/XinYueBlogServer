package db

import (
	"fmt"
	"xybs/models"
)

// QueryUser
func QueryUser(username, password string) (user models.User, err error) {
	sqlSrt := "SELECT * FROM `user` WHERE username=? AND `password`=?"
	//var user models.User
	err = db.Get(&user, sqlSrt, username, password)
	if err != nil {
		fmt.Println("找不到该用户")
		return
	}
	return
}

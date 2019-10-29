package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xybs/db"
	"xybs/models"
)

func ArticleListHandler(c *gin.Context) {
	articleList, err := db.QueryAllArticle()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": articleList,
	})
}

func LoginHandler(c *gin.Context) {
	var tempUser models.User
	err:=c.ShouldBind(&tempUser)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "用户名或密码不能为空！",
		})
		return
	}
	user, err := db.QueryUser(tempUser.Username, tempUser.Password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "用户名或密码错误！",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "登录成功！",
		"user": user,
	})
}

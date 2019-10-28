package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xybs/db"
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
	//username := c.PostForm("username")
	//password := c.PostForm("password")
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	user, err := db.QueryUser(username, password)
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

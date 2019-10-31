package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
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
	err := c.ShouldBind(&tempUser)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "用户名或密码不能为空！",
		})
		return
	}
	user, err := db.QueryUser(tempUser.Username, tempUser.Password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "用户名或密码错误！",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "登录成功！",
		"user": user,
	})
}

func RegisterHandler(c *gin.Context) {
	var tempUser models.User
	err := c.ShouldBind(&tempUser)
	fmt.Println("username:", tempUser.Username, "password:", tempUser.Password, "password2:", tempUser.Password2)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "用户名或密码不能为空！",
		})
		return
	}
	// 去除空格
	username := strings.Replace(tempUser.Username, " ", "", -1)
	password := strings.Replace(tempUser.Password, " ", "", -1)
	password2 := strings.Replace(tempUser.Password2, " ", "", -1)
	fmt.Printf("%#v\n", username)
	fmt.Printf("%#v\n", password)
	fmt.Printf("%#v\n", password2)
	if username == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"msg": "用户名或密码不能为空！",
		})
		return
	}
	if password != password2 {
		c.JSON(http.StatusOK, gin.H{
			"msg": "两次密码不一致！",
		})
		return
	}
	if db.QueryUserByUsername(username) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "该用户名已被注册！",
		})
		return
	}
	now := time.Now()        //获取当前时间
	userID := now.UnixNano() //获取纳秒时间戳作为用户ID
	db.InsertUser(userID, username, password)
	user, err := db.QueryUser(username, password)
	c.JSON(http.StatusOK, gin.H{
		"msg":  "注册成功！",
		"user": user,
	})
}

func ArticleDetailHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 0, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	articleDetail, err := db.QueryAnArticleByID(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": articleDetail,
	})
}

package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
	"xybs/db"
	"xybs/models"
)

type Result struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func returnMsg(ctx *gin.Context, code int, data interface{}, msg string) {
	ctx.JSON(code, Result{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func ArticleListHandler(c *gin.Context) {
	articleList, err := db.QueryAllArticle()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err,
		})
		return
	}
	// 根据文章的ID进行降序排序
	sort.Slice(articleList, func(i, j int) bool {
		return articleList[i].Id > articleList[j].Id
	})
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": articleList,
	})
}

// LoginHandler 用户登录处理函数
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

// RegisterHandler 用户注册处理函数
func RegisterHandler(c *gin.Context) {
	var tempUser models.User
	err := c.ShouldBind(&tempUser)
	fmt.Println("username:", tempUser.Username, "password:", tempUser.Password, "checkPass:", tempUser.CheckPass)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "用户名或密码不能为空！",
		})
		return
	}
	// 去除空格
	username := strings.Replace(tempUser.Username, " ", "", -1)
	password := strings.Replace(tempUser.Password, " ", "", -1)
	checkPass := strings.Replace(tempUser.CheckPass, " ", "", -1)
	fmt.Printf("%#v\n", username)
	fmt.Printf("%#v\n", password)
	fmt.Printf("%#v\n", checkPass)
	if username == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"msg": "用户名或密码不能为空！",
		})
		return
	}
	if password != checkPass {
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
	// 更新文章阅读量
	db.UpdateViewCount(id)
	articleDetail.ArticleInfo.ViewCount += 1
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": articleDetail,
	})
}

// @Summary 发布一篇文章
// @Description 用于用户发布文章
// @Tags 文章模块
// @Accept  json
// @Produce json
// @Param title formData string true "文章标题"
// @Param content formData string true "文章内容"
// @Param summary formData string true "文章摘要"
// @Success 200 {object} Result "{"code":200,"data":nil,"msg": "发布成功！"}"
// @Failure 400 {object} Result "{"code":400,"data":nil,"msg": "后端获取文章失败！"}"
// @Router /PublishArticle [post]
func PublishArticleHandler(c *gin.Context) {
	var article models.ArticleDetail
	err := c.ShouldBind(&article)
	fmt.Printf("%#v\n", article)
	if err != nil {
		returnMsg(c, 400, nil, "后端获取文章失败！")
		//c.JSON(http.StatusUnauthorized, gin.H{
		//	"msg": "后端获取文章失败！",
		//})
		return
	}
	db.InsertArticle(article.Content, article.Title, article.Username, article.Summary, article.ViewCount, article.CommentCount)
	//c.JSON(http.StatusOK, gin.H{
	//	"msg": "发布成功！",
	//})
	returnMsg(c, http.StatusOK, nil, "发布成功！")
}

// @Summary 发布评论
// @Description 用于用户发布评论
// @Tags 评论模块
// @Accept  json
// @Produce json
// @Param content formData string true "评论内容"
// @Param username formData string true "发布者"
// @Param article_id formData string true "关联文章ID"
// @Success 200 {object} Result "{"code":200,"data":nil,"msg": "发布评论成功！"}"
// @Failure 400 {object} Result "{"code":400,"data":nil,"msg": "后端获取评论失败！"}"
// @Router /PublishComment [post]
func PublishCommentHandler(c *gin.Context) {
	var comment models.Comment
	fmt.Printf("%#v\n", comment)
	err := c.ShouldBind(&comment)
	if err != nil {
		returnMsg(c, 400, nil, "后端获取评论失败！")
		return
	}
	db.InsertComment(comment.Content, comment.Username, comment.ArticleID)
	db.UpdateCommentCount(comment.ArticleID)
	returnMsg(c, 200, nil, "发布评论成功！")
}

// @Summary 获取文章评论
// @Description 用于获取某篇文章的评论
// @Tags 评论模块
// @Accept  json
// @Produce json
// @Param ArticleID formData string true "文章id"
// @Success 200 {object} Result "{"code":200,"data":nil,"msg": "查询评论成功！"}"
// @Failure 400 {object} Result "{"code":400,"data":nil,"msg": "查询评论失败！"}"
// @Router /Comment/:ArticleID [get]
func QueryCommentHandler(c *gin.Context) {
	articleIDStr := c.Param("ArticleID")
	articleID, err := strconv.ParseInt(articleIDStr, 0, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	commentList, err := db.QueryCommentList(articleID)
	if err != nil {
		returnMsg(c, 400, nil, "查询评论失败！")
		return
	}
	returnMsg(c, 200, commentList, "查询评论成功！")
}

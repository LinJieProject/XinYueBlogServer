package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xybs/db"
)

func ArticleListHandler(c *gin.Context)  {
	articleList,err:=db.QueryAllArticle()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  err,
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"code": 0,
		"data": articleList,
	})
}
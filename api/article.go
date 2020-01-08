package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jotyy/go-crud-example/model"
	"github.com/jotyy/go-crud-example/response"
)

func GetArticles(c *gin.Context) {
	article := model.Article{}

	articles, err := article.GetArticles(10)
	if err != nil {
		c.JSON(200, response.Response{
			Code: 40001,
			Msg:  "获取文章列表失败",
		})
		return
	}

	c.JSON(200, response.Response{
		Code: 0,
		Msg:  "获取文章列表成功",
		Data: articles,
	})
}

func AddArticles(c *gin.Context) {
	var article model.Article
	if err := c.ShouldBind(&article); err != nil {
		c.JSON(200, response.Response{
			Code: 40001,
			Msg:  "上传格式有误",
		})
		return
	}

	articleAdded, err := article.AddArticle()
	if err != nil {
		c.JSON(200, response.Response{
			Code: 40001,
			Msg:  "上传格式有误",
		})
		return
	}

	c.JSON(200, response.Response{
		Code: 0,
		Msg:  "上传文章成功",
		Data: articleAdded,
	})
}

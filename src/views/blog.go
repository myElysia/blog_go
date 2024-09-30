package views

import (
	"blogGo/src/model"
	"blogGo/src/query"
	"blogGo/src/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func BlogRouter(router *gin.RouterGroup) {
	route := router.Group("/blog")
	route.GET("/article/:uid", getPost)
	route.GET("/article/", getPosts)
	route.POST("/article/", createPost)
	route.PATCH("/article/:uid", updatePost)
	route.DELETE("/article/:uid", deletePost)
}

// @Summary 获取单个文章详情
// @Description 通过传递文章md5获取单个文章
// @Tags posts
// @Produce  json
// @Param  uid path string true "文章的MD5"  // 定义路径参数
// @Success 200 {object} model.Article
// @Router /post/{uid} [get]
func getPost(c *gin.Context) {
	a := query.Use(model.DB).Article
	uid := c.Param("uid")

	if article, err := a.WithContext(c.Request.Context()).Where(a.Md5.Eq(uid)).First(); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, utils.Response{Data: map[string]interface{}{"error": "Article not found"},
				Status: http.StatusNotFound}.Fail())
		} else {
			c.JSON(http.StatusInternalServerError, utils.Response{Data: err.Error(), Status: http.StatusInternalServerError}.Fail())
		}
	} else {
		c.JSON(http.StatusOK, utils.Response{Data: map[string]interface{}{"data": article}}.Success())
	}
}

// @Summary 获取文章列表
// @Description 获取所有文章的列表
// @Tags posts
// @Produce  json
// @Success 200 {array} []model.Article "成功返回文章列表"
// @Router /post [get]
func getPosts(c *gin.Context) {
	a := query.Use(model.DB).Article
	ctx := c.Request.Context()

	if articles, err := a.WithContext(ctx).Where(a.DeletedAt.IsNull()).Find(); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{Data: err.Error(), Status: http.StatusInternalServerError}.Fail())
	} else {
		c.JSON(http.StatusOK, utils.Response{Data: articles}.Success())
	}
}

// @Summary 创建新的Post
// @Description 通过请求体创建新的 Post
// @Tags posts
// @Accept  json
// @Produce  json
// @Param post body model.Article true "创建的文章json数据"
// @Success 200 {object} model.Article
// @Router /post [post]
func createPost(c *gin.Context) {
	var newArticle model.Article

	if err := c.ShouldBindJSON(&newArticle); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{Data: err.Error(), Status: http.StatusBadRequest}.Fail())
		return
	}
	a := query.Use(model.DB).Article
	if err := a.WithContext(c.Request.Context()).Create(&newArticle); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{Data: err.Error(), Status: http.StatusInternalServerError}.Fail())
	} else {
		c.JSON(http.StatusCreated, utils.Response{Data: newArticle}.Success())
	}
}

// @Summary 更新文章
// @Description 通过请求体更新文章内容等
// @Tags posts
// @Accept  json
// @Produce  json
// @Param post body model.Article true "更新数据"
// @Param uid path string true "文章md5"
// @Success 200 {object} model.Article
// @Router /post/{uid} [patch]
func updatePost(c *gin.Context) {
	var post map[string]interface{}
	uid := c.Param("uid")

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{Data: err.Error(), Status: http.StatusBadRequest}.Fail())
		return
	}
	a := query.Use(model.DB).Article

	if dbPost, err := a.WithContext(c.Request.Context()).Where(a.Md5.Eq(uid)).Updates(post); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{Data: err.Error(), Status: http.StatusInternalServerError}.Fail())
	} else {
		c.JSON(http.StatusOK, utils.Response{Data: dbPost, Status: http.StatusOK}.Success())
	}
}

// @Summary 删除文章
// @Description 通过路径中的文章md5对数据进行删除
// @Tags posts
// @Produce  json
// @Param uid path string true "文章md5"
// @Success 200 {object} []string
// @Router /post/{uid} [delete]
func deletePost(c *gin.Context) {
	uid := c.Param("uid")

	a := query.Use(model.DB).Article
	if _, err := a.WithContext(c.Request.Context()).Where(a.Md5.Eq(uid)).Delete(); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{Data: err.Error(), Status: http.StatusInternalServerError}.Fail())
	}
	c.JSON(http.StatusOK, utils.Response{Data: "success"}.Success())
}

package views

import (
	"blogGo/src/model"
	"blogGo/src/query"
	"blogGo/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func TodoRouter(router *gin.RouterGroup) {
	route := router.Group("/blog")
	route.GET("/todo/:uid", getPost)
	route.GET("/todo/", getPosts)
	route.POST("/todo/", createPost)
	route.PATCH("/todo/:uid", updatePost)
	route.DELETE("/todo/:uid", deletePost)
}

// @Summary 获取文章列表
// @Description 获取所有文章的列表
// @Tags posts
// @Produce  json
// @Success 200 {array} []model.TodoInfo "成功返回文章列表"
// @Router /post [get]
func getTodos(c *gin.Context) {
	tm := query.Use(model.DB).TodoInfo
	ctx := c.Request.Context()

	if todos, err := tm.WithContext(ctx).Where(tm.DeletedAt.IsNull()).Find(); err != nil {
		c.JSON(http.StatusInternalServerError,
			utils.Response{Data: err.Error(), Status: http.StatusInternalServerError}.Fail())
	} else {
		c.JSON(http.StatusOK, utils.Response{Data: todos}.Success())
	}
}

func createTodo(c *gin.Context) {
	var newTodo model.TodoInfo

	if err := c.ShouldBindJSON(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{Data: err.Error(), Status: http.StatusBadRequest}.Fail())
		return
	}
	tm := query.Use(model.DB).TodoInfo
	if err := tm.WithContext(c.Request.Context()).Create(&newTodo); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{Data: err.Error(), Status: http.StatusInternalServerError}.Fail())
		return
	}
	c.JSON(http.StatusOK, utils.Response{Data: newTodo}.Success())
}

func updateTodo(c *gin.Context) {
	var updTodo map[string]interface{}
	tm := query.Use(model.DB).TodoInfo
	uid := c.Param("id")
	id, err := strconv.ParseUint(uid, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{Data: err.Error(), Status: http.StatusBadRequest}.Fail())
		return
	}

	if err := c.ShouldBindJSON(&updTodo); err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{Data: err.Error(), Status: http.StatusBadRequest}.Fail())
		return
	}
	if dbTodo, err := tm.WithContext(c.Request.Context()).Where(tm.ID.Eq(uint(id))).Updates(updTodo); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{Data: err.Error(), Status: http.StatusInternalServerError}.Fail())
	} else {
		c.JSON(http.StatusOK, utils.Response{Data: dbTodo, Status: http.StatusOK}.Success())
	}
}

func deleteTodo(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	tm := query.Use(model.DB).TodoInfo
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Response{Data: err.Error(), Status: http.StatusBadRequest}.Fail())
		return
	}

	if dbTodo, err := tm.WithContext(c.Request.Context()).Where(tm.ID.Eq(uint(id))).Delete(); err != nil {
		c.JSON(http.StatusInternalServerError, utils.Response{Data: err.Error(), Status: http.StatusInternalServerError}.Fail())
	} else {
		c.JSON(http.StatusOK, utils.Response{Data: dbTodo, Status: http.StatusOK}.Success())
	}
}

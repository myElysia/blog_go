package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Data   any `json:"data"`
	Status int `json:"status"`
}

func (response Response) Success() gin.H {
	if response.Status <= 0 {
		response.Status = http.StatusOK
	}
	return gin.H{"status": response.Status, "data": response.Data}
}

func (response Response) Fail() gin.H {
	if response.Status <= 0 {
		response.Status = http.StatusBadRequest
	}
	return gin.H{"status": response.Status, "data": response.Data}
}

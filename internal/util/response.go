package util

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Errors  any    `json:"errors,omitempty"`
}

func OK(c *gin.Context, statusCode int, message string, data any) {
	c.JSON(statusCode, Response{Success: true, Message: message, Data: data})
}

func Fail(c *gin.Context, statusCode int, message string, errors any) {
	c.JSON(statusCode, Response{Success: false, Message: message, Errors: errors})
}

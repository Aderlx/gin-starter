package demoapp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello q1mi!",
	})
}

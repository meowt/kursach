package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var client *http.Client

type Response struct {
	Message string
}

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, Response{Message: "ok"})
}

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var router *gin.Engine

func server() {
	router = gin.Default()
	router.LoadHTMLGlob(cfg.Html + "*")
	router.Static("/assets", cfg.Assets)
	router.GET("/", index)
	e := router.Run(cfg.ServerHost + ":" + cfg.ServerPort)
	if e == nil {
		fmt.Println(e.Error())
		panic("Не удалось запустить сервер")
	}
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "header.html", nil)
	c.HTML(http.StatusOK, "indexBody.html", nil)
	c.HTML(http.StatusOK, "footer.html", nil)
}

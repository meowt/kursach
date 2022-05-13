package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var router *gin.Engine

func httpHandle() {
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":9090", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	router = gin.Default()
	router.LoadHTMLFiles(cfg.Html + "index.html")
	//router.Static("/assets", cfg.Assets)
	router.GET("/", index)
	router.Run(cfg.ServerHost + ":" + cfg.ServerPort)
}

func index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

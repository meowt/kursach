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
	router.POST("/posts/login", login)
	e := router.Run(cfg.ServerHost + ":" + cfg.ServerPort)
	if e == nil {
		fmt.Println(e.Error())
		panic("Не удалось запустить сервер")
	}
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "header.html", gin.H{
		"title": "Nymph",
	})
	c.HTML(http.StatusOK, "indexBody.html", nil)
	c.HTML(http.StatusOK, "footer.html", nil)
}

func login(c *gin.Context) {
	var CurrentUser user
	var e error
	e = c.Request.ParseForm()
	if e != nil {
		fmt.Println(e.Error())
		panic("Не удалось распарсить форму авторизации")
	}
	var LoginData struct {
		email, password string
	}
	LoginData.email = c.Request.FormValue("email")
	LoginData.password = c.Request.FormValue("password")
	if e != nil {
		fmt.Println(e.Error())
	}
	res, data := dbRequestLogin(LoginData)
	if res {
		c.HTML(http.StatusOK, "trueHeader.html", gin.H{
			"Username": data.Username,
		})
		c.HTML(http.StatusOK, "lol.html", gin.H{
			"CurrentUser": CurrentUser,
		})
		c.HTML(http.StatusOK, "footer.html", nil)
	} else {
		c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9090/")
	}
}

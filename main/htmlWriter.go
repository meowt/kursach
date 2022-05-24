package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var router *gin.Engine
var data user

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

func logged(c *gin.Context) {
	c.HTML(http.StatusOK, "trueHeader.html", gin.H{
		"username": data.Username,
	})
	c.HTML(http.StatusOK, "authIndexBody.html", nil)
	c.HTML(http.StatusOK, "footer.html", nil)
}

func login(c *gin.Context) {
	var e error
	//parsing POST form data
	e = c.Request.ParseForm()
	if e != nil {
		fmt.Println(e.Error())
		panic("Не удалось распарсить форму авторизации")
	}
	//creating struct from POST data
	var LoginData struct {
		email, password string
	}
	LoginData.email = c.Request.FormValue("email")
	LoginData.password = c.Request.FormValue("password")
	//checking
	var res bool
	res, data = dbRequestLogin(LoginData)
	if res {
		if isUrlFree(urls, data.Username) {
			addUserToSlice(data.Username)
			router.GET("/"+data.Username, logged)
		}
		c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9090/"+data.Username)
	} else {
		c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9090/")
	}
}

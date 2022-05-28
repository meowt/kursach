package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"net/http"
)

var router *gin.Engine
var data user
var store = sessions.NewCookieStore([]byte("random-hash-secret"))

func server() {
	router = gin.Default()
	router.LoadHTMLGlob(cfg.Html + "*")
	router.Static("/assets", cfg.Assets)
	router.GET("/", index)
	router.POST("/posts/login", login)
	router.POST("/posts/reg", reg)
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

		//r.GET("/incr", func(c *gin.Context) {
		//	session := sessions.Default(c)
		//	var count int
		//	v := session.Get("count")
		//	if v == nil {
		//		count = 0
		//	} else {
		//		count = v.(int)
		//		count++
		//	}
		//	session.Set("count", count)
		//	session.Save()
		//	c.JSON(200, gin.H{"count": count})
		//})

		c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9090/user/"+data.Username)
	} else {
		c.Writer.WriteString("error happened(")
	}
}

func reg(c *gin.Context) {
	var e error
	//parsing POST form data
	e = c.Request.ParseForm()
	if e != nil {
		fmt.Println(e.Error())
		panic("Не удалось распарсить форму регистрации")
	}
	//creating struct from POST data
	var RegData struct {
		email, password string
	}
	RegData.email = c.Request.FormValue("email")
	RegData.password = c.Request.FormValue("password")
	//uploading to db
	_, e = dbRequestReg(RegData)

	println()
	if e == nil {
		c.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:9090/user/"+data.Username)
	} else {
		c.Writer.WriteString("<script>" +
			"alert('Email is busy.')" +
			"</script>")
		c.Writer.WriteString("<script>" +
			"window.location.href = 'http://127.0.0.1:9090/'" +
			"</script>")
	}
}

var currentVar string

func getAllUsers() {
	var allUsers UsersIdUsername
	var e error
	allUsers, e = receiveAllUsersID()
	if e != nil {
		fmt.Println(e.Error())
		panic("Не удалось загрузить список пользователей")
	}
	for i := 0; i < len(allUsers.id); i++ {
		router.GET("/"+allUsers.id[i], userPage)
	}
}

func userPage(c *gin.Context) {
	c.HTML(http.StatusOK, "trueHeader.html", gin.H{
		"username": data.Username,
	})
	c.HTML(http.StatusOK, "authIndexBody.html", nil)
	c.HTML(http.StatusOK, "footer.html", nil)
}

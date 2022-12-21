package handler

import (
	"Diploma/pkg/theme"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"net/url"
	"time"
)

type Handler struct {
	delegate theme.Delegate
}

type Response struct {
	Message string
}

func (h *Handler) InitThemeRoutes(router *gin.Engine) {
	themeRouter := router.Group("/theme")
	{
		themeRouter.GET("/", h.GetInfo)
	}
}

func SetupThemeHandler(themeDelegate theme.Delegate) Handler {
	return Handler{
		delegate: themeDelegate,
	}
}

func (h *Handler) GetInfo(c *gin.Context) {
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	head := http.Header{}
	head.Add("Content-Type", "application/json")
	head.Add("Accept", "*/*")
	head.Add("Authorization", viper.GetString("yadisk.oauth"))
	rawUrl, err := url.Parse("https://cloud-api.yandex.net/v1/disk/")
	if err != nil {
		return
	}
	req := &http.Request{URL: rawUrl, Header: head}
	response, err := client.Do(req)
	if err != nil {
		return
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(response.Body)
	if err != nil {
		return
	}
	s := buf.String()
	fmt.Println(s) //s is a json!!!
	c.JSON(http.StatusOK, Response{Message: s})
}

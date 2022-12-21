package server

import (
	"Diploma/app/modules"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Start(handlers modules.HandlerModule) (err error) {
	router := SetupGinRouter(handlers)
	if err = router.Run(viper.GetString("server.address")); err != nil {
		return err
	}
	return
}

func SetupGinRouter(handlers modules.HandlerModule) *gin.Engine {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	handlers.UserHandler.InitUserRoutes(router)
	handlers.ThemeHandler.InitThemeRoutes(router)

	return router
}

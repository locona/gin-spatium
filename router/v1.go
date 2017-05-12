package router

import (
	"github.com/gin-gonic/gin"
	users "github.com/locona/gin-spatium/api/users/v1"
)

func V1(app *gin.Engine) *gin.Engine {
	api := app.Group("/v1")
	{
		api.GET("/users", users.Index)
	}
	return app
}

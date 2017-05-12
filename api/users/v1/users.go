package users

import (
	"github.com/gin-gonic/gin"
	"github.com/locona/gin-spatium/models"
)

func Index(ctx *gin.Context) {
	users := models.MockUsers()
	ctx.JSON(200, users)
}

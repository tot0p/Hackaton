package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tot0p/env"
	"hackaton/utils/db/mongodb"
)

func DataSetNamesAPIController(ctx *gin.Context) {
	names, err := mongodb.DB.GetAllCollectionsNames(env.Get("DB_MONGODB"))
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, names)
}

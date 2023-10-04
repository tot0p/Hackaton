package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tot0p/env"
	"hackaton/utils/db/mongodb"
)

func DataSetByNameGetFieldsAPIController(ctx *gin.Context) {
	name := ctx.Param("name")
	data, err := mongodb.DB.GetFields(env.Get("DB_MONGODB"), name)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, data)
}

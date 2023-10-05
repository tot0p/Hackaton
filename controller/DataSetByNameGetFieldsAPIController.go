package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tot0p/env"
	"hackaton/utils/db/mongodb"
	"strconv"
)

// DataSetByNameGetFieldsAPIController handles the GET request to get the fields of a dataset
func DataSetByNameGetFieldsAPIController(ctx *gin.Context) {
	name := ctx.Param("name")
	_id := true
	_mulfields := true
	var err error
	rmId := ctx.Query("rmId")
	if rmId != "" {
		_id, err = strconv.ParseBool(rmId)
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
	}
	rmMulFields := ctx.Query("rmMulFields")
	if rmMulFields != "" {
		_mulfields, err = strconv.ParseBool(rmMulFields)
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
	}
	data, err := mongodb.DB.GetFields(env.Get("DB_MONGODB"), name, _id, _mulfields)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, data)
}

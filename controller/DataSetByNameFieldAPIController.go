package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tot0p/env"
	"go.mongodb.org/mongo-driver/bson"
	"hackaton/utils/db/mongodb"
)

// DataSetByNameFiltersAPIController handles the GET request to get the data of a dataset by name and filters
func DataSetByNameFieldAPIController(ctx *gin.Context) {
	name := ctx.Param("name")
	field := ctx.Param("field")
	value := ctx.Param("value")

	data, err := mongodb.DB.Find(env.Get("DB_MONGODB"), name, bson.M{field: value})
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, data)

}

package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tot0p/env"
	"go.mongodb.org/mongo-driver/bson"
	"hackaton/utils/db/mongodb"
)

// DataSetByNameFiltersAPIController handles the GET request to get the data of a dataset by name and filters
func DataSetByNameFilterAPIController(ctx *gin.Context) {
	name := ctx.Param("name")

	// verify if param get are 'q'
	q := ctx.Query("q")
	if q == "" {
		ctx.JSON(400, gin.H{
			"error": "q is required",
		})
		return
	}

	// unmarshal q to bson.M
	filter := bson.M{}
	err := bson.UnmarshalExtJSON([]byte(q), true, &filter)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	data, err := mongodb.DB.FindOne(env.Get("DB_MONGODB"), name, filter)
	ctx.JSON(200, data)

}

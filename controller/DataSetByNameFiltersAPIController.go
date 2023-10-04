package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tot0p/env"
	"go.mongodb.org/mongo-driver/bson"
	"hackaton/utils/db/mongodb"
)

func DataSetByNameFiltersAPIController(ctx *gin.Context) {
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
	data, err := mongodb.DB.Find(env.Get("DB_MONGODB"), name, filter)
	ctx.JSON(200, data)

}

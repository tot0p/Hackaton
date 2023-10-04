package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"os"
)

const filePath = "src/data/categories.json"

func DataSetCategoriesAPIController(ctx *gin.Context) {
	jsonData, err := os.ReadFile(filePath)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	var data []map[string]interface{}
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, data)

}

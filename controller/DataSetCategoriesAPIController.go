package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"os"
)

const filePathCategory = "src/data/category.json"

// DataSetCategoriesAPIController handles the GET request to get the graph
func DataSetCategoriesAPIController(ctx *gin.Context) {
	jsonData, err := os.ReadFile(filePathCategory)
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

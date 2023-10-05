package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"os"
)

// DataSetByCategoryAPIController handles the GET request to get the graph by name
func DataSetByCategoryAPIController(ctx *gin.Context) {
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

	for _, v := range data {
		if v["theme"] == ctx.Param("category") {
			ctx.JSON(200, v)
			return
		}
	}

	ctx.JSON(404, gin.H{
		"error": "Category not found",
	})
}

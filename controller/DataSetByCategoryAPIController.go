package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"os"
)

func DataSetByCategoryAPIController(ctx *gin.Context) {
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

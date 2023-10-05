package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"os"
)

// GraphByNameAPIController handles the GET request to get the graph by name
func GraphByNameAPIController(ctx *gin.Context) {

	jsonData, err := os.ReadFile(filePathGraph)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	var data map[string]interface{}
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	if d, ok := data[ctx.Param("name")]; ok {
		ctx.JSON(200, d)
		return
	}

	ctx.JSON(404, gin.H{
		"error": "Category not found",
	})
}

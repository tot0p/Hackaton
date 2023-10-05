package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"os"
)

const filePathGraph = "src/data/graph.json"

func GraphAPIController(ctx *gin.Context) {
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
	ctx.JSON(200, data)
}

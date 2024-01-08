package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetIntParam(ctx *gin.Context, key string, defaultValue int) int {
	value, err := strconv.Atoi(ctx.Query(key))
	if err != nil {
		return defaultValue
	}
	return value
}

func GetStringParam(ctx *gin.Context, key, defaultValue string) string {
	value := ctx.Query(key)
	if value == "" {
		return defaultValue
	}
	return value
}

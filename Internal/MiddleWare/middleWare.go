package MiddleWare

import "github.com/gin-gonic/gin"

func checkHeader(ctx *gin.Context) bool {
	var result = false
	if val, exist := ctx.Request.Header["X-PING"]; exist && containsElem(val) {
		result = true
	}
	return result
}

func containsElem(slice []string) bool {
	var result = false
	for _, a := range slice {
		if a == "ping" {
			result = true
		}
	}
	return result
}

func SendHeader(ctx *gin.Context) {
	if checkHeader(ctx) {
		ctx.Writer.Header().Set("X-PONG", "pong")
	}
}

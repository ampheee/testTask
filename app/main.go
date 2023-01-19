package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"webPageComparer/Internal/MiddleWare"
	"webPageComparer/Internal/calculator"
)

func main() {
	r := gin.Default()
	r.GET("/when/:year", func(ctx *gin.Context) {
		defer ctx.Request.Body.Close()
		year, err := calculator.GetYear(ctx)
		var ans = []byte("")
		if err == nil {
			obtainedDateTime := calculator.FormateDate(year)
			ans = []byte(strconv.Itoa(calculator.Since(obtainedDateTime)))
		}
		MiddleWare.SendHeader(ctx)
		ctx.Data(http.StatusOK, "", ans)
	})

	err := r.Run(":8080")
	if err != nil {
		panic("We can`t start router!")
	}
}

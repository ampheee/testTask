package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	format  = "01.02.2006"
	ERROR   = "ERROR! "
	WARNING = "WARNING! "
)

func main() {
	r := gin.Default()
	r.GET("/when/:year", func(ctx *gin.Context) {
		defer ctx.Request.Body.Close()
		year, err := getYear(ctx)
		var ans = []byte("")
		if err == nil {
			obtainedDateTime := FormateDate(year)
			ans = []byte(strconv.Itoa(calculateSince(obtainedDateTime)))
		}
		ctx.Data(http.StatusOK, "", ans)
	})
	r.Run(":8080")
}

func getYear(ctx *gin.Context) (string, error) {
	year := ctx.Param("year")
	_, err := strconv.ParseInt(year, 10, 0)
	if err != nil {
		log.Printf("\n[%s]%s", ERROR, "Wrong type of year")
	}
	return year, err
}

func calculateSince(obtainedDateTime time.Time) int {
	days := time.Now().Sub(obtainedDateTime)
	return int(days.Hours() / 24)
}

func FormateDate(year string) time.Time {
	obtainedDateFormatted := fmt.Sprintf("%s.%s.%s", "01", "01", year)
	obtainedDateTime, err := time.Parse(format, obtainedDateFormatted)
	if err != nil {
		log.Printf("\n\n[%s] Cant parse date from page\n%s\n\n", WARNING, err)
	}
	return obtainedDateTime
}

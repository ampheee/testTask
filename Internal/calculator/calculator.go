package calculator

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math"
	"strconv"
	"time"
)

const (
	format  = "01.02.2006"
	ERROR   = "ERROR! "
	WARNING = "WARNING! "
)

func GetYear(ctx *gin.Context) (string, error) {
	year := ctx.Param("year")
	_, err := strconv.ParseInt(year, 10, 0)
	if err != nil {
		log.Printf("\n[%s]%s", ERROR, "Wrong type of year")
	}
	return year, err
}

func Since(obtainedDateTime time.Time) int {
	days := time.Now().Sub(obtainedDateTime)
	return int(math.Abs(days.Hours() / 24))
}

func FormateDate(year string) time.Time {
	obtainedDateFormatted := fmt.Sprintf("%s.%s.%s", "01", "01", year)
	obtainedDateTime, err := time.Parse(format, obtainedDateFormatted)
	if err != nil {
		log.Printf("\n\n[%s] Cant parse date from page\n%s\n\n", WARNING, err)
	}
	return obtainedDateTime
}

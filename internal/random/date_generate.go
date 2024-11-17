package random

import (
	"github.com/brianvoe/gofakeit/v7"
	"time"
)

//func for generating correct date

func GenerateDateAfter2010() time.Time {
	year := gofakeit.Number(2011, 2024)

	startDate := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Now().Add(-1)

	return gofakeit.DateRange(startDate, endDate)
}

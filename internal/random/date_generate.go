package random

import (
	"github.com/brianvoe/gofakeit/v7"
	"time"
)

func GenerateDateAfter2010() time.Time {
	year := gofakeit.Number(2011, 2024)

	startDate := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(year, time.December, 31, 23, 59, 59, 0, time.UTC)

	return gofakeit.DateRange(startDate, endDate)
}

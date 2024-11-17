package random

import (
	"math/rand"
	"time"
)

// func for generate correct order date

func GenerateRandomOrderDate(times []time.Time) []time.Time {
	now := time.Now()

	randomDates := make([]time.Time, len(times))

	for i, t := range times {
		if t.After(now) {
			randomDates[i] = now
			continue
		}
		duration := now.Sub(t)
		randomOffset := time.Duration(rand.Int63n(int64(duration)))
		randomDates[i] = t.Add(randomOffset)
	}

	return randomDates
}

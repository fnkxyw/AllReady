package random

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateWorkHours() string {

	rand.Seed(time.Now().UnixNano())

	startHour := rand.Intn(8) + 6

	endHour := rand.Intn(8) + 14

	return fmt.Sprintf("%02d-%02d", startHour, endHour)
}

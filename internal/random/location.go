package random

import "math/rand"

func GenerateRandomLocation() string {
	locations := []string{
		"Table #1 (Window side)",
		"Table #2 (Center of the room)",
		"Table #3 (Near the wall)",
		"Bar Counter",
		"Kitchen",
		"VIP Zone",
		"Terrace",
		"Coat Room",
		"Restroom",
	}

	randomIndex := rand.Intn(len(locations))

	return locations[randomIndex]
}

package helpers

import (
	"math/rand"
	"time"
)

func GenerateRandom(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func CheckWater(water int) string {

	if water <= 5 {
		return "aman ygy"
	} else if water > 5 && water <= 7 {
		return "siaga gaes"
	} else {
		return "awas tengge..."
	}

}

func CheckAir(air int) string {
	if air <= 6 {
		return "aman ygy"
	} else if air >= 7 && air <= 14 {
		return "siaga gaes"
	} else {
		return "bahaya ges bisa terbang"
	}
}

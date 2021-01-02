package main

import "fmt"

func main() {
	var deg, dis int
	fmt.Scan(&deg, &dis)
	directions := []string{
		"N", "NNE", "NE", "ENE",
		"E", "ESE", "SE", "SSE",
		"S", "SSW", "SW", "WSW",
		"W", "WNW", "NW", "NNW",
	}
	powers := []int{2, 15, 33, 54, 79, 107, 138, 171, 207, 244, 284, 326, 2000}
	t := (dis*100/60 + 5) / 10
	direction := directions[(deg*10+1125)/2250%16]
	power := 0
	for i, v := range powers {
		if t <= v {
			power = i
			if i == 0 {
				direction = "C"
			}
			break
		}
	}

	fmt.Println(direction, power)
}

package main

import "fmt"

func main() {
	var n int
	var s string
	var sh, sm, eh, em int
	ts := [1500]int{}
	start := -1

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&s)
		fmt.Sscanf(s, "%02d%02d-%02d%02d", &sh, &sm, &eh, &em)

		sm -= sm % 5
		if em%5 != 0 {
			em += 5 - (em % 5)
		}
		if em >= 60 {
			eh++
			em = 0
		}

		ts[sh*60+sm]++
		ts[eh*60+em+1]--
	}

	for i := 0; i < len(ts)-1; i++ {
		ts[i+1] += ts[i]
	}

	for i := 0; i <= 1441; i++ {
		if start == -1 && ts[i] > 0 {
			start = i
		} else if start != -1 && ts[i] == 0 {
			ans := fmt.Sprintf("%02d%02d-%02d%02d", start/60, start%60, (i-1)/60, (i-1)%60)
			fmt.Println(ans)
			start = -1
		}
	}
}

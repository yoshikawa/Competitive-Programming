package main

import "fmt"

func main() {
	var x string
	fmt.Scan(&x)

	for i := 0; i < len(x); i++ {
		s := x[i]
		if (s != 'a') && (s != 'i') && (s != 'u') && (s != 'e') && (s != 'o') {
			fmt.Printf("%c", s)
		}
	}
	fmt.Println("")
}

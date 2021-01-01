package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var input int

	fmt.Fscan(r, &input)

	fmt.Fprintln(w, calc(input))
}

func calc(input int) (ans string) {
	if input < 100 {
		ans = "00"
	} else if input <= 5000 {
		if input < 1000 {
			ans = "0"
		}
		ans += strconv.Itoa(input / 100)
	} else if input <= 30000 {
		ans = strconv.Itoa(input/1000 + 50)
	} else if input <= 70000 {
		ans = strconv.Itoa((input/1000-30)/5 + 80)
	} else {
		ans = "89"
	}

	return ans
}

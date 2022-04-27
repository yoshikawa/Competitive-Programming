package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var h1, h2 int

	fmt.Fscan(r, &h1)
	fmt.Fscan(r, &h2)
	fmt.Fprintln(w, diff(h1, h2))
}

func diff(h1, h2 int) (abs int) {
	return h1 - h2
}

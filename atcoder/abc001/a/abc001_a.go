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
	fmt.Fprintln(w, Diff(h1, h2))
}

// Diff returns Snow depth's difference.
func Diff(h1, h2 int) (abs int) {
	return h1 - h2
}

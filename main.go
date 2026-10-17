package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	// TODO: accumulate the sum of 1..n with a loop, and print that instead of the placeholder.

	acc := 0
	for i := 1; i <= n; i++ {
		acc = acc + i
	}

	fmt.Printf("%d\n", acc)
}

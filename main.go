package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	mustPrintNum := true

	// TODO: choose the right label for n instead of always printing it.
	if n%3 == 0 {
		fmt.Print("Fizz")
		mustPrintNum = false
	}

	if n%5 == 0 {
		fmt.Print("Buzz")
		mustPrintNum = false
	}

	if mustPrintNum {
		fmt.Printf("%d", n)
	}
}

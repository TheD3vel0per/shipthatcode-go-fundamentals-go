package main

import "fmt"

func main() {
	var first, second int
	fmt.Scan(&first)
	fmt.Scan(&second)

	// TODO: line 1 - print the sum of the two integers.
	sum := int(first) + int(second)
	fmt.Println(sum)

	// TODO: line 2 - print the result of subtracting the second from the first.
	diff := int(first) - int(second)
	fmt.Println(diff)

	// TODO: line 3 - print whether the two integers are equal.
	equal := first == second
	fmt.Println(equal)
}

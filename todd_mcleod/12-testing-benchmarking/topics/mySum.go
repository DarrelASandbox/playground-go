// Package topics (mySum.go)
package topics

import "fmt"

// MySum() adds up a set of numbers with values of type int
func MySum(xi ...int) int {
	fmt.Println("\n\nmySum:")

	sum := 0
	for _, v := range xi {
		sum += v
	}

	return sum
}

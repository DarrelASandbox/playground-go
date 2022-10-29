// Package topics (mySum.go)
package topics

// MySum() adds up a set of numbers with values of type int
func MySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}

	return sum
}

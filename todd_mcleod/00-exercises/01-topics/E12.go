// package word provides custom functions for working with strings and words
package topics

import (
	"fmt"

	with_tests "github.com/DarrelASandbox/playground-go/todd_mcleod/00-exercises/02-with-tests"
)

func E12() {
	fmt.Println("\n\n##################################################")

	countWords()
	calCenteredAvg()
}

func countWords() {
	fmt.Println("\n\ncountWords:")
	fmt.Println(with_tests.Count(with_tests.SunAlso))

	for k, v := range with_tests.UseCount(with_tests.SunAlso) {
		fmt.Println(v, k)
	}
}

func calCenteredAvg() {
	fmt.Println("\n\ncalCenteredAvg:")

	gen := func() [][]int {
		a := []int{1, 4, 6, 8, 100}
		b := []int{0, 8, 10, 1000}
		c := []int{9000, 4, 10, 8, 6, 12}
		d := []int{123, 744, 140, 200}
		e := [][]int{a, b, c, d}
		return e
	}

	xxi := gen()
	for _, v := range xxi {
		fmt.Println(with_tests.CenteredAvg(v))
	}
}

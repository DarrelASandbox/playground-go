package main

import (
	"fmt"
	"testing"

	"github.com/DarrelASandbox/playground-go/todd_mcleod/12-testing-benchmarking/topics"
)

type test struct {
	data   []int
	answer int
}

var tests = []test{
	test{[]int{21, 21}, 42},
	test{[]int{3, 4, 5}, 12},
	test{[]int{1, 1}, 2},
	test{[]int{-1, 0, 1}, 0},
}

func TestMySum(t *testing.T) {
	for i, v := range tests {
		x := topics.MySum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		} else {
			fmt.Print("Test Input ", i+1, ": ok")
		}
	}

	fmt.Print("\n\n")
}

func ExampleMySum() {
	fmt.Println(topics.MySum(2, 3))
	// Output:
	// 5
}

func BenchmarkMySum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, v := range tests {
			topics.MySum(v.data...)
		}
	}
}

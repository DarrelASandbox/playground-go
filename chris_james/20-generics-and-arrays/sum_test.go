package main

import "testing"

func TestReduce(t *testing.T) {
	/*
	  In the multiplication example, we show the reason for having a default value as an argument to `Reduce`. If we relied on Go's default value of 0 for `int`, we'd multiply our initial value by 0, and then the following ones, so you'd only ever get 0. By setting it to 1, the first element in the slice will stay the same, and the rest will multiply by the next elements.
	*/
	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}

		AssertEqual(t, Reduce([]int{1, 2, 3}, multiply, 1), 6)
	})

	t.Run("concatenate strings", func(t *testing.T) {
		concatenate := func(x, y string) string {
			return x + y
		}

		AssertEqual(t, Reduce([]string{"a", "b", "c"}, concatenate, ""), "abc")
	})
}

func TestFind(t *testing.T) {
	t.Run("find first even number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		firstEvenNumber, found := Find(numbers, func(x int) bool { return x%2 == 0 })
		AssertTrue(t, found)
		AssertEqual(t, firstEvenNumber, 2)
	})
}

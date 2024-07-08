package main

/*
Recurring pattern:
- Create some kind of "initial" result value.
- Iterate over the collection, applying some kind of operation (or function) to the result and
  the next item in the slice, setting a new value for the result
- Return the result.

`Sum` and `SumAllTails` now describe the behavior of their computations as
the functions declared on their first lines respectively.
The act of running the computation on the collection is abstracted away in `Reduce`.
*/

// Sum calculates the total from a slice of numbers.
func Sum(numbers []int) int {
	add := func(acc, x int) int { return acc + x }
	return Reduce(numbers, add, 0)
}

// SumAllTails calculates the sums of all but the first number given a collection of slices.
func SumAllTails(numbers ...[]int) []int {
	sumTail := func(acc, x []int) []int {
		if len(x) == 0 {
			return append(acc, 0)
		} else {
			tail := x[1:]
			return append(acc, Sum(tail))
		}
	}

	return Reduce(numbers, sumTail, []int{})
}

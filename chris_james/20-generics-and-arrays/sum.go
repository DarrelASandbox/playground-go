package main

/*
Recurring pattern:
- Create some kind of "initial" result value.
- Iterate over the collection, applying some kind of operation (or function) to the result and
  the next item in the slice, setting a new value for the result
- Return the result.
*/

// Sum calculates the total from a slice of numbers.
func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}

	return sum
}

// SumAllTails calculates the sums of all but the first number given a collection of slices.
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}

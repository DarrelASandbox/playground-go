package sum

/*
The size of an array is encoded in its type.
If you try to pass an [4]int into a function that expects [5]int, it won't compile.

Go has slices which do not encode the size of the collection and instead can have any size.
*/
func Sum(numbers []int) int {
	sum := 0

	// Blank identifier
	for _, number := range numbers {
		sum += number
	}
	return sum
}

/*
As mentioned, slices have a capacity.
If you have a slice with a capacity of 2 and try to do mySlice[10] = 1 you will get a runtime error.
However, you can use the append function which takes a slice and a new value,
then returns a new slice with all the items in it.
*/
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

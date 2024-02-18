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

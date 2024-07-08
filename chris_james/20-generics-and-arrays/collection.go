package main

/*
`Reduce` captures the essence of the pattern, it's a function that takes a collection,
an accumulating function, an initial value, and returns a single value.
There's no messy distractions around concrete types.

We've added a second type constraint which has allowed us to loosen the constraints on `Reduce`.
This allows people to `Reduce` from a collection of `A` into a `B`. In our case from `Transaction` to `float64`.
*/
func Reduce[A, B any](collection []A, f func(B, A) B, initialValue B) B {
	var result = initialValue
	for _, x := range collection {
		result = f(result, x)
	}
	return result
}

func Find[A any](items []A, predicate func(A) bool) (value A, found bool) {
	for _, v := range items {
		if predicate(v) {
			return v, true
		}
	}
	return
}

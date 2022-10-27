package topics

import "fmt"

/*
Recover is a built-in function that regains control of a panicking goroutine. Recover is only useful inside deferred functions. During normal execution, a call to recover will return nil and have no other effect. If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.
*/

func Recovering() {
	fmt.Println("\n\ntopics.Recovering():")

	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	fmt.Println("Calling g.")
	g(0)

	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}

	// At the top of the call stack, executing any defer statement first
	// LIFO: 3, 2, 1 & 0
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

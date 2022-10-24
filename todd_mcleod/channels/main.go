package main

import "fmt"

func main() {
	// blockingDeadlock()
	successfulBuffer()
	// unsuccessfulBuffer()
	directionalChannel()
	// invalidDirectionalChannel1()
	// invalidDirectionalChannel2()
}

func blockingDeadlock() {
	fmt.Println("\n\nblockingDeadlock:")
	c := make(chan int)
	c <- 42

	// fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-c)
}

func successfulBuffer() {
	fmt.Println("\n\nsuccessfulBuffer:")
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

func unsuccessfulBuffer() {
	fmt.Println("\n\nunsuccessfulBuffer:")
	c := make(chan int, 2)

	// capacity is 2, added 2 values in buffer
	c <- 42
	c <- 43

	// taken out 1 value
	// fatal error: all goroutines are asleep - deadlock!
	fmt.Println(<-c)

	// adding 1 more since buffer only has 1 value
	c <- 44

	//adding 1 more, now this should block because added more than capacity
	c <- 45

	fmt.Println("this doesn't print because the now its blocking", <-c)
}

func directionalChannel() {
	fmt.Println("\n\ndirectionalChannel:")
	c := make(chan int, 2)

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)

	fmt.Println("-----")
	fmt.Printf("%T\n", c)
}

/*
func invalidDirectionalChannel1() {
	fmt.Println("\n\ninvalidDirectionalChannel1:")
	c := make(chan<- int, 2)

	c <- 42
	c <- 43

	// invalid operation: cannot receive from send-only channel c (variable of type chan<- int)
	fmt.Println(<-c)
	fmt.Println(<-c)

	fmt.Println("-----")
	fmt.Printf("%T\n", c)
}
*/

/*
func invalidDirectionalChannel2() {
	fmt.Println("\n\ninvalidDirectionalChannel2:")
	c := make(<-chan int, 2)

	// invalid operation: cannot send to receive-only channel c (variable of type <-chan int)
	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)

	fmt.Println("-----")
	fmt.Printf("%T\n", c)
}
*/

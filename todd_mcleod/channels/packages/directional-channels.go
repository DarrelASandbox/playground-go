package packages

import (
	"fmt"
)

func DirectionalChannel() {
	fmt.Println("\n\npackages.DirectionalChannel():")
	c := make(chan int, 2)

	c <- 42
	c <- 43

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("-----")
	fmt.Printf("%T\n", c)
}

/*
func InvalidDirectionalChannel1() {
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
func InvalidDirectionalChannel2() {
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

func BidirectionalChannel() {
	fmt.Println("\n\npackages.BidirectionalChannel():")
	c := make(chan int)
	cr := make(<-chan int) // receive
	cs := make(chan<- int) // send

	fmt.Println("-----")
	fmt.Printf("c\t%T\n", c)
	fmt.Printf("c\t%T\n", cr)
	fmt.Printf("c\t%T\n", cs)

	// specific to general doesn't assign
	// c = cr
	// c = cs
}

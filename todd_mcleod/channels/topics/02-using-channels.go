package topics

import "fmt"

func SendReceive() {
	fmt.Println("\n\ntopics.SendReceive():")

	c := make(chan int)

	// send
	go routine1(c)
	// receive
	routine2(c)

	fmt.Println("about to exit in SendReceive() main func")
}

func routine1(c chan<- int) {
	fmt.Println("routine1:")
	c <- 42
}

func routine2(c <-chan int) {
	fmt.Println("routine2 has no goroutine for channel")
	fmt.Println(<-c)
}

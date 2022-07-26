package main

import (
	"fmt"

	"github.com/DarrelASandbox/playground-go/todd_mcleod/09-channels/topics"
)

/*
WaitGroups are just enough if you don’t need any data returned from a goroutine.
However, you’ll often need to pass data around when building concurrent applications, which channels are extremely helpful for.
*/

func main() {
	// blockingDeadlock()
	successfulBuffer()
	// unsuccessfulBuffer()

	topics.DirectionalChannel()
	// topics.InvalidDirectionalChannel1()
	// topics.InvalidDirectionalChannel2()
	topics.BidirectionalChannel()

	topics.SendReceive()
	topics.Range()
	topics.Select()
	topics.CommaWithBool()
	topics.CommaWithInt()
	topics.CommaOnly()

	topics.FanIn()
	topics.FanInBoring()
	topics.FanOut()

	topics.Ctx()
	topics.CtxWithGoRoutine()
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

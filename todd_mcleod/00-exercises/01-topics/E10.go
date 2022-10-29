package topics

import (
	"fmt"
	"runtime"
)

func E10() {
	fmt.Println("\n\n##################################################")
	channelEx1()
	channelEx2()
	channelEx3()
	channelEx4()
	channelEx5()
	channelEx6()
	channelEx7()
}

func channelEx1() {
	fmt.Println("\n\nchannelEx1:")

	// func literal, aka, anonymous self-executing func
	c1 := make(chan int)
	go func() {
		c1 <- 42
	}()

	// buffered channel
	c2 := make(chan int, 1)
	c2 <- 42

	fmt.Println("func literal:", <-c1)
	fmt.Println("buffered channel:", <-c2)
}

func channelEx2() {
	fmt.Println("\n\nchannelEx2:")

	// cs := make(chan<- int)
	cs := make(chan int)
	go func() {
		cs <- 42
	}()

	fmt.Println("bidirectional channel cs:", <-cs)
	fmt.Printf("%T\n", cs)

	// cr := make(<-chan int)
	cr := make(chan int)
	go func() {
		cr <- 42
	}()

	fmt.Println("\nbidirectional channel cr:", <-cr)
	fmt.Printf("%T\n", cr)
}

/*
channelEx3: In this program there are a total of two goroutines running. The "main" goroutine (that func main started in) calls gen to make a channel and start the other goroutine, then it goes and reads from the channel until it is closed using the range loop on it. The generator goroutine needs to close the channel inside itself then, so that the main goroutine can exit the loop. If the channel was closed outside the goroutine, then the goroutine would start attempting to send data into a closed channel, which would cause a panic error.
*/

func channelEx3() {
	fmt.Println("\n\nchannelEx3: range over channel")

	// Using Lambdas for nested functions
	gen := func() <-chan int {
		c := make(chan int)

		go func() {
			for i := 0; i < 3; i++ {
				c <- i
			}
			close(c)
		}()

		return c
	}

	receive := func(c <-chan int) {
		for v := range c {
			fmt.Println(v)
		}
	}

	c := gen()
	receive(c)

	fmt.Println("about to exit channelEx3")
}

func channelEx4() {
	fmt.Println("\n\nchannelEx4: selecting channels")

	// Using Lambdas for nested functions
	gen := func(q chan int) <-chan int {
		c := make(chan int)

		go func() {
			for i := 0; i < 3; i++ {
				c <- i
			}

			// When a channel e.g. close(c) has been closed, reading from it returns the zero value and the ok flag to false.
			// 0 false
			// close(c)
			close(q)
		}()

		return c
	}

	receive := func(c, q <-chan int) {
		for {
			select {
			case v := <-c:
				fmt.Println(v)

			case <-q:
				return
			}
		}
	}

	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit channelEx4")
}

func channelEx5() {
	fmt.Println("\n\nchannelEx5: comma, ok")

	c := make(chan int, 1)
	c <- 42

	v, ok := <-c
	fmt.Println(v, ok)
	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}

/*
https://go.dev/ref/spec
A new, initialized channel value can be made using the built-in function make, which takes the channel type and an optional capacity as arguments

The capacity, in number of elements, sets the size of the buffer in the channel. If the capacity is zero or absent, the channel is unbuffered and communication succeeds only when both a sender and receiver are ready. Otherwise, the channel is buffered and communication succeeds without blocking if the buffer is not full (sends) or not empty (receives). A nil channel is never ready for communication.
*/

func channelEx6() {
	fmt.Println("\n\nchannelEx6: looping buffered channel")

	sameCapacity := 4
	c := make(chan int, sameCapacity)
	for i := 0; i < sameCapacity; i++ {
		c <- i
	}
	close(c)

	for v := range c {
		fmt.Println(v)
	}
}

func channelEx7() {
	fmt.Println("\n\nchannelEx7:")

	// Using Lambdas for nested functions
	gen := func(x, y int) <-chan int {
		c := make(chan int)

		for i := 0; i < x; i++ {
			go func() {
				for j := 0; j < y; j++ {
					c <- j
				}
			}()
			fmt.Println("ROUTINES", runtime.NumGoroutine())
		}

		return c
	}

	x := 3 // number of goroutines
	y := 4 // number of values
	c := gen(x, y)

	for i := 0; i < x*y; i++ {
		fmt.Println(i, <-c)
	}

	fmt.Println("ROUTINES", runtime.NumGoroutine())
}

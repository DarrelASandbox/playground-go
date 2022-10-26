package topics

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func FanIn() {
	fmt.Println("\n\ntopics.FanIn():")

	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go send4(even, odd)
	go receive4(even, odd, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("about to exit in FanIn() main func")
}

func send4(e, o chan<- int) {
	for i := 1; i <= 16; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}

	close(e)
	close(o)
}

func receive4(e, o <-chan int, fi chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range e {
			fi <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range o {
			fi <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fi)
}

/* #################################################################################################### */

/*
code source:
Rob Pike
https://talks.golang.org/2012/concurrency.slide#25

source:
https://blog.golang.org/pipelines
*/

func FanInBoring() {
	fmt.Println("\n\ntopics.FanInBoring():")

	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 1; i <= 17; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("You're both boring; I'm leaving.")
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			c <- <-input1 // pull value off input1 and put into channel c
		}
	}()

	go func() {
		for {
			c <- <-input2
		}
	}()

	return c
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}

/* #################################################################################################### */

func FanOut() {
	fmt.Println("\n\ntopics.FanOut():")

	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)
	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("about to exit in FanOut() main func")
}

func populate(c chan int) {
	for i := 0; i < 18; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	for v := range c1 {
		wg.Add(1)
		go func(v2 int) {
			c2 <- timeConsumingWork(v2)
			wg.Done()
		}(v)
	}

	wg.Wait()
	close(c2)
}

func timeConsumingWork(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}

/* #################################################################################################### */

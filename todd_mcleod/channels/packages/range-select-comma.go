package packages

import "fmt"

func Range() {
	fmt.Println("\n\npackages.Range():")

	c := make(chan int)

	// send
	go func() {
		for i := 1; i <= 13; i++ {
			c <- i
		}

		// required to prevent fatal error: all goroutines are asleep - deadlock!
		close(c)
	}()

	// receive
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit in Range() main func")
}

/* #################################################################################################### */

func Select() {
	fmt.Println("\n\npackages.Select():")

	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go send1(even, odd, quit)
	// receive
	receive1(even, odd, quit)

	fmt.Println("about to exit in Select() main func")
}

func send1(e, o, q chan<- int) {
	for i := 1; i <= 15; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	// remove close() to remove 0s
	close(e)
	close(o)

	q <- 0
	close(q)
}

func receive1(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("from the even channel:\t", v)
		case v := <-o:
			fmt.Println("from the odd channel:\t", v)
		case v := <-q:
			fmt.Println("from the quit channel:\t", v)
			return
		}
	}
}

/* #################################################################################################### */

func CommaWithBool() {
	fmt.Println("\n\npackages.CommaWithBool():")

	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	// send
	go send2(even, odd, quit)
	// receive
	receive2(even, odd, quit)

	fmt.Println("about to exit in CommaWithBool() main func")
}

func send2(e, o chan<- int, q chan<- bool) {
	for i := 1; i <= 16; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}

	close(q)
}

func receive2(e, o <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Println("from the even channel:\t", v)
		case v := <-o:
			fmt.Println("from the odd channel:\t", v)
		case i, ok := <-q:
			if !ok {
				fmt.Println("from the NOT comma ok:\t", i, "\tok:", ok)
				return
			} else {
				fmt.Println("from the comma ok:\t", i)
			}
		}
	}
}

/* #################################################################################################### */

func CommaWithInt() {
	fmt.Println("\n\npackages.CommaWithInt():")

	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go send3(even, odd, quit)
	// receive
	receive3(even, odd, quit)

	fmt.Println("about to exit in CommaWithInt() main func")
}

func send3(e, o, q chan<- int) {
	for i := 1; i <= 16; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}

	close(q)
}

func receive3(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("from the even channel:\t", v)
		case v := <-o:
			fmt.Println("from the odd channel:\t", v)
		case i, ok := <-q:
			if !ok {
				fmt.Println("from the NOT comma ok:\t", i, "\tok:", ok)
				return
			} else {
				fmt.Println("from the comma ok:\t", i)
			}
		}
	}
}

/* #################################################################################################### */

func CommaOnly() {
	fmt.Println("\n\npackages.CommaOnly():")

	c := make(chan int)
	go func() {
		c <- 42
		close(c)
	}()

	v, ok := <-c
	fmt.Println(v, ok)

	v, ok = <-c
	fmt.Println(v, ok)
}

/* #################################################################################################### */

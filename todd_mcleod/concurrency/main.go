package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/DarrelASandbox/playground-go/todd_mcleod/concurrency/packages"
)

var wg sync.WaitGroup

func main() {
	fooNeverRunWithoutWG()
	fooRunsWithWG()
	raceConditionFunc()

	// Setup go.work -> setup go.mod -> setup packages -> local import functions
	packages.Mutex()
	packages.Atomic()
}

func fileLine() string {
	_, _, fileLine, ok := runtime.Caller(1)
	var s string
	if ok {
		s = fmt.Sprintf("Line %d: ", fileLine)
	} else {
		s = ""
	}
	return s
}

func foo() {
	for i := 0; i < 2; i++ {
		fmt.Println("fool from foo()", i)
	}
}

func fooWithWG() {
	for i := 0; i < 2; i++ {
		fmt.Println("foo", i)
	}

	fmt.Println()
	wg.Done()
}

func bar() {
	for j := 0; j < 3; j++ {
		fmt.Println("bar", j)
	}
}

func fooNeverRunWithoutWG() {
	fmt.Println("\n\nfooNeverRunWithoutWG:")
	fmt.Println("OS:\t\t", runtime.GOOS)
	fmt.Println("ARCH:\t\t", runtime.GOARCH)
	fmt.Println("CPUs:\t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())
	fmt.Println()

	go foo() // launch the second goroutine
	bar()

	fmt.Println()
	fmt.Println("CPUs:\t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())
	// Program exited including second goroutine
	// So we need WaitGroup
}

func fooRunsWithWG() {
	fmt.Println("\n\nfooRunsWithWG:")
	fmt.Println("OS:\t\t", runtime.GOOS)
	fmt.Println("ARCH:\t\t", runtime.GOARCH)
	fmt.Println("CPUs:\t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())
	fmt.Println()

	wg.Add(1) // Add 1 to WaitGroup
	go fooWithWG()
	wg.Wait() // Wait for fooWithWG
	bar()

	fmt.Println()
	fmt.Println("CPUs:\t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())
}

func raceConditionFunc() {
	fmt.Println("\n\nraceConditionFunc:")
	fmt.Println("CPUs:\t\t", runtime.NumCPU())
	fmt.Println("Goroutines:\t", runtime.NumGoroutine())

	counter := 0
	gs := 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < 100; i++ {
		go func() {
			v := counter
			time.Sleep(time.Second)
			runtime.Gosched()

			v++
			counter = v
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(fileLine(), "Goroutines:\t", runtime.NumGoroutine())
	fmt.Println(fileLine(), "counter:", counter)
}

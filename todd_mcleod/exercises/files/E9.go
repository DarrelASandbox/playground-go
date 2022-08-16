package files

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func E9() {
	fmt.Println("\n\n##################################################")
	raceCondition()
	raceConditionMutexFix()
	raceConditionAtomicFix()
}

func raceCondition() {
	fmt.Println("\n\nraceCondition:")
	var wg sync.WaitGroup

	counter := 0
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			fmt.Println(counter)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Final value:", counter)
}

func raceConditionMutexFix() {
	fmt.Println("\n\nraceConditionMutexFix:")
	var wg sync.WaitGroup

	counter := 0
	gs := 100
	wg.Add(gs)
	var m sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			m.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			m.Unlock()
			fmt.Println(counter)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Final value:", counter)
}

func raceConditionAtomicFix() {
	fmt.Println("\n\nraceConditionAtomicFix:")
	var wg sync.WaitGroup
	var counter int64

	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println(atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Final value:", counter)
}

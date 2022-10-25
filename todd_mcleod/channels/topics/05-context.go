package topics

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func Ctx() {
	fmt.Println("\n\ntopics.Ctx():")

	ctx := context.Background()

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n\n", ctx)

	ctx, cancel := context.WithCancel(ctx)

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("context cancel:\t", cancel)
	fmt.Printf("context type:\t%T\n\n", cancel)

	cancel()

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("context cancel:\t", cancel)
	fmt.Printf("context type:\t%T\n\n", cancel)
}

func CtxWithGoRoutine() {
	fmt.Println("\n\ntopics.CtxWithGoRoutine():")

	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("error check 1:", ctx.Err())
	fmt.Println("NumGoroutine1:", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return

			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("working", n)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("error check 2:", ctx.Err())
	fmt.Println("NumGoroutine2:", runtime.NumGoroutine())

	fmt.Println("cancelling context")
	cancel()
	fmt.Println("cancelled context")

	time.Sleep(time.Second * 2)
	fmt.Println("error check 3:", ctx.Err())
	fmt.Println("NumGoroutine3:", runtime.NumGoroutine())
}

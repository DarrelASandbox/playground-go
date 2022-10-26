package main

import "fmt"

func main() {
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println("i:", i, " j:", j)
		}
	}

	fmt.Println("")

	x := 1
	for x < 3 {
		fmt.Println("x:", x)
		x++
	}

	fmt.Println("")

	for i := 33; i <= 122; i++ {
		fmt.Printf("%v\t%#x\t%#U\n", i, i, i)
	}
}

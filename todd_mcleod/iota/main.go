package main

import (
	"fmt"
)

func main() {

	const (
		_        = iota
		kilobyte = 1 << (iota * 10)
		megabyte
		gigabyte
	)
	fmt.Printf("%d\t%b\n", kilobyte, kilobyte)
	fmt.Printf("%d\t%b\n", megabyte, megabyte)
	fmt.Printf("%d\t%b\n", gigabyte, gigabyte)
}

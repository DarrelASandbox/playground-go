package main

import (
	"fmt"

	"github.com/DarrelASandbox/playground-go/todd_mcleod/12-testing-benchmarking/topics"
)

func main() {
	fmt.Println("\n\ntopics.mySum():")
	fmt.Println("2 + 3 =", topics.MySum(2, 3))
}

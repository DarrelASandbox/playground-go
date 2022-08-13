package files

import "fmt"

func E3() {
	fmt.Println("\n\n##################################################")
	fmt.Println("E3:")

	x := "Moneypenny"

	if x == "Moneypenny" {
		fmt.Println(x)
	} else if x == "James Bond" {
		fmt.Println("BONDDONBONDONBOND", x)
	} else {
		fmt.Println("neither")
	}

	fmt.Println()

	switch {
	case false:
		fmt.Println("should not print")
	case true:
		fmt.Println("should print")
	}

	fmt.Println()

	favSport := "surfing"
	switch favSport {
	case "skiing":
		fmt.Println("go to the mountains!")
	case "swimming":
		fmt.Println("go to the pool!")
	case "surfing":
		fmt.Println("go to hawaii!")
	case "wingsuit flying":
		fmt.Println("what would you like me to say at your funeral?")
	}
}

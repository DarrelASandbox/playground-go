package topics

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func E4() {
	fmt.Println("\n\n##################################################")
	fmt.Println("E4:")

	overwriteSlice()
	statesInUSA()
	sliceOfSlice()
	mapping()
}

func overwriteSlice() {
	fmt.Println("\n\noverwriteSlice:")
	x := make([]string, 5)
	fmt.Println("x := make([]string, 5, 5)", x)

	x = []string{"one", "two", "three", "four"} // 4 elements

	fmt.Println("capacity:", cap(x)) // prints 4 instead of 5

	x = append(x, "five") // exceeds previous capacity of 4
	// thus allocates a new underlying array
	// with capacity 8

	fmt.Println("capacity:", cap(x))
	// Prints 8. Which should still be 5
	// if only make would have been used
}

func statesInUSA() {
	fmt.Println("\n\nstatesInUSA:")
	y := make([]string, 50)
	fmt.Println("Check slice after make:\t\tlength:", len(y), "capacity:", cap(y))

	states := []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}

	fmt.Println()

	for i, v := range states {
		y[i] = v
		fmt.Println(i, y[i])
	}

	fmt.Println("\nCheck slice after for loop:\tlength:", len(y), "capacity:", cap(y))
}

func sliceOfSlice() {
	fmt.Println("\n\nsliceOfSlice:")
	x := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Hello, James."},
	}

	for i, xs := range x {
		fmt.Println("record: ", i)
		for j, val := range xs {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
		}
	}
}

func mapping() {
	fmt.Println("\n\nmapping:")
	// key of string & slice of strings
	m := map[string][]string{
		"bond_james":      {`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": {`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	m[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}
	delete(m, `no_dr`)

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 3, 3, 0, ' ', 0)

	for k, v := range m {
		fmt.Fprintln(w, "Record for", k, ":")
		for i, v2 := range v {
			fmt.Fprintln(w, "\t", i, "\t", v2)
		}
	}
	w.Flush()
}

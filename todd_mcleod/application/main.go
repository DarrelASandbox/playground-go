package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

// json is for unmarshalling
type Person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

// https://pkg.go.dev/sort@go1.18.4#example-package
// ByAge implements sort.Interface for []Person based on the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByName []Person

func (n ByName) Len() int           { return len(n) }
func (n ByName) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n ByName) Less(i, j int) bool { return n[i].First < n[j].First }

func main() {
	p1 := Person{First: "James", Last: "Bomb", Age: 81}
	p2 := Person{First: "Larry", Last: "Fairy", Age: 18}
	people := []Person{p1, p2}

	marshalFunc(people...)
	unmarshalFunc()
	sortFunc()
	customSortFunc(people...)
}

func marshalFunc(people ...Person) {
	fmt.Println("\n\nmarshalFunc:")

	fmt.Println(people)

	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}
	// returns [{},{}] if struct keys (first, last & age) is lowercase
	fmt.Println(string(bs))
}

func unmarshalFunc() {
	fmt.Println("\n\nunmarshalFunc:")
	s := `[{"First":"James","Last":"Bomb","Age":81},{"First":"Larry","Last":"Fairy","Age":18}]`
	bs := []byte(s)
	fmt.Printf("%T\ns type:", s)
	fmt.Printf("%T\nbs type:", bs)

	people := []Person{} // var people = []Person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}

	for i, v := range people {
		fmt.Println(i, v)
	}
}

func sortFunc() {
	fmt.Println("\n\nsortFunc:")
	xi := []int{4, 7, 3, 42, 99, 18, 16, 56, 12}
	xs := []string{"James", "Q", "M", "Larry", "Dr. No"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}

func customSortFunc(people ...Person) {
	fmt.Println("\n\ncustomSortFunc:")
	sort.Sort(ByAge(people))
	fmt.Println(people)
}

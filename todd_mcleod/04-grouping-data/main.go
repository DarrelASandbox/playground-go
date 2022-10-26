package main

import "fmt"

func main() {
	var a [3]int // array
	fmt.Println("array of size 3:", a)
	a[2] = 42
	fmt.Println("assigning value to array at specific index:", a)

	// a slice allows you to compose together many values of the SAME type
	// type{values}: composite literal
	c := []int{4, 5, 6, 7, 42}
	fmt.Println("\nslice c:", c)

	// printSlice(c)
	// slicingTheSlice(c)
	// appendAndDelete(c)
	// buildinPKGMake()
	// multiDimensionalSlice()
	printMap()
}

func printSlice(c []int) {
	fmt.Println("length of slice:", len(c))

	for i, v := range c {
		fmt.Println(i, v)
	}
}

func slicingTheSlice(c []int) {
	fmt.Println("\n\nslicingTheSlice:")
	fmt.Println(c[1:])
	fmt.Println(c[1:3])
}

func appendAndDelete(c []int) {
	fmt.Println("\n\nappendAndDelete:")
	c = append(c, 567, 568)
	fmt.Println("append:", c)

	pc := []int{4445, 4446, 4447}
	c = append(c, pc...)
	fmt.Println("append another slice:", c)

	c = append(c[:4], c[7:]...)
	fmt.Println("delete:", c)
}

func buildinPKGMake() {
	fmt.Println("\n\nbuildinMake:")
	m := make([]int, 2, 3)
	fmt.Println("slice:", m)
	fmt.Println("slice of length:", len(m))
	fmt.Println("underlying array of size:", cap(m))

	fmt.Println()

	m = append(m, 1)
	fmt.Println("append 1:", m)
	fmt.Println("slice of length:", len(m))
	fmt.Println("underlying array of size:", cap(m))

	fmt.Println()

	// Optimized memory management
	// Appending to a slice when the underlying array has reached capacity requires a new array be instantiated and all values be copied into it.
	m = append(m, 2)
	fmt.Println("append 2:", m)
	fmt.Println("slice of length:", len(m))
	fmt.Println("underlying array of size:", cap(m))
}

func multiDimensionalSlice() {
	fmt.Println("\n\nmultiDimensionalSlice:")
	jb := []string{"James", "Bomb", "Chicky", "Nugget"}
	fmt.Println(jb)

	lf := []string{"Larry", "Fairy", "Hotdog", "Popcorn"}
	fmt.Println(lf)

	mds := [][]string{jb, lf}
	fmt.Println(mds)
}

func printMap() {
	fmt.Println("\n\nprintMap:")
	m := map[string]int{
		"James Bomb":  18,
		"Larry Fairy": 81,
	}

	fmt.Println(m)
	fmt.Println("\nm[\"James Bomb\"]:", m["James Bomb"])
	fmt.Println("m[\"James\"]:", m["James"])

	v, ok := m["James"]
	fmt.Println("\nv, ok := m[\"James\"]:\tv:", v, "\tok:", ok)

	if v, ok := m["Larry Fairy"]; ok {
		fmt.Println("if m[\"Larry Fairy\"] print key value:", v)
	}

	fmt.Println("\nAdd new key:")
	m["New Key"] = 1
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("\nDelete key:")
	delete(m, "New Key")
	fmt.Println(m)
}

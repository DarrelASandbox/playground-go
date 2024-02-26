// https://go.dev/blog/slices-intro

package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	// t.Run("collection of 5 numbers", func(t *testing.T) {
	// 	numbers := []int{1, 2, 3, 4, 5}

	// 	got := Sum(numbers)
	// 	want := 15

	// 	if got != want {
	// 		t.Errorf("got %d want %d given, %v", got, want, numbers)
	// 	}
	// })

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	/*
		  Go does not let you use equality operators with slices.
			It's important to note that reflect.DeepEqual is not "type safe" -
			the code will compile even if you did something a bit silly.
	*/

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

// Change `SumAll` to `SumAllTails`
func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}

// Example of slicing an array and how changing the slice affects the original array;
// but a "copy" of the slice will not affect the original array.
func TestSliceCopy(t *testing.T) {
	gotX, gotY, gotZ := SliceCopy()
	wantX := []string{"Лайка", "Belka", "Стрелка"}
	wantY := []string{"Лайка", "Belka", "Стрелка"}
	wantZ := []string{"Лайка", "Белка", "Стрелка"}

	if !reflect.DeepEqual(gotX, wantX) {
		t.Errorf("gotX %v want %v", gotX, wantX)
	}
	if !reflect.DeepEqual(gotY, wantY) {
		t.Errorf("gotY %v want %v", gotY, wantY)
	}
	if !reflect.DeepEqual(gotZ, wantZ) {
		t.Errorf("gotZ %v want %v", gotZ, wantZ)
	}
}

func TestAnotherSliceCopy(t *testing.T) {
	// Call the function to get the copied slice
	got := AnotherSliceCopy()

	// Check if the length of the returned slice is 2
	if len(got) != 2 {
		t.Errorf("Expected slice length of 2, got %d", len(got))
	}

	// Since the original slice `a` was made with make([]int, 1e6),
	// and no values were assigned, it should contain zeros.
	// Therefore, the copied slice should also contain zeros.
	want := []int{0, 0}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected slice values to be %v, got %v", want, got)
	}

	// Modify the original slice `a` to ensure it does not affect the copied slice.
	// This step assumes you have access to modify `a` from the test,
	// which in this case, you don't. This is more of a conceptual check.
	// If `a` was accessible: a[0] = 10

	// Assuming you could modify `a`, you would then check that `copiedSlice` is unaffected.
	// This part of the test is more theoretical in this context, as `a` is not accessible here.
	// However, it's an important aspect to consider when writing tests for real applications
	// where the original data source might be modified after a copy operation.
}

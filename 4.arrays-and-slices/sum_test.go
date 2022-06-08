package arraysandslices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

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
	t.Run("collection of two slices", func(t *testing.T) {
		slice1 := []int{1,2}
		slice2 := []int{0,9}

		got := SumAll(slice1, slice2)
		want := []int{3, 9} 

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()

		// reflect.DeepEqual and why it's useful but can reduce the type-safety of your code
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func (t *testing.T)  {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func (t *testing.T)  {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
	
}

func TestSlices(t *testing.T) {
	numbers := []int{1, 3, 6, 8, 9, 20 ,11}

	fmt.Println(numbers[1:]) // tail: 3 6 8 9 20 11
	fmt.Println(numbers[:1]) // head: 1
	fmt.Println(numbers[:])  // all:  1 3 6 8 9 20 11
}
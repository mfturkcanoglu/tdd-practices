package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	// t.Run("run collection of sized array", func(t *testing.T) {
	// 	// array numbers := [5]int{1, 2, 3, 4, 5}
	// 	numbers := []int{1, 2, 3, 4, 5} // slice

	// 	got := Sum(numbers)
	// 	want := 15

	// 	if got != want {
	// 		t.Errorf("got %d but expected %d, given %v", got, want, numbers)
	// 	}
	// })

	t.Run("run collection of unsized slice", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d but expected %d, given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{4, 5})
	want := []int{3, 9}

	checkSums(t, got, want)
}

func TestSumAllTails(t *testing.T) {
	t.Run("make sum of sum slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{4, 5})
		want := []int{2, 5}

		checkSums(t, got, want)
	})

	t.Run("make sum of empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})

}

func checkSums(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %v, but got %v", want, got)
	}
}

package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("slice of integers", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6
		if got != want {
			t.Errorf("got %d, want %d, given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sum a single slice of integers", func(t *testing.T) {
		slice := []int{1, 2, 3}

		got := SumAll(slice)
		want := []int{6}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v, given %v", got, want, slice)
		}
	})
	t.Run("sum of several slices of integers", func(t *testing.T) {
		slice1 := []int{1, 2}
		slice2 := []int{3, 9}
		given := [][]int{slice1, slice2}

		got := SumAll(slice1, slice2)
		want := []int{3, 12}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v, given %v", got, want, given)
		}
 	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}
	t.Run("tail sum several slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{3, 9, 0})
		want := []int{2, 9}
		checkSums(t, got, want)
	})
	t.Run("tail sum empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}
		checkSums(t, got, want)
	})
}

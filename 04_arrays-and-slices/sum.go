package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// Sums all integers in each slice and returns the sum of each of the slices.
func SumAll(slicesToSum ...[]int) []int {
	var sums []int

	for _, numbers := range slicesToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

// Sums all elements of a slice except the integer in the 0th index position.
func SumAllTails(slicesToSum ...[]int) []int {
	var sums []int

	for _, numbers := range slicesToSum {
		if len(numbers) > 0 {
			tail := Sum(numbers[1:])
			sums = append(sums, tail)
		} else {
			sums = append(sums, 0)
		}
	}
	return sums
}

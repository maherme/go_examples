package main

// Sum returns the total of an array of numbers.
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll takes a varying number of slices, returning a new slice containing
// the totals of each slice passed in.
func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// SumAllTails calculates the totals of the "tails" of each slice.
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}

	return sums
}

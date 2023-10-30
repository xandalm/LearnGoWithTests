package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(collectionsToSum ...[]int) []int {
	var sums []int
	for _, numbers := range collectionsToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

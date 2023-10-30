package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(collectionsToSum ...[]int) []int {
	collectionsCount := len(collectionsToSum)
	sums := make([]int, collectionsCount)

	for i, numbers := range collectionsToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}

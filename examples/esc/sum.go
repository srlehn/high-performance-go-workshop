package main

import "fmt"

// Sum returns the sum of the numbers 1 to 100 tag::sum[]
func Sum() int {
	numbers := make([]int, 100)
	for i := range numbers {
		numbers[i] = i + 1
	}

	var sum int
	for _, i := range numbers {
		sum += i
	}
	return sum
}

func main() {
	answer := Sum()
	fmt.Println(answer)
}

// end::sum[]

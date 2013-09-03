package main

import (
	"fmt"
	"math"
)

func main() {
	i := 8
	for {
		n := GetTriangleNumber(i)
		if GetNumberOfDivisors(n) > 500 {
			fmt.Print(n)
			break
		}
		i++
	}
}

func GetTriangleNumber(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum = sum + i
	}

	return sum
}

func GetNumberOfDivisors(n int) int {
	count := 2

	//No need to check all the numbers. If we found one divisor it means, that we can figure out second one.
	//It is needed to check all numbers up to sqrt(n)
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			count = count + 2
		}
	}

	return count
}

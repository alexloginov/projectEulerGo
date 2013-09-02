package main

import (
	"fmt"
)

func main() {
	num := 600851475143
	median := (num / 2) - 1

	for i := 2; i <= median; i++ {
		if num%i == 0 {
			if IsPrime(num / i) {
				fmt.Print(num / i)
				break
			}
		}
	}

}

func IsPrime(n int) bool {
	if n%2 == 0 {
		return false
	}
	for i := 3; i <= n/2; i = i + 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

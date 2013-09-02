package main

import (
	"fmt"
)


func main() {
	sum := 2
	
	for i := 3; i < 2000000; i = i + 2 {
		if IsPrime(i) {
			sum = sum + i
			fmt.Printf("%v \n",i);
		}
	}
	fmt.Print(sum)
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

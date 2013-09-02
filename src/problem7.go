package main 

import (
	"fmt"
)

func main() {
	count:=6
	i:=13
	for{
		i=i+2
		if IsPrime(i){
			count++
		}
		if count==10001{
			fmt.Print(i)
			break
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

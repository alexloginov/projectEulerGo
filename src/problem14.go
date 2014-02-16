package main

import (
	"fmt"
)

func main(){
	max:=1
	maxCount:=1

	for i:=2; i<1000000; i++{
		count := getChain(i)
		if count>maxCount{
			max = i
			maxCount = count
		}
	}

	fmt.Printf("Result: %d with chain %d",max,maxCount)
}

func getChain(n int) int {
	if (n==1){
		return 1
	}
	count:=1
	for{
		n = next(n)
		count++
		if (n==1){
			return count
		}
	}
}

func next(n int) int {
	if n%2==0{
		return n/2
	} else {
		return 3*n+1
	}
}

package main

import (
	"fmt"
)

func main() {
	sum:=0
	for i:=1; i<10000; i++{
		num1 := getSumOfDividers(i)
		num2 := getSumOfDividers(num1)
		if num2 == i && i<num1{
			sum = sum + i + num1
			fmt.Printf("Found pair: (%d)=%d and (%d)=%d\n ",i,num1,num1,num2)
		}
	}
	fmt.Printf("Sum is: %d",sum)
}

func getSumOfDividers(n int) int{
	sum:=0
	for i:=n/2; i>=1; i--{
		if(n%i==0){
			sum = sum+i
		}
	}
	return sum
}

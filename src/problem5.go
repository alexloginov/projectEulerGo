package main 

import (
	"fmt"
)

func main() {
	i := 20
	for{
		i=i+10;
		if check(i){
			fmt.Print(i) 
			break
		}	
	}
}

func check(n int) bool{
	for i:=2; i<=20; i++{
		if (n%i!=0){
			return false
		}
	}
	return true
}


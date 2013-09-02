package main 

import (
	"fmt"
)

func main() {
	sumOfSquares:=0
	squareOfSum:=0
	
	for i:=1; i<=100; i++{
		sumOfSquares=sumOfSquares+(i*i)
		squareOfSum=squareOfSum+i
	}
	squareOfSum=squareOfSum*squareOfSum
	
	fmt.Print(squareOfSum-sumOfSquares)
}


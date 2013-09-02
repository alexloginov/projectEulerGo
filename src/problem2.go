package main

import (
	"fmt"
)

func main() {
	sum := 0
	prePrev := 1
	prev := 1
	for {
		current := prePrev + prev
		if current < 4000000 {
			if current%2 == 0 {
				sum = sum + current
			}
		} else {
			break
		}
		prePrev = prev
		prev = current
	}
	
	fmt.Print(sum)

}

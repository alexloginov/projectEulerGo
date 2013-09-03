package main

import (
	"fmt"
)

func main() {
	for c := 999; c >= 0; c-- {
		for b := (1000 - c); b > 0; b-- {
			a := 1000 - c - b
			if (a*a + b*b) == c*c{
				fmt.Printf("%v \n", a*b*c)
			}
		}
	}
}

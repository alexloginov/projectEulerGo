package main

import (
	"fmt"
	"math/big"
)

func main() {
	n1 := big.NewInt(1)
	n2 := big.NewInt(1)
	count := 2
	for{
		temp := big.NewInt(1)
		temp.Set(n2)
		n2.Add(n2,n1)
		n1 = temp
		count++

		if len(n2.String())==1000{
			break
		}
	}

	fmt.Print(count)
}



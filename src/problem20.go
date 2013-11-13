package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	fact := big.NewInt(1)
	n := 100
	for i := 1; i <= n; i++ {
		fact.Mul(fact, big.NewInt(int64(i)))
	}
	sum := 0
	for _, r := range fact.String() {
		n, _ := strconv.Atoi(string(r))
		sum = sum + n
	}
	fmt.Print(sum)
}

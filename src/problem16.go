package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	num := big.NewInt(1)
	n := 1000
	for i := 1; i <= n; i++ {
		num.Mul(num, big.NewInt(int64(2)))
	}
	sum := 0
	for _, r := range num.String() {
		n, _ := strconv.Atoi(string(r))
		sum = sum + n
	}
	fmt.Print(sum)
}

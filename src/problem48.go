package main

import(
	"fmt"
	"math/big"
)

func main(){
	sum := big.NewInt(0)
	for i:=1; i<=1000; i++{
		sum.Add(sum,exp(i))
	}
	fmt.Print(sum)
}

func exp(n int) *big.Int{
	num := big.NewInt(int64(n))
	exp := big.NewInt(int64(n))
	for i:=1; i<n; i++{
		exp.Mul(exp,num)
	}
	return exp
}

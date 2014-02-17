package main

import(
	"fmt"
)

func main(){
	a:=1000;
	b:=1000;

	max :=0
	max_a:=0
	max_b:=0
	for i:=-a; i<=a; i++{
		for j:=-b; j<=b; j++{
			val := getChainCount(i,j)
			if val > max{
				max = val
				max_a = i
				max_b = j
			}
		}
	}
	fmt.Printf("%d and %d = %d chain ", max_a, max_b, max)
	fmt.Print(max_a*max_b)

}

func isPrime(n int) bool {
	if (n<0) {
		n = -n
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i <= n/2; i = i + 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func getChainCount(a,b int) int{
	n:=0
	for{
		if isPrime(n*n+a*n+b)==false{
			return n
		}
		n++
	}
}

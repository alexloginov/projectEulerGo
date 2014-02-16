package main 

import (
	"fmt"
	"strconv"
)

func main() {
	max:=0
	for i:=999; i>=100; i--{
		for j:=999; j>=100; j--{
			val := i*j
			if isPalindrome(strconv.Itoa(i*j)) && max<val{
				max = val
			}
		}
	}

	fmt.Printf("Result: %d",max)
}

func isPalindrome(p string) bool{
	s := []byte(p)
	j := len(s)-1
	for i:=0;i<len(s)/2;i++{
		if s[i] != s[j]{
			return false
		}
		j--
	}
	return true
}

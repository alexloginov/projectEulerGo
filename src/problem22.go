package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sort"
)

func main() {

	buf,err := ioutil.ReadFile("res/problem22.txt")
	if err!=nil{
		return
	}
	//Prepare array of names
	val := string(buf)
	val = strings.Replace(val,`"`,"",-1)
	arr := strings.Split(val,",")
	sort.Strings(arr)

	sum:=0

	for pos,name := range arr{

		nameValue := 0
		for i:=0;i<len(name);i++{
			nameValue=nameValue+getValue(string(name[i]))
		}
		sum = sum + (pos+1) * nameValue
	}

	fmt.Print(sum)
}

func getValue(s string) int{
	//Prepare map of char values
	//Very slow and stupid way, could be replaced by map with actual values or use ASCII code with some offset
	arr := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	for n,c := range arr {
		if c==s{
			return n+1
		}
	}
	return 0
}


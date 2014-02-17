package main

import(
	"fmt"
	"strconv"
)

func main(){
	sum := 0
	for i:=2; i<=999999; i++{ //"magic" number selected as upper limit by test
		if couldBeWritten(i){
			sum = sum + i
			fmt.Printf("%d \n",i)
		}
	}
	fmt.Printf("------")
	fmt.Print(sum)
}

func couldBeWritten(a int) bool{
	s := strconv.Itoa(a)
	sum:=0
	for _,n := range s{
		num,_ := strconv.ParseInt(string(n),10,0)
		sum = sum+pow(int(num),5)
	}

	if sum==a{
		return true
	}
	return false
}

func pow(num int, power int)  int{
	result := num;
	for i := 2; i <= power; i++ {
		result = result*num
	}
	return result
}

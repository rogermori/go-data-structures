package main

import (
	"fmt"
)

func main() {
   sum := addIntegers(1,2,3)
   fmt.Println("sum",sum)
   sum = addNumbers(1,2,3)
   fmt.Println("sum",sum)
   sum = addNumbers([]int{1,2,3})
   fmt.Println("sum",sum)
   sum = addNumbers(1, []int{2,3})
   fmt.Println("sum",sum)
   
}

func addIntegers(x ...int) int {
	fmt.Printf("x = %v of type %T\n",x,x)
	sum := 0
	for _, n := range x {
		sum += n
	}
	return sum
}

func addNumbers(x ...interface{}) int {
	fmt.Println(x)
	sum := 0
	for _, n := range x {
		switch n.(type) {
		case int:
			sum += n.(int)
		case []int :	
		    for _, v := range n.([]int) {
				sum += v
			}
		}
	}
	return sum
}




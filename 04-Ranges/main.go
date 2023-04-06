package main

import "fmt"

func main() {
	numbers := []int{33,34,35,36}
	
	for index, value := range numbers {
		fmt.Printf("numbers[%v] = %v\n",index,value)
	}
	
}
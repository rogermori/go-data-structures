package main

import "fmt"

func main() {
	type Person struct {
		Name string
		Grades []int
	}
	
	roger := Person {"Roger",[]int{98,99,100}}
	
	fmt.Println("Name: ",roger.Name)
	for index, value := range roger.Grades {
		fmt.Printf("grades[%d] = %d\n",index,value)
	}
	
	
}
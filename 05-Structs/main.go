package main

import "fmt"

func main() {
	type Person struct {
		Name string
		Age int
		Student bool
	}
	
	var tony  = Person{
		Name:    "Roger",
		Age:     50,
		Student: true,
	}
	
	peter := Person{"Tony",60,false}
	
	fmt.Println(tony)
	fmt.Println(peter)
	
}
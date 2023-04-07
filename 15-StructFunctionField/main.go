package main

import (
	"fmt"
)

type Person struct {
	Name string
	BirthYear  int
	Age func(int, int) int
}


func ComputeAge(currentYear int, yearBorn int) int {
	return currentYear - yearBorn	 
}

func main() {
     roger := Person{
		 Name:    "Roger",
		 BirthYear: 1980,
		 Age:     ComputeAge,
	 }
	 
	 fmt.Println(roger.Name,"is",roger.Age(2021,roger.BirthYear),"years old")
}
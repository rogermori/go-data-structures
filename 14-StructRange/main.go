package main

import (
	"github.com/fatih/structs"
	"fmt"
)

type Person struct {
	Name string
	Age  int
	Student bool
}



func main() {
     roger := Person{
		 Name:    "Roger",
		 Age:     20,
		 Student: true,
	 }
	 fields := structs.Fields(roger)
	 for _, field := range fields {
		 fmt.Printf("%v = %v (%T)\t", field.Name(), field.Value(),field.Value())	 
	}
}
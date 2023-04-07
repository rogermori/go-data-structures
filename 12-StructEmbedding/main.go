package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Student struct {
	Person
	Id string
}


func main() {
	roger := Student{Person{"Roger",19},"A123"}
	fmt.Println(roger)
	fmt.Println(roger.Name)
}
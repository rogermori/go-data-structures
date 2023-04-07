package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	var mary Person
	mary.Name = "Maria"

	tony := Person{"Tony de Araujo", 99}

	tonya := Person{Name: "Tony de Araujo", Age: 99}

	andre := &Person{Name: "Andre", Age: 33}

	cristina := new(Person)

	fmt.Println(mary)     // {Maria 0}
	fmt.Println(tony)     // {Tony de Araujo 99}
	fmt.Println(tonya)    // {Tony de Araujo 99}
	fmt.Println(andre)    // &{Andre 33}
	fmt.Println(cristina) // &{ 0}

}

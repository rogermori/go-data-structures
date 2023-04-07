package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Student struct {
	Person
	Name string
	Id string
}

func (x *Person) updateAge(newAge int) {
	x.Age = newAge
}

func (x *Student) updateAge(newAge int) {
	x.Age = newAge + 1
}


func main() {
	roger := Student{Person{Age: 19},"Ivan","A123"}
	fmt.Println(roger, roger.Name)
	roger.updateAge(20)
	fmt.Println(roger, roger.Name)
	
	
}
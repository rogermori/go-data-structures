package main

import "fmt"

func main() {
	type Person struct {
		Name string
		Age  int
	}
	
	roger := &Person{"Roger",19}
	var joe = new(Person)
	joe.Name = "Joe"
	joe.Age = 16
	
	fmt.Printf("%v %T\n",*roger,roger)
	fmt.Printf("%v %T",*joe,joe)
	
}
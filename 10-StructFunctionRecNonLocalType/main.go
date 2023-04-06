package main

import "fmt"

type String string

func (x *String) add(word String) {
	*x+= word
}

func main() {
	var hello String = "bana"
	hello.add("na")
	fmt.Println(hello)
	
	
}
package main

import (
	"fmt"
)

func saySomething(message interface{}) {
	fmt.Printf("%v \t %T\n",message,message)
}

func main() {
	saySomething("banana")
	saySomething(55)
	saySomething(true)
}

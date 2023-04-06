package main

import "fmt"

func main() {
	
	type NewStringType string
	type NewStringAlias = string
	
	var myName NewStringType = "Roger Mori"
	var myDogName NewStringAlias = "Pepito"
	
	fmt.Printf("%v is of type %T\n",myName,myName)
	fmt.Printf("%v is of type %T\n",myDogName,myDogName)
	
}
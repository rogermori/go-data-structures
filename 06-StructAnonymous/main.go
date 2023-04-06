package main

import "fmt"

func main() {
	roger := struct {
		name string
		age int
	} {"Roger",19}
	fmt.Println(roger)
}
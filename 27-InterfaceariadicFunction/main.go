package main

import (
	"fmt"
)

func main() {
	var m map[string]string
	var n map[string]bool
	testMap(n)
	testMap(m)
}

func testMap(x interface{}) {
	switch x.(type) {
	   case map[string]bool:
		   fmt.Println(x)
	default :
		fmt.Println("this is not a map[string]bool")
	}
}

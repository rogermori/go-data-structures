package main

import (
	"errors"
	"fmt"
)

func add(a,b interface{}) (int, error) {
	b1,ok2 := b.(int)
	if   a1,ok1 := a.(int); ok1 == true && ok2 == true {
		return a1 + b1, nil
	}
	return 0 ,errors.New("Invalid datatypes")
	
}

func main() {
	i, _ := add(1,2)
	fmt.Println(i)
	_, err := add("a","b")
	fmt.Println(err)
	
}

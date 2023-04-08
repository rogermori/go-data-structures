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
	i, err := add(1,2)
	if err == nil {
		fmt.Println(i)
	}
	i, err = add("a","b")
	if err != nil {
		fmt.Println(err)
	}
	
}

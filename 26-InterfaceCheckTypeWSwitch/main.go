package main

import (
	"errors"
	"fmt"
)

func inc(a interface{}) interface{} {
	switch t := a.(type) {
	case int32:
		return t +1
	case int64:
		return t +1
	case int8:
		return t +1
	case byte:
		return t +1
	case int:
		return t +1
	case uint16:
		return t +1
	case uint:
		return t +1
	case float32:
		return t +1
	case float64:
		return t +1
	default:
		fmt.Printf("%T\n",t)
		
	}
    return errors.New("invalid type")	
	
}

func main() {
	 fmt.Println(inc(1))
	 fmt.Println(inc(1.3))
	 fmt.Println(inc("a"))
}

package main

import (
	"errors"
	"fmt"
)

func inc(a interface{}) interface{} {
	switch t := a.(type) {
	case int32:
	case int64:	
	case int8:
	case byte:
	case int:
	case uint16:	
	case uint:
	case float32:	
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

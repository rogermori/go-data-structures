package main

import (
	"fmt"
	

)

var ages = make(map[string]int,100)

func main() {
	if len(ages) == 0 {
		fmt.Println("Map is empty")
	}
	ages["Jose"] = 42
	ages["Mary"] =  28
	ages["Peter"] =  55
	ages["Antonio" ] =  39
	
	fmt.Println(ages)
	
	keepValue := ages["Jose"]
	ages["Joseph"] = keepValue
	delete(ages,"Jose")
	
	fmt.Println(ages)
	 
}

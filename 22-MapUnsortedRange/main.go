package main

import (
	"fmt"
	"sort"
)

var ages = map[string]int {
	"Jose":42,
	"Mary": 28,
	"Peter": 55,
	"Antonio" : 39,
}

func main() {
     fmt.Printf("Sorted %v\n\n",ages)
	 
	 fmt.Println("Unsorted looping")
	 for key, value := range ages {
		 fmt.Printf("%v = %v\n",key,value)
	 }
	 sortedLoop()
	 
}

func sortedLoop(){
	keys := make([]string,len(ages))
	for key, _ := range ages {
		keys = append(keys,key)
	}
	sort.Strings(keys)
	fmt.Println("Sorted looping")
	for _, key := range keys {
		fmt.Printf("%v = %v\n",key,ages[key])
	}
	
}
package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	id int
}



func main() {
   var roger = Person{Name: "Roger"}
   var joe  = Person{Name: "Roger",id:444}
   var phil Person;
   if roger.id == 0 {
	  roger.id = 444   
   }
   fmt.Println(roger,joe)
   if roger == joe {   // compilation time
	   fmt.Println("Structs equate each other")
   }
   
   if (Person{}) == phil {
	   fmt.Println("phil is empty at compilation time")
	}
  
	// phil = loadFromDB()
    if (reflect.DeepEqual(Person{},phil)) { // run time
		fmt.Println("phil is empty at run time")
	}
   
   

}
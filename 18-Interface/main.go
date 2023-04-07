package main
import "fmt"
type reader interface {
	printMe(p interface{}) interface{}
}
func printVehicle(x reader, p interface{}) {
	x.printMe(p)
}
type car struct {
	make  string
	model string
}
type truck struct {
	make  string
	model string
}
func (a car)  printMe(p interface{}) interface{} {
	fmt.Println("This is car")
	fmt.Printf("Make: %v \t Model: %v ID: %v\n\n", a.make, a.model,p)
	return p
} 

func (a truck)  printMe(p interface{}) interface{} {
	fmt.Println("This is truck")
	fmt.Printf("Make: %v \t Model: %v ID: %v\n\n", a.make, a.model,p)
	return p
}


func main() {
	ford := car{make: "Ford", model: "Capri"}
	volvo := truck{"Volvo", "VHD"}
	bmw := car{"BMW","Z3"}
	subaru := car{"Subaru","Outback"}
	
	listOfVehicles := []reader{ford,subaru,bmw,volvo} 
	
	for index, vehicle := range listOfVehicles {
		printVehicle(vehicle, index+1)
	}
	
}
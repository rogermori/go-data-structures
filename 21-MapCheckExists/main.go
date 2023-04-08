package main
import "fmt"

var stockList = make(map[string]string)

func main() {
     
	 stockList["OPESY"] = "Opera Software"
	 stockList["SINA"]  = "SIna Corp"
	 
	 value, ok := stockList["OPESY"]
	 fmt.Printf("Value: %s exists?: %t\n",value,ok)
	 value, ok = stockList["XOPESY"]
	 fmt.Printf("Value: %s exists?: %t\n",value,ok)
	 
	
}
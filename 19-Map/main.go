package main
import "fmt"


func main() {
     stockList := make(map[string]string)
	 stockList["OPESY"] = "Opera Software"
	 stockList["SINA"]  = "SIna Corp"
	 
	 shoppingList := map[string]bool{
		  "milk": false,
		  "bananas": false,
		  "apples": true,
	 }
	 shoppingList["chocolate"] = true
	 shoppingList["apples"] = false
	 delete(shoppingList,"milk")
	 
	 fmt.Println(stockList)
	 fmt.Println(shoppingList)
	 
	
}
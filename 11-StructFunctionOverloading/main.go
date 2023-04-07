package main

import "fmt"

type num int
type str string

func (x *num) doubleMe() {
	*x *= 2
}
func (x *str) doubleMe() {
	*x += *x
}

func main() {
	var myNumber num = 2
	var myStr str = "2"
	myNumber.doubleMe()
	myStr.doubleMe()
	fmt.Println(myNumber)
	fmt.Println(myStr)
}
package main

import "fmt"

type bankAccount struct {
	name string
	balance int
}
func (b *bankAccount) print() {
	fmt.Println(b)
}
func (b *bankAccount) deposit(amount int) {
	b.balance += amount
}

func main() {
	
	checking  := bankAccount{"Checking",1000}
	checking.print()
	checking.deposit(100)
	checking.print()
	
	saving  := &bankAccount{"Saving",1000}
	saving.print()
	saving.deposit(100)
	saving.print()
		
}
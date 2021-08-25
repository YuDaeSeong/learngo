package main

import (
	"fmt"

	"github.com/learngo/banking"
)

// main - nomadcoders example
func main() {
	account := banking.NewAccount("nicol")

	account.Deposit(10)

	fmt.Println(account.Balance())

	if err := account.Withdraw(20); err != nil {
		fmt.Println(err)
	}

	fmt.Println(account.Balance())
}

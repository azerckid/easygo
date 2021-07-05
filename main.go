package main

import (
	"fmt"

	"github.com/azerckid/easygo/accounts"
)

func main() {
	account := accounts.NewAccount("tessa")
	account.Deposit(10)
	balance := account.Balance()
	fmt.Println(account, balance)
}

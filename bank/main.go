package main

import (
	"fmt"

	"github.com/rnjshippo/bank/mydict"
)

// "github.com/rnjshippo/bank/accounts"

func main() {

	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	dictionary.Search(baseWord)
	dictionary.Delete(baseWord)
	word, err := dictionary.Search(baseWord)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(word)
	}

	// dictionary := mydict.Dictionary{"first": "First word"}
	// definition, err := dictionary.Search("second")
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(definition)
	// }

	// account := accounts.NewAccount("kwon")
	// account.Deposit(10)
	// err := account.Withdraw(20)
	// if err != nil {
	// 	// log.Fatalln(err) //이렇게하면 에러 띄우고 프로그램 종료시킴
	// 	fmt.Println(err)
	// }
	// fmt.Println(account.Balance(), account.Owner())
	// fmt.Println(account.String())
}

package accounts

import (
	"errors"
	"fmt"
)

// Account struct 대문자로 시작해야 public 임, 소문자는 private
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw! you are poor")

// NewAccount create Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

//Deposit 아래처럼 하면 Account에 Deposit func가 추가됨. : receiver라고 부름
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Withdraw x amount from your ccount
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

//Balance of yout account :: 아래처럼 *를 안쓰면 값에 접근해서 수정불가
func (a Account) Balance() int {
	return a.balance
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

//Owner of account
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas ", a.Balance())
}

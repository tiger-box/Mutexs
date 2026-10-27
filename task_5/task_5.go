package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	balance int64
}

var errNegativeamount = errors.New("amount can not negative")
var errNegativeWithdraw = errors.New("It is not possible to withdraw funds.")
var errNegativAamountWithdraw = errors.New("negative amount")

func (account *BankAccount) Deposit(amount int64) error {

	if amount < 0 {
		return errNegativeamount
	}

	account.balance += amount

	return nil

}

func (account *BankAccount) Withdraw(amount int64) error {
	if account.balance < amount {
		return errNegativeWithdraw
	}

	if amount < 0 {
		return errNegativAamountWithdraw
	}

	account.balance -= amount

	return nil
}

func (account *BankAccount) Balance() int64 {

	if account.balance < 0 {
		return 0
	}

	return account.balance
}

func main() {

	bank := &BankAccount{}

	fmt.Println(bank.balance)
	fmt.Println(bank.Balance())

	bank.Deposit(1000)
	bank.Deposit(500)
	fmt.Println(bank.Balance())

	errDeposit := bank.Deposit(-100)
	errWithdrawAmount := bank.Withdraw(-100)
	errbalanceWithdraw := bank.Withdraw(2500)

	if errors.Is(errDeposit, errNegativeamount) {
		fmt.Println("нельзя внесси сумма отрецательного числа")
	}

	if errors.Is(errWithdrawAmount, errNegativAamountWithdraw) {
		fmt.Println("не должно быть отрецательно число в ()")
	}

	if errors.Is(errbalanceWithdraw, errNegativeWithdraw) {
		fmt.Println("недостаточно средств")
	}

}

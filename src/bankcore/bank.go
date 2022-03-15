package bank

import (
  "errors"
  "fmt"
)

type Customer struct {
  Name string
  Address string
  Phone string
}

type Account struct {
  Customer
  Number int32
  Balance float64
}


type Bank interface {
  Statement() string
}

func Statement(b Bank) string {
  return b.Statement()
}

func (a *Account) Deposit(amount float64) error {
  if amount <= 0 {
    return errors.New("the amount to deposit should be greater than zero")
  }

  a.Balance += amount
  return nil
}

func (a *Account) Withdraw(amount float64) error {
  if amount <= 0 {
    return errors.New("the amount to withdraw should be greater than zero")
  }

  if a.Balance < amount {
    return errors.New("the amount to withdraw should be greater than the account's balance")
  }

  a.Balance -= amount
  return nil
}

func (a *Account) Statement() string {
  return fmt.Sprintf("%v - %v - %v", a.Number, a.Name, a.Balance)
}

func (a *Account) Transfer(amount float64, dest *Account) error {
  if amount<= 0 {
    return errors.New("the amount to transfer shoulf be greater than zero")
  }

  if a.Balance < amount {
    return errors.New("amount to transfer should be greater than the account's balance")
  }

  a.Withdraw(amount)
  dest.Deposit(amount)
  return nil
}

func Hello() string {
  return "Hey! I'm working!"
}

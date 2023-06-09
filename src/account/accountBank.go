package account

import client "projects/goStudyng/POO/Bank/src/client"

type AccountCurrent struct {
	Title                       client.Customers
	NumberAgency, NumberAccount int
	Balance                     float64
}

func (c *AccountCurrent) WithDrawMoney(value float64) string {
	if value > 0 && value <= c.Balance {
		c.Balance -= value
		return "Successful withdrawal"
	} else {
		return "Insufficient Balance"
	}
}

func (c *AccountCurrent) DepositMoney(value float64) (string, float64) {
	if value > 0 {
		c.Balance += value
		return "\nDeposit made successfully Mr. " + c.Title.Name + " your current Balance is: ", c.Balance
	} else {
		return "\nDeposit has not been made Mr. " + c.Title.Name + " your current Balance is: ", c.Balance
	}
}

func (c *AccountCurrent) TransferAccount(value float64, targetAccount *AccountCurrent) string {
	if value > 0 && value < c.Balance {
		c.Balance -= value
		targetAccount.DepositMoney(value)
		return "\nMr. " + c.Title.Name + " Transfer successfully for " + targetAccount.Title.Name
	} else {
		return "\nMr. " + c.Title.Name + " Transfer unsuccessfuly for " + targetAccount.Title.Name
	}
}

func (c *AccountCurrent) GetBalanceUser() float64 {
	return c.Balance
}

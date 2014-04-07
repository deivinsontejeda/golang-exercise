package funding

type Fund struct {
	// balance is unexported (private), because it's lowercase
	balance int
}

func NewFund(initBalance int) *Fund {
	return &Fund{
		balance: initBalance,
	}
}

// Methods start with a *receiver*, in this case a Fund pointer
func (f *Fund) Balance() int {
	return f.balance
}

func (f *Fund) Withdraw(amount int) {
	f.balance -= amount
}

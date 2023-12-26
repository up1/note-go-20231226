package demo

type DepositRequest struct {
	AccountNo string
	Amount    int
}

type Bank struct {
}

func (b Bank) Deposit(req DepositRequest) error {
	return nil
}

func (b Bank) Withdraw() error {
	return nil
}

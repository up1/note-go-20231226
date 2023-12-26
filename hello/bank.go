package demo

type Bank struct {
}

func (b Bank) Deposit(req DepositRequest) error {
	return nil
}

func (b Bank) Withdraw() error {
	return nil
}

package demo

type Bank struct {
	Message string
}

func NewBank(message string) Bank {
	return Bank{
		Message: message,
	}
}

func (b *Bank) Deposit(req DepositRequest) error {
	b.Message = "Deposit"
	return nil
}

func (b Bank) Withdraw() error {
	return nil
}

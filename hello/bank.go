package demo

type DepositRequest struct {
	AccountNo string `json:"account_no"`
	Amount    int    `json:"amount"`
}

type DepositRequestV2 struct {
	DepositRequest
	Credit bool `json:"credit"`
}

type Bank struct {
}

func (b Bank) Deposit(req DepositRequest) error {
	return nil
}

func (b Bank) Withdraw() error {
	return nil
}

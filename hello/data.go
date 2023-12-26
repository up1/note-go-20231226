package demo

type DepositRequest struct {
	AccountNo string `json:"account_no"`
	Amount    int    `json:"amount"`
}

func (dr DepositRequest) String() string {
	return "AccountNo=" + dr.AccountNo + ", Amount=" + string(dr.Amount)
}

type DepositRequestV2 struct {
	DepositRequest
	Credit bool `json:"credit"`
}

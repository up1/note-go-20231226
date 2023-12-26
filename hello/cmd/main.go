package main

import (
	"demo"
	"fmt"
)

func main() {
	b := demo.Bank{}
	req := demo.DepositRequest{
		AccountNo: "1234567890",
		Amount:    100,
	}
	b.Deposit(req)
	b.Withdraw()
	fmt.Println(req)

	r2 := demo.DepositRequestV2{
		Credit: true,
		DepositRequest: demo.DepositRequest{
			AccountNo: "1234567890",
			Amount:    100,
		},
	}
	println(r2.DepositRequest)
	println(r2.Amount)
	println(r2.Credit)

	r, err := demo.TryError()
	if err != nil {
		println(err.Error())
	}
	println(r)

	result := demo.SayHi()
	println(result)
}

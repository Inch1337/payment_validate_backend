package main

import (
	"fmt"
	"payment_backend/internal/factory"
	"payment_backend/internal/payments"
)

func ProcessPayment(
	p payments.Payment,
	vf *factory.ValidatorFactory,
	amount float64,
) error {

	validator := vf.Get(p)
	if validator == nil {
		return fmt.Errorf("no validator for payment type")
	}

	if err := validator.Validate(p); err != nil {
		return err
	}

	return p.Pay(amount)
}

func main() {
	vf := factory.NewValidatorFactory()

	paymentsList := []payments.Payment{
		&payments.PayPal{User: "John", Email: "john@example.com"},
		&payments.Card{CardNumber: "1234567812345678", DateOfExpiry: "12/26"},
		&payments.Crypto{Wallet: "0x1234567890abcdef1234567890abcdef12345678"},
	}

	for _, p := range paymentsList {
		if err := ProcessPayment(p, vf, 100); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Payment successful")
		}
	}
}

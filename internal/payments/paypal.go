package payments

import "fmt"

type PayPal struct {
	User  string
	Email string
}

func (p *PayPal) Pay(amount float64) error {
	fmt.Printf("Paid %.2f via PayPal (%s)\n", amount, p.Email)
	return nil
}
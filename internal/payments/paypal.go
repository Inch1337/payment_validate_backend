package payments

import "fmt"

type PayPal struct {
	User  string
	Email string
}

func (p *PayPal) Pay(amount int) error {
	fmt.Printf("Paid %d via PayPal (%s)\n", amount/100, p.Email)
	return nil
}

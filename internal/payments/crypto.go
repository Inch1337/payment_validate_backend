package payments

import "fmt"

type Crypto struct {
	Wallet string
}

func (c *Crypto) Pay(amount float64) error {
	fmt.Printf("Paid %.2f via Crypto wallet %s\n", amount, c.Wallet)
	return nil
}
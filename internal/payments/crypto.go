package payments

import "fmt"

type Crypto struct {
	Wallet string
}

func (c *Crypto) Pay(amount int) error {
	fmt.Printf("Paid %d via Crypto wallet %s\n", amount/100, c.Wallet)
	return nil
}

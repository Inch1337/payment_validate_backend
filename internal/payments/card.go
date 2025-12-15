package payments

import "fmt"

type Card struct {
	CardNumber   string
	DateOfExpiry string // MM/YY
}

func (c *Card) Pay(amount float64) error {
	fmt.Printf("Paid %.2f via Card %s\n", amount, c.CardNumber)
	return nil
}
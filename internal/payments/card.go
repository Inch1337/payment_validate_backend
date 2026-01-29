package payments

import "fmt"

type Card struct {
	CardNumber   string
	DateOfExpiry string // MM/YY
}

func (c *Card) Pay(amount int) error {
	fmt.Printf("Paid %d via Card %s\n", amount/100, c.CardNumber)
	return nil
}

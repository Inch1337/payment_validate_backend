package main

import (
	"errors"
	"fmt"
	"strings"
)

type Payment interface {
	Pay(amount float64) error
}

type PayPal struct {
	Email string
}

type Visa struct {
	CardNumber string
}

type Crypto struct {
	Wallet string
}

func (p *PayPal) Pay(amount float64) error {
	fmt.Printf("Paid: %.2f via PayPal to email %s\n", amount, p.Email)
	return nil
}

func (p *PayPal) Validate() error {
	if p.Email == "" {
		return errors.New("имя не может быть пустым")
	}
	if !strings.Contains(p.Email, "@") {
		return errors.New("не соответствует стандартам email")
	}
	return nil
}

func main() {
	pay := PayPal{
		Email: "inchgmail.com",
	}

	err := pay.Validate()
	if err != nil {
		fmt.Print(err)
	}
}

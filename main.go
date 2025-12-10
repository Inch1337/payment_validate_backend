package main

import (
	"errors"
	"fmt"
	"strings"
)

// интерфейсы
type Validateble interface {
	Validate() error
}

type Payment interface {
	Pay(amount float64) error
	Validateble
}

// cтруктуры
type PayPal struct {
	User  string
	Email string
}

type Visa struct {
	CardNumber   string
	DataOfExpiry string
}

type Crypto struct {
	Wallet string
}

// методы PayPal

func (p *PayPal) Pay(amount float64) error {
	fmt.Printf("Paid: %.2f via PayPal to email %s\n", amount, p.Email)
	return nil
}

func (p *PayPal) Validate() error {
	return ValidateEmail(p.Email)
}

// методы Visa

func (v *Visa) Pay(amount float64) error {
	fmt.Printf("Paid: %.2f via Visa to cardnumber %s\n", amount, v.CardNumber)
	return nil
}

func (v *Visa) Validate() error {
	return ValidateCardNumber(v)
}

// методы Crypto

func (c *Crypto) Pay(amount float64) error {
	fmt.Printf("Paid: %.2f via Crypto to wallet %s\n", amount, c.Wallet)
	return nil
}

func (c *Crypto) Validate() error {
	return ValidateWallet(c.Wallet)
}

// функции

func ValidateEmail(email string) error {
	if email == "" {
		return errors.New("emal пуст")
	}
	if !strings.Contains(email, "@") {
		return errors.New("неправильный формат emal")
	}
	return nil
}

func ValidateCardNumber(v *Visa) error {
	if len(v.CardNumber) != 16 {
		return errors.New("номер карты должен состоять из 16 цифр")
	}
	if !strings.Contains(v.DataOfExpiry, "26") {
		return errors.New("срок действия карты истек") // доделать с помощью библиотеки с датами
	}
	return nil
}

func ValidateWallet(wallet string) error {
	if len(wallet) < 10 {
		return errors.New("крипто-кошелёк слишком короткий")
	}
	return nil
}

func ProcessPayment(p Payment, amount float64) error {
	if err := p.Validate(); err != nil {
		return err
	}
	return p.Pay(amount)
}

func main() {
	payments := []Payment{
		&PayPal{Email: "inch@gmail.com"},
		&Visa{CardNumber: "1234567891234567", DataOfExpiry: "26"},
		&Crypto{Wallet: "0xABC1234567890"},
	}

	for _, p := range payments {
		if err := ProcessPayment(p, 100); err != nil {
			fmt.Println("Ошибка:", err)
		}
	}
}

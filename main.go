package main

// распределить на пакеты, переписать всё на английский и доделать остальные идеи

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

// интерфейсы

type Payment interface {
	Pay(amount float64) error
}

type Validator interface {
	Validate(p Payment) error
}

type CardValidator struct{}

func (c *CardValidator) Validate(p Payment) error {
	// сделать расшифровку карт, отдельно visa, mastercard etc
	card, ok := p.(*Card)
	if !ok {
		return errors.New("ожидалось Card, но получен другой тип платежа")
	}

	if len(card.CardNumber) != 16 {
		return errors.New("номер карты должен состоять из 16 цифр")
	}
	// в будущем использовать библиотеку time для правильной валидации даты карты
	cardDate, err := strconv.Atoi(card.DateOfExpiry)

	if err != nil {
		return errors.New("failed to convert string to integer")
	}

	if cardDate >= 26 {
		return errors.New("срок действия карты истек")
	}

	return nil
}

type PayPalValidator struct{}

func (pp *PayPalValidator) Validate(p Payment) error {
	paypal, ok := p.(*PayPal)

	if !ok {
		return errors.New("ожидалось PayPal, но получен другой тип платежа")
	}

	if paypal.User == "" {
		return errors.New("имя пользователя не может быть пустым")
	}
	// переместить в отдельный блок с валидацией email для удобства использования в других местах где есть email
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]{1,64}@[a-zA-Z0-9._%+\-]{1,255}\.[a-zA-Z]{2,}$`)

	if !re.MatchString(paypal.Email) {
		return errors.New("email не соответсвует стандартам")
	}

	if regexp.MustCompile(`\.\.`).MatchString(paypal.Email) {
		return errors.New("email содержит недопустимую последовательность `..`")
	}

	if regexp.MustCompile(`@\.`).MatchString(paypal.Email) {
		return errors.New("email содержит недопустимую последовательность `@.`")
	}

	return nil
}

type CryptoValidator struct{}

func (c *CryptoValidator) Validate(p Payment) error {
	// сделать разделение на криптокошельки btc, eth, ton etc
	crypto, ok := p.(*Crypto)

	if !ok {
		return errors.New("ожидалось Crypto, но получен другой тип платежа")
	}

	re := regexp.MustCompile(`^(0x[a-fA-F0-9]{40}|[13][a-km-zA-HJ-NP-Z1-9]{25,34}|bc1[ac-hj-np-z02-9]{11,71})$`)

	if !re.MatchString(crypto.Wallet) {
		return errors.New("ваш крипто кошелек не соответствует стандартам")
	}
	return nil

}

// cтруктуры
type PayPal struct {
	User  string
	Email string
}

type Card struct {
	CardNumber   string
	DateOfExpiry string
}

type Crypto struct {
	Wallet string
}

// методы PayPal

func (p *PayPal) Pay(amount float64) error {
	fmt.Printf("Paid: %.2f via PayPal to email %s\n", amount, p.Email)
	return nil
}

// методы Visa

func (v *Card) Pay(amount float64) error {
	fmt.Printf("Paid: %.2f via Card to cardnumber %s\n", amount, v.CardNumber)
	return nil
}

// методы Crypto

func (c *Crypto) Pay(amount float64) error {
	fmt.Printf("Paid: %.2f via Crypto to wallet %s\n", amount, c.Wallet)
	return nil
}

// функции

func ProcessPayment(p Payment, v Validator, amount float64) error {
	if err := v.Validate(p); err != nil {
		return err
	}
	return p.Pay(amount)
}

func main() {
	payments := []Payment{
		&PayPal{User: "John", Email: "john@example.com"},
		&Card{CardNumber: "1234567890123456", DateOfExpiry: "25"},
		&Crypto{Wallet: "0xABC1234567890"},
	}

	for _, p := range payments {
		var v Validator

		switch p.(type) {
		case *PayPal:
			v = &PayPalValidator{}
		case *Card:
			v = &CardValidator{}
		case *Crypto:
			v = &CryptoValidator{}
		default:
			fmt.Println("Неизвестный тип платежа")
			continue
		}

		if err := ProcessPayment(p, v, 100); err != nil {
			fmt.Println("Ошибка:", err)
		} else {
			fmt.Println("Платёж прошёл успешно")
		}
	}
}

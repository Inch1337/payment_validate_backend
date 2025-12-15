package validators

import (
	"errors"
	"payment_backend/internal/payments"
	"regexp"
)

type PayPalValidator struct{}

var emailRegex = regexp.MustCompile(
	`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`,
)

func (v *PayPalValidator) Validate(p payments.Payment) error {
	pp, ok := p.(*payments.PayPal)
	if !ok {
		return errors.New("invalid payment type for paypal validator")
	}

	if pp.User == "" {
		return errors.New("user is required")
	}

	if !emailRegex.MatchString(pp.Email) {
		return errors.New("invalid email format")
	}

	return nil
}

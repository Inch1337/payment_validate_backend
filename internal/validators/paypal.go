package validators

import (
	"errors"
	"payment_backend/internal/payments"
	"payment_backend/internal/validation"
)

type PayPalValidator struct{}

func (v *PayPalValidator) Validate(p payments.Payment) error {
	pp, ok := p.(*payments.PayPal)
	if !ok {
		return errors.New("invalid payment type for paypal validator")
	}

	if pp.User == "" {
		return errors.New("user is required")
	}

	if !validation.IsValidEmail(pp.Email) {
		return errors.New("invalid email format")
	}

	return nil
}

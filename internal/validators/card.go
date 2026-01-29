package validators

import (
	"errors"
	"payment_backend/internal/payments"
	"payment_backend/internal/validation"
)

type CardValidator struct{}

func (v *CardValidator) Validate(p payments.Payment) error {
	card, ok := p.(*payments.Card)
	if !ok {
		return errors.New("invalid payment type for card validator")
	}

	if !validation.IsValidCardNumber(card.CardNumber) {
		return errors.New("card number must be 16 digits")
	}

	if err := validation.IsExpired(card.DateOfExpiry); err != nil {
		return err
	}

	return nil
}

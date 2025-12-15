package validators

import (
	"errors"
	"payment_backend/internal/payments"
	"time"
)

type CardValidator struct{}

func (v *CardValidator) Validate(p payments.Payment) error {
	card, ok := p.(*payments.Card)
	if !ok {
		return errors.New("invalid payment type for card validator")
	}

	if len(card.CardNumber) != 16 {
		return errors.New("card number must be 16 digits")
	}

	expiry, err := time.Parse("01/06", card.DateOfExpiry)
	if err != nil {
		return errors.New("invalid card expiry format")
	}

	if expiry.Before(time.Now()) {
		return errors.New("card expired")
	}

	return nil
}

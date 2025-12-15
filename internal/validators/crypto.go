package validators

import (
	"errors"
	"payment_backend/internal/payments"
	"payment_backend/internal/validation"
)

type CryptoValidator struct{}

func (v *CryptoValidator) Validate(p payments.Payment) error {
	crypto, ok := p.(*payments.Crypto)
	if !ok {
		return errors.New("invalid payment type for crypto validator")
	}

	if !validation.IsValidCrypto(crypto.Wallet) {
		return errors.New("invalid crypto wallet")
	}

	return nil
}

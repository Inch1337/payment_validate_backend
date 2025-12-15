package validators

import (
	"errors"
	"payment_backend/internal/payments"
	"regexp"
)

type CryptoValidator struct{}

var walletRegex = regexp.MustCompile(
	`^(0x[a-fA-F0-9]{40}|bc1[a-z0-9]{25,39})$`,
)

func (v *CryptoValidator) Validate(p payments.Payment) error {
	crypto, ok := p.(*payments.Crypto)
	if !ok {
		return errors.New("invalid payment type for crypto validator")
	}

	if !walletRegex.MatchString(crypto.Wallet) {
		return errors.New("invalid crypto wallet")
	}

	return nil
}

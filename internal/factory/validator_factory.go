package factory

import (
	"payment_backend/internal/payments"
	"payment_backend/internal/validators"
	"reflect"
)

type ValidatorFactory struct {
	registry map[string]validators.Validator
}

func NewValidatorFactory() *ValidatorFactory {
	return &ValidatorFactory{
		registry: map[string]validators.Validator{
			"Card":   &validators.CardValidator{},
			"PayPal": &validators.PayPalValidator{},
			"Crypto": &validators.CryptoValidator{},
		},
	}
}

func (f *ValidatorFactory) Get(p payments.Payment) validators.Validator {
	typeName := reflect.TypeOf(p).Elem().Name()
	return f.registry[typeName]
}

package validators

import "payment_backend/internal/payments"

// add generics
type Validator interface {
	Validate(p payments.Payment) error
}

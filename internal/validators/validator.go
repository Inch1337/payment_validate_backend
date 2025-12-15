package validators

import "payment_backend/internal/payments"

type Validator interface {
	Validate(p payments.Payment) error
}

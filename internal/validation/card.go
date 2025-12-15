package validation

import (
	"time"
)

func IsValidCardNumber(number string) bool{
	return len(number) == 16 
}

func IsExpired(expiry string) bool {
	t, err := time.Parse("01/06", expiry)
	if err != nil {
		return true
	}

	return t.Before(time.Now())
}
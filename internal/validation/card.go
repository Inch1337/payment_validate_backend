package validation

import (
	"errors"
	"time"
)

func IsValidCardNumber(number string) bool{
	return len(number) == 16 
}

func IsExpired(expiry string) (error){
	ok, err := validateTime(expiry)
	if err != nil {
		return err
	}else if ok{
		return errors.New("card expired")
	}
	
	return nil
	// t, err := time.Parse("01/06", expiry)
	// if err != nil {
	// 	return t.Before(time.Now()), err
	// }

	// return t.Before(time.Now()), nil

}
func validateTime(expiry string) (bool, error) {
	t, err := time.Parse("01/06", expiry)
	if err != nil {
		return t.Before(time.Now()), err
	}

	return t.Before(time.Now()), nil
}
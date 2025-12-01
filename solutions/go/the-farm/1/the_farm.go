package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(f FodderCalculator, cows int) (float64, error) {
	fatteningFactor, ok1 := f.FatteningFactor()
	if ok1 != nil {
		return 0, ok1
	}
	fodderAmount, ok2 := f.FodderAmount(cows)
	if ok2 != nil {
		return 0, ok2
	}

	result := fatteningFactor * fodderAmount / float64(cows)
	return result, nil
}

func ValidateInputAndDivideFood(f FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(f, cows)
}

type InvalidCowsError struct {
	cows    int
	message string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.message)
}

func ValidateNumberOfCows(cows int) error {
	var message string
	switch {
	case cows < 0:
		message = "there are no negative cows"
	case cows == 0:
		message = "no cows don't need food"
	default:
		return nil
	}
	return &InvalidCowsError{
		cows:    cows,
		message: message,
	}
}

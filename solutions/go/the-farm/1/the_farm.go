package thefarm
import "errors"
import "fmt"



// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, numberOfCows int) (float64, error) {
	
	
	fodder, err := fc.FodderAmount(numberOfCows)

	if err != nil {
		return 0, err
	}

	factor, err := fc.FatteningFactor()

	if err != nil {
		return 0, err
	}

	

	return (fodder * factor) / float64(numberOfCows), nil

}
// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, numberOfCows int) (float64, error) {
	var ErrNotFound = errors.New("invalid number of cows")
	if numberOfCows <= 0 {
		return 0, ErrNotFound
	}

   return DivideFood(fc, numberOfCows)
}
// TODO: define the 'ValidateNumberOfCows' function

type InvalidCowsError struct {
	Number  int
	Message string
}

// Implement the `error` interface
func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.Number, e.Message)
}


func ValidateNumberOfCows(numberOfCows int) error {
	if numberOfCows < 0 {
		return &InvalidCowsError{
			Number:  numberOfCows,
			Message: "there are no negative cows",
		}
	}
	if numberOfCows == 0 {
		return &InvalidCowsError{
			Number:  numberOfCows,
			Message: "no cows don't need food",
		}
	}
	return nil
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.

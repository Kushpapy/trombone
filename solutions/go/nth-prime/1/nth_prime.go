package nthprime

import "errors"



// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
        err := errors.New("there is no zeroth prime")
        return 0, err
    }

    count := 0
	number := 1
      
    for count < n {
        number++

        isPrime := true

        for i := 2; i*i <= number; i++ {
            if number % i == 0{
                isPrime = false
                break
            }
        }

        if isPrime {
            count++
        }
 
    }


    return number, nil
}

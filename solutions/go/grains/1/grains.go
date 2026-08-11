package grains

import (
	"errors"
	"math"
)

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
        return 0, errors.New("invalid input: number must be >= 1 or < 64")
    }
    
	base := 2.0
    exponent := float64(number - 1)

    result := math.Pow(base, exponent)

    return uint64(result), nil
}

func Total() uint64 {
	n := 64 

    var result uint64
    
    for n >= 1 {
        val, err := Square(n)
		if err != nil {
            break
        }
        
        result += val

       n = n - 1
    }

    
    return result
}

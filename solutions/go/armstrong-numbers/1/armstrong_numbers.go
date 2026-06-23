package armstrongnumbers

import (
	"fmt"
	"math"
)

func IsNumber(n int) bool {
	count := 0
	str := fmt.Sprintf("%d",n)

    for _,v := range str {
        print(v)
        count += 1
    }

    num := 0
    for _,v := range str{
        pot := int(v - '0')
        num += int(math.Pow(float64(pot),float64(count)))
    }

    return num == n
}

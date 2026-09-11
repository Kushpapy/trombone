package complexnumbers

import "math"

// Define the Number type here.

type Number struct {
    real float64
    imaginary float64
   
}

func (n Number) Real() float64 {
	return n.real
}

func (n Number) Imaginary() float64 {
	return n.imaginary
}

func (n1 Number) Add(n2 Number) Number {
	return Number{n1.real + n2.real, n1.imaginary + n2.imaginary}
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number{n1.real - n2.real, n1.imaginary - n2.imaginary}
}

func (n1 Number) Multiply(n2 Number) Number {
	return Number {n1.real * n2.real - n1.imaginary * n2.imaginary,n1.real * n2.imaginary + n1.imaginary*n2.real}
    
}

func (n Number) Times(factor float64) Number {
	return Number {n.real * factor, n.imaginary * factor}
}

func (n1 Number) Divide(n2 Number) Number {
	d := n2.real * n2.real + n2.imaginary * n2.imaginary
    realPart := (n1.real*n2.real + n1.imaginary*n2.imaginary) / d
    imagPart := (n1.imaginary*n2.real - n1.real*n2.imaginary) / d
    return Number{realPart, imagPart}
}

func (n Number) Conjugate() Number {
	return Number{n.real, -n.imaginary}
}

func (n Number) Abs() float64 {
	return math.Hypot(n.real, n.imaginary)
}

func (n Number) Exp() Number {
	m := math.Exp(n.real)
    return Number{m * math.Cos(n.imaginary),m * math.Sin(n.imaginary)}
}

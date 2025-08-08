package sorting
import "fmt"
//import "math"
import "strconv"

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
    //x := float64(math.Round(f))
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	x := float64(nb.Number())

    return fmt.Sprintf("This is a box containing the number %.1f", x)
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	fn, ok := fnb.(FancyNumber)
	
    if !ok {
      return 0  
    }

    var number, _ = strconv.Atoi(fn.Value())
    if number == 0 {
        return 0
    }
        
    return number
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	i := ExtractFancyNumber(fnb)
    var str string
    if i == 0 {
        str = fmt.Sprintf("This is a fancy box containing the number %.1f", float64(i))
    }else{
        str = fmt.Sprintf("This is a fancy box containing the number %.1f", float64(i))
    }

    return str
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	switch v := i.(type) {
        case int:
        	return DescribeNumber(float64(v))
        case float64:
        	return DescribeNumber(v)
        case NumberBox:
            return DescribeNumberBox(v)
        case FancyNumberBox:
           return DescribeFancyNumberBox(v)
    }

    return "Return to sender"
}

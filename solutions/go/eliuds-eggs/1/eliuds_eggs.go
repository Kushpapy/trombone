package eliudseggs

func EggCount(displayValue int) int {
   count := 0

    for displayValue > 0 {
     test := displayValue % 2
        if test == 1{
            count++
        }
      displayValue  = displayValue / 2
    }

    return count
}

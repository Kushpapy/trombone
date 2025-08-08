package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	pr := (successRate/float64(100)) * float64(productionRate);

    return float64(pr);
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	//every 60 minutes -> CalculateWorkingCarsPerHour(productionRate, successRate)
	return int(CalculateWorkingCarsPerHour(productionRate, successRate)/float64(60))
    
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	st := int(carsCount/10);
    rem := int(carsCount % 10);

    total := (st * 95000) + (rem * 10000);

    return uint(total)
}

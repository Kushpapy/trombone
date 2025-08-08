package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    count := 0;
	for _,value := range birdsPerDay{
        count += value
    }

    return count
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	daysInWeek := (week - 1) * 7;
    lastDayOfWeek := 7 * week;
    counter := 0;

    for i := daysInWeek; i < lastDayOfWeek; i++{
        counter += birdsPerDay[i];
    }

    return counter
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
		for i := 0; i < len(birdsPerDay); i++ {
		if i%2 == 0 { 
			birdsPerDay[i] += 1
		}
	}


    return birdsPerDay
}

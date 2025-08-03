package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
    
    if time <= 0{
        return len(layers) * 2
    }

    return len(layers) * time
    
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64){
    noodles := 0
    sauce := 0.0
    for _,value := range layers{
        if value == "noodles"{
            noodles += 50
        }
        if value == "sauce"{
            sauce += 0.2
        }
    }
    return noodles,sauce
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string){
     if len(friendsList) == 0 || len(myList) == 0 {
        return // optionally handle invalid input
    }
    secret := friendsList[len(friendsList)-1]       // Get the last item from friend's list
    myList[len(myList)-1] = secret 
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, scaledQuantities int) [] float64 {
	var s []float64
	for _, value := range quantities {
		newValue := (value * float64(scaledQuantities)) / 2
		s = append(s, newValue)
	}
	return s
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.

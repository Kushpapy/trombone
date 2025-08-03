package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := make(map[string]int)
    units["quarter_of_a_dozen"] = 3
    units["half_of_a_dozen"] = 6
    units["dozen"] = 12
    units["small_gross"] = 120
    units["gross"] = 144
    units["great_gross"] = 1728

    return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := make(map[string]int);
	
    return bill;
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	quantity, ok := units[unit]

    if !ok{
       return false 
    }


    bill[item] += quantity
    
   

    return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
     currQty,ok := bill[item]
    

    if !ok{
        return false
    }

    qty,_ := units[unit]
    if qty == 0 {
        return false
    }

    newQty := currQty - qty 

    if newQty < 0{
        return false
    }

    if newQty == 0 {
        delete(bill, item)
        return true
    }else{
        bill[item] = newQty
    }

    return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	qty,ok := bill[item]

    return qty,ok
}

package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if _, ok := units[unit]; !ok {
		return ok
	} else {
		bill[item] += units[unit]
	}

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	itemQuantity, itemInBill := bill[item]
	unitQuantity, unitInMap := units[unit]
	if !itemInBill || !unitInMap {
		return false
	}
	diff := itemQuantity - unitQuantity
	if diff < 0 {
		return false
	} else if diff == 0 {
		delete(bill, item)
	} else {
		bill[item] = diff
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	if itemQtd, ok := bill[item]; !ok {
		return 0, false
	} else {
		return itemQtd, true
	}
}

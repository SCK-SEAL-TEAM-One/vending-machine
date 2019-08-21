package vending

type VendingMachine struct {
	Drinks       map[string]int
	TotalBalance int
}

func (vendingMachine VendingMachine) calculateChange(drinkPrice int) int {
	return 5
}

func (vendingMachine VendingMachine) canBuyDrink(drinkPrice int) bool {
	return true
}

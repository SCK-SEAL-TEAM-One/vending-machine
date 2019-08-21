package vending

type VendingMachine struct {
	Drinks       map[string]int
	TotalBalance int
}

func (vendingMachine VendingMachine) BuyDrink(item string) (string, int) {
	price := vendingMachine.getDrinkPrice(item)
	if vendingMachine.canBuyDrink(price) {
		change := vendingMachine.calculateChange(price)
		return item, change
	}
	return "", vendingMachine.TotalBalance
}

func (vendingMachine VendingMachine) canBuyDrink(drinkPrice int) bool {
	if vendingMachine.TotalBalance > drinkPrice {
		return true
	}
	return false
}

func (vendingMachine VendingMachine) calculateChange(drinkPrice int) int {
	return vendingMachine.TotalBalance - drinkPrice
}

func (vendingMachine VendingMachine) getDrinkPrice(drink string) int {
	return vendingMachine.Drinks[drink]
}

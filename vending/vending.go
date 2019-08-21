package vending

type VendingMachine struct {
	Drinks       map[string]int
	TotalBalance int
	Coins        map[string]int
}

func (vendingMachine VendingMachine) BuyDrink(item string) (string, int) {
	price := vendingMachine.getDrinkPrice(item)
	if !vendingMachine.canBuyDrink(price) {
		return "", vendingMachine.TotalBalance
	}
	change := vendingMachine.calculateChange(price)
	return item, change
}

func (vendingMachine VendingMachine) canBuyDrink(drinkPrice int) bool {
	return vendingMachine.TotalBalance >= drinkPrice
}

func (vendingMachine VendingMachine) calculateChange(drinkPrice int) int {
	return vendingMachine.TotalBalance - drinkPrice
}

func (vendingMachine VendingMachine) getDrinkPrice(drink string) int {
	return vendingMachine.Drinks[drink]
}

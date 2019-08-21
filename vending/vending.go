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
func (vendingMachine VendingMachine) getDrinkPrice(item string) int {
	if item == "CF" {
		return 15
	}
	return 8
}

func (vendingMachine VendingMachine) canBuyDrink(price int) bool {
	return true
}

func (vendingMachine VendingMachine) calculateChange(drinkPrice int) int {
	return vendingMachine.TotalBalance - drinkPrice
}

package vending

type VendingMachine struct {
	Drinks       map[string]int
	TotalBalance int
}

func (vendingMachine VendingMachine) BuyDrink(item string) (string, int) {
	price := vendingMachine.getDrinkPrice(item)
	vendingMachine.canBuyDrink(price)
	change := vendingMachine.calculateChange(price)
	return item, change
}
func (vendingMachine VendingMachine) getDrinkPrice(item string) int {
	return 8
}
func (vendingMachine VendingMachine) calculateChange(price int) int {
	return 0
}

func (vendingMachine VendingMachine) canBuyDrink(price int) bool {
	return true
}

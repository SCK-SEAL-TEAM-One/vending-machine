package vending

type VendingMachine struct {
	Drinks       map[string]int
	TotalBalance int
	Coins        map[string]int
}

func (vendingMachine VendingMachine) BuyDrink(item string) (string, []string) {
	price := vendingMachine.getDrinkPrice(item)
	if !vendingMachine.canBuyDrink(price) {
		return "", []string{}
	}
	change := vendingMachine.calculateChange(price)
	return item, vendingMachine.getChangeCoins(change)
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

func (vendingMachine VendingMachine) getChangeCoins(change int) []string {
	return []string{"T"}
}

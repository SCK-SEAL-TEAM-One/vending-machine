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

func (vendingMachine *VendingMachine) InsertCoin(coin string) {
	vendingMachine.TotalBalance += vendingMachine.getCoinValue(coin)
}

<<<<<<< HEAD
func (vendingMachine VendingMachine) getCoinValue(coin string) int {
	return vendingMachine.Coins[coin]
=======
func getCoinsValue(coin string) int {
	if coin == "O" {
		return 1
	}
	if coin == "TW" {
		return 2
	}
	if coin == "F" {
		return 5
	}
	return 10
>>>>>>> 3473eafa115321721fe77f292e059b4f7db2e420
}

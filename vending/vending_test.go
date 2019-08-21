package vending

import (
	"testing"
)

func Test_BuyDrink_Input_TotalBalance_8_Item_DW_Should_Be_Item_DW_Balance_0(t *testing.T) {
	expetedItem, expetedBalance := "DW", 0
	drink := "DW"
	drinks := map[string]int{
		"SD": 25,
		"CF": 15,
		"DW": 8,
	}
	vendingMachine := VendingMachine{
		TotalBalance: 8,
		Drinks:       drinks,
	}

	actualItem, actualBalance := vendingMachine.BuyDrink(drink)

	if expetedItem != actualItem {
		t.Errorf("Expect %s but got %s", expetedItem, actualItem)
	}
	if expetedBalance != actualBalance {
		t.Errorf("Expect %d but got %d", expetedBalance, actualBalance)
	}
}

func Test_BuyDrink_Input_TotalBalance_20_Item_CF_Should_Be_Item_CF_Balance_5(t *testing.T) {
	expetedItem, expetedBalance := "CF", 5
	dirnk := "CF"
	drinks := map[string]int{
		"SD": 25,
		"CF": 15,
		"DW": 8,
	}
	vendingMachine := VendingMachine{
		TotalBalance: 20,
		Drinks:       drinks,
	}

	actualItem, actualBalance := vendingMachine.BuyDrink(dirnk)

	if expetedItem != actualItem {
		t.Errorf("Expect %s but got %s", expetedItem, actualItem)
	}
	if expetedBalance != actualBalance {
		t.Errorf("Expect %d but got %d", expetedBalance, actualBalance)
	}
}

func Test_CalculateChange_Input_TotalBalance_20_DrinkPrice_15_Should_Be_5(t *testing.T) {
	expect := 5
	vendingMachine := VendingMachine{
		TotalBalance: 20,
	}
	dirnkPrice := 15

	actual := vendingMachine.calculateChange(dirnkPrice)

	if expect != actual {
		t.Errorf("Expect %d but got %d", expect, actual)
	}
}

func Test_CalculateChange_Input_TotalBalance_8_DrinkPrice_8_Should_Be_0(t *testing.T) {
	expect := 0
	vendingMachine := VendingMachine{
		TotalBalance: 8,
	}
	dirnkPrice := 8

	actual := vendingMachine.calculateChange(dirnkPrice)

	if expect != actual {
		t.Errorf("Expect %d but got %d", expect, actual)
	}
}

func Test_CanBuyDrink_Input_Price_8_Should_True(t *testing.T) {
	expected := true
	drinkPrice := 8
	vendingMachine := VendingMachine{
		TotalBalance: 8,
	}

	actual := vendingMachine.canBuyDrink(drinkPrice)

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_CanBuyDrink_Input_Price_15_Should_True(t *testing.T) {
	expected := true
	drinkPrice := 15
	vendingMachine := VendingMachine{
		TotalBalance: 20,
	}

	actual := vendingMachine.canBuyDrink(drinkPrice)

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_CanBuyDrink_Input_Price_15_Should_False(t *testing.T) {
	expected := false
	drinkPrice := 15
	vendingMachine := VendingMachine{
		TotalBalance: 10,
	}

	actual := vendingMachine.canBuyDrink(drinkPrice)

	if expected != actual {
		t.Errorf("Expect %v but got %v", expected, actual)
	}
}

func Test_GetDrinkPrice_Input_Drink_CF_Should_Be_15(t *testing.T) {
	expect := 15
	drink := "CF"
	drinks := map[string]int{
		"SD": 25,
		"CF": 15,
		"DW": 8,
	}
	vendingMachine := VendingMachine{
		Drinks: drinks,
	}

	actual := vendingMachine.getDrinkPrice(drink)
	if expect != actual {
		t.Errorf("Expect %d but got %d", expect, actual)
	}
}

func Test_GetCoinValue_Input_Coin_F_Should_Be_5(t *testing.T) {
	expected := 5
	coin := "F"
	coins := map[string]int{
		"T":  10,
		"F":  5,
		"TW": 2,
		"O":  1,
	}
	vendingMachine := VendingMachine{
		Coins: coins,
	}

	actual := vendingMachine.getCoinValue(coin)

	if expected != actual {
		t.Errorf("Expect %d but got %d", expected, actual)
	}
}

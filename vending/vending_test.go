package vending

import (
	"testing"
)

func Test_BuyDrink_Input_TotalBalance_8_Item_DW_Should_Be_Item_DW_No_Change(t *testing.T) {
	expetedItem, expetedChange := "DW", []string{}
	drink := "DW"
	drinks := map[string]int{
		"SD": 25,
		"CF": 15,
		"DW": 8,
	}
	coins := map[string]int{
		"T":  10,
		"F":  5,
		"TW": 2,
		"O":  1,
	}
	vendingMachine := VendingMachine{
		Drinks: drinks,
		Coins:  coins,
	}

	vendingMachine.InsertCoin("F")
	vendingMachine.InsertCoin("TW")
	vendingMachine.InsertCoin("O")
	actualItem, actualChange := vendingMachine.BuyDrink(drink)

	if expetedItem != actualItem {
		t.Errorf("Expect %s but got %s", expetedItem, actualItem)
	}
	if expetedChange != actualChange {
		t.Errorf("Expect %d but got %d", expetedChange, actualChange)
	}
}

func Test_BuyDrink_Input_TotalBalance_20_Item_CF_Should_Be_Item_CF_Change_F(t *testing.T) {
	expetedItem, expetedChange := "CF", []string{"F"}
	dirnk := "CF"
	drinks := map[string]int{
		"SD": 25,
		"CF": 15,
		"DW": 8,
	}
	coins := map[string]int{
		"T":  10,
		"F":  5,
		"TW": 2,
		"O":  1,
	}
	vendingMachine := VendingMachine{
		Drinks: drinks,
		Coins:  coins,
	}

	vendingMachine.InsertCoin("T")
	vendingMachine.InsertCoin("T")
	actualItem, actualChange := vendingMachine.BuyDrink(dirnk)

	if expetedItem != actualItem {
		t.Errorf("Expect %s but got %s", expetedItem, actualItem)
	}
	if expetedChange != actualChange {
		t.Errorf("Expect %d but got %d", expetedChange, actualChange)
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

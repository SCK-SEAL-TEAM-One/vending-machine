package vending

import (
	"testing"
)

func Test_BuyDrink_Input_TotalBalance_8_Item_DW_Should_Be_Item_DW_Balance_0(t *testing.T) {
	expetedItem, expetedBalance := "DW", 0
	item := "DW"
	vendingMachine := VendingMachine{
		TotalBalance: 8,
	}

	actualItem, actualBalance := vendingMachine.BuyDrink(item)

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

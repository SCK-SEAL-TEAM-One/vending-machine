package vending

import "testing"

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

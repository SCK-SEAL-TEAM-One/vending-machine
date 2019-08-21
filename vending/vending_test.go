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

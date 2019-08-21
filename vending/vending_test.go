package vending

import (
	"testing"
)

func Test_BuyDrink_Input_TotalBalance_8_Item_DW_Should_Be_Item_DW_Balance_0(t *testing.T) {
	expetedItem, expetedBalance := "DW", 0
	item := "DW"
	vending := VendingMachine{
		TotalBalance: 8,
	}

	actualItem, actualBalance := vending.BuyDrink(item)

	if expetedItem != actualItem {
		t.Errorf("Expect %s but got %s", expetedItem, actualItem)
	}
	if expetedBalance != actualBalance {
		t.Errorf("Expect %d but got %d", expetedBalance, actualBalance)
	}
}

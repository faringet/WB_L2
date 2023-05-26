package main

import "fmt"

type HasMoneyState struct {
	VendingMachine *VendingMachine
}

func (i *HasMoneyState) AddItem(count int) error {
	return fmt.Errorf("item dispense in progress")
}

func (i *HasMoneyState) RequestItem() error {
	return fmt.Errorf("item dispense in progress")
}

func (i *HasMoneyState) InsertMoney(money int) error {
	return fmt.Errorf("item out of stock")
}

func (i *HasMoneyState) DispenseItem() error {
	fmt.Println("Dispensing Item")
	i.VendingMachine.itemCount = i.VendingMachine.itemCount - 1
	if i.VendingMachine.itemCount == 0 {
		i.VendingMachine.setState(i.VendingMachine.noItem)
	} else {
		i.VendingMachine.setState(i.VendingMachine.hasItem)
	}
	return nil
}

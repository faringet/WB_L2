package main

import "fmt"

type NoItemState struct {
	VendingMachine *VendingMachine
}

func (i *NoItemState) AddItem(count int) error {
	i.VendingMachine.IncrementItemCount(count)
	i.VendingMachine.setState(i.VendingMachine.hasItem)
	return nil
}

func (i *NoItemState) RequestItem() error {
	return fmt.Errorf("item out of stock")
}

func (i *NoItemState) InsertMoney(money int) error {
	return fmt.Errorf("item out of stock")
}

func (i *NoItemState) DispenseItem() error {
	return fmt.Errorf("item out of stock")
}

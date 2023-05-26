package main

import "fmt"

type ItemRequestedState struct {
	VendingMachine *VendingMachine
}

func (i *ItemRequestedState) AddItem(count int) error {
	return fmt.Errorf("item dispense in progress")
}

func (i *ItemRequestedState) RequestItem() error {
	return fmt.Errorf("item already requested")
}

func (i *ItemRequestedState) InsertMoney(money int) error {
	if money < i.VendingMachine.itemPrice {
		fmt.Errorf("insered money is less. Please insert %d", i.VendingMachine.itemPrice)
	}
	fmt.Println("Money entered is ok")
	i.VendingMachine.setState(i.VendingMachine.hasMoney)
	return nil
}

func (i *ItemRequestedState) DispenseItem() error {
	return fmt.Errorf("please insert money first")
}

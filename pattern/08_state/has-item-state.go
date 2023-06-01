package main

import "fmt"

type HasItemState struct {
	VendingMachine *VendingMachine
}

func (i *HasItemState) AddItem(count int) error {
	fmt.Printf("%d items added\n", count)
	i.VendingMachine.IncrementItemCount(count)
	return nil

}

func (i *HasItemState) RequestItem() error {
	if i.VendingMachine.itemCount == 0 {
		i.VendingMachine.setState(i.VendingMachine.noItem)
		return fmt.Errorf("no item present")
	}
	fmt.Printf("Item requested\n")
	i.VendingMachine.setState(i.VendingMachine.itemRequested)
	return nil
}

func (i *HasItemState) InsertMoney(money int) error {
	return fmt.Errorf("please select item first")

}

func (i *HasItemState) DispenseItem() error {
	return fmt.Errorf("please select item first")
}

package main

import "time"

type Card struct {
	Name    string
	Balance float64
	Bank    *Bank
}

func (card Card) CheckBalance() error {
	println("[Карта] Запрос в банк для проверки остатка по карте")
	time.Sleep(time.Millisecond * 800)
	return card.Bank.CheckBalance(card.Name)

}

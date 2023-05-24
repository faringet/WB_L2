package main

import "fmt"

var (
	bank = Bank{
		Name:  "БАНК",
		Cards: []Card{},
	}
	card1 = Card{
		Name:    "CRD-1",
		Balance: 200,
		Bank:    &bank,
	}
	card2 = Card{
		Name:    "CRD-2",
		Balance: 5,
		Bank:    &bank,
	}
	user = User{
		Name: "Покупатель-1",
		Card: &card1,
	}
	user2 = User{
		Name: "Покупатель-2",
		Card: &card2,
	}
	prod = Product{
		Name:  "Сыр",
		Price: 150,
	}
	shop = Shop{
		Name: "SHOP",
		Products: []Product{
			prod,
		},
	}
)

func main() {
	println("[Банк] Выпуск кард")
	bank.Cards = append(bank.Cards, card1, card2)
	fmt.Printf("[%s]", user.Name)
	err := shop.Sell(user, prod.Name)
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("[%s]", user2.Name)
	err = shop.Sell(user2, prod.Name)
	if err != nil {
		println(err.Error())
		return
	}
}

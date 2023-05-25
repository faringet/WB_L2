package main

import (
	"errors"
	"fmt"
	"time"
)

type Shop struct {
	Name     string
	Products []Product
}

// Sell пользователь, приходя в магазин покупает товар
func (shop Shop) Sell(user User, product string) error {
	println("[Магазин] Запрос к пользователю, для получения остатка по карте")
	time.Sleep(time.Millisecond * 500)
	err := user.Card.CheckBalance()
	if err != nil {
		return err
	}
	fmt.Printf("[Магазин] проверка - может ли [%s] купить товар", user.Name)
	time.Sleep(time.Millisecond * 500)
	for _, prod := range shop.Products {
		if prod.Name != product {
			continue
		}
		if prod.Price > user.GetBalance() {
			return errors.New("[Магазин] недостаточно средств для покупки товара")
		}
		fmt.Printf("[Магазин] Товар [%s] - куплен\n", prod.Name)
	}
	return nil
}

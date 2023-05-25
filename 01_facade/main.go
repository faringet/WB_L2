package main

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
Структурный паттерн проектирования. Предоставляет простой доступ (интерфейс) к сложной системе.

Когда у нас много разных подсистем, которые реализуют свой функционал поведения.
func (shop Shop) Sell - как раз и будет общим фасадом над всей платежной системой.
Как пример - интернет магазин, который осуществляет продажу товаров.
Скрываем основную бизнес логику поведения, а юзеру даем минимум функционала.

+ изолирует клиент от поведения сложной системы
+ сам фасад простой

- сам фасад может стать божественным объектом (мы привязаны к этому объекту и он будет во всей системе) - фсе функции
будут проходить через этот объект

Реализация покупки товара каким-либо пользователем интернет магазина:
*/

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

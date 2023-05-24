package main

type User struct {
	Name string
	Card *Card
}

// GetBalance у ползователя нет поля баланс, но у него есть карта по которой он может запросить инфу у банка. Банк вернет баланс.
func (user User) GetBalance() float64 {
	return user.Card.Balance

}

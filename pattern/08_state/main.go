package main

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

/*
Поведенческий паттерн проектирования.
Позволяет объектам менять поведение в зависимости от своего (извне может показаться что изменился весь объект).
Основная идея в том что программа может находиться в одном из нескольких состояний, которые все время сменяют друг друга.
Набор этих состояний, а также переходов между ними всегда имеет окончание.
Находясь в разных состояниях программа может по-разному реагировать на одни и те же события (которые происходят с ней).
Такой подход можно применить и к объектам.
Например - документ. Он может принимать 3 состояния (черновик, модерация, опубликоваться). В каждом из состояний метод
опубликовать будет работать по-разному. Из черновика отправит на модерацию, а с модерации - в публикацию (при условии, что
это сделал админ).
Паттерн состояние предлагает создать отдельные объекты для каждого состояния в котором может прибывать объект, а затем внести
туда поведение - соответсвующее текущему состоянию.
Контекст (первоначальный объект) будет содержать ссылку на один из объектов состояний и делегировать ему работу (зависящей от состояния).

+ избавляет от множества больших условных операторов
+ концентрирует в одном месте код, связанный с определенным состоянием
+ упрощает код контекста

- может неоправданно усложнить код если состояний мало или они редко меняются

В качестве примера напишем софт для торгового автомата.
1) торговый автомат может выдавать только один товар
2) автомат может находиться только в одном из 4 состояний, а именно - [выдать предмет], [добавить предмет], [унести деньги], [добавить предмет]

*/

func main() {

}
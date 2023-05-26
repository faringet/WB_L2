package main

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

/*
Порождающий паттерн проектирования.
Определяет общий интерфейс поведения для создаваемых объектов. Есть общий интерфейс этих объектов.
Какие проблемы он решает ?
Предположим у нас есть интернет магазин, который продает серваки. Мы полностью реализовали поведение интернет магазина
по продаже сервера. Но через какое-то время мы хотим продавать и ПК. Но основная логика у нас завязана на серверах.
Чтобы не создавать еще одну логику поведения для ПК.
Как решим ?
Создадим общую реализацию создания объектов при помощи фабричного метода. В этот метод передаем название типа того объекта,
который хотим создать. В зависимости от этого типа и будет возвращаться создаваемый объект.

+ избавляет от привязки к конкретному типу объекта (есть интерфейс поведения)
+ создается общий конструктор создания объектов
+ упрощает добавление новых объектов
+ реализует принцип открытости и закрытости

- может привести к большим параллельным иерархий объектов (может быть большое количество структур)
- божественный конструктор
*/

var types = []string{PersonalComputerType, NotebookType, ServerType, "monoblock"}

func main() {
	for _, typeName := range types {
		computer := New(typeName)
		if computer == nil {
			continue
		}
		computer.PrintDetails()
	}

}

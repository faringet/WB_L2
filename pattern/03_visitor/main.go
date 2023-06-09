package main

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

/*
Поведенческий паттерн проектирования.
Позволяет добавлять поведение в структуру без ее изменения.
Какие проблемы он решает ?
Предположим мы написали библиотеку, которая содержит три структуры - квадрат, круг и прямоугольник.
Нас попросили добавить функционал, который будет возвращать площадь фигуры (getArea).
В теории мы можем тупо добавить getArea в интерфейс фигуры, а затем реализовать.
Но это вообще не бест практис.
Как решим ?
Очевидно, с помощью посетителя!

1) Для этого определим интерфейс посетителя:

type visitor interface {
   visitForSquare(square) // добавляем функционал для квадрата
   visitForCircle(circle) // добавляем функционал для круга
   visitForRectangle(rectangle) // добавляем функционал для прямоугольника
}

2) Добавление метода accept в интерфейс фигуры:

func accept(v visitor)

Все структуры фигур должны определять этот метод похожим способом:

func (obj *square) accept(v visitor){
    v.visitForSquare(obj)
}

В конечном итоге, структуры нужно изменить лишь единожды, и все будущие запросы нового функционала можно будет
реализовать с помощью функции accept(v visitor). Если команда запросит поведение getArea, мы можем просто определить
явную реализацию интерфейса посетителя и прописать логику вычисления площади в этой конкретной имплементации.

+ упрощает добавление операций, работающих со сложными структурами объектов
+ объединяет родственные операции в одном классе

- не оправдан, если иерархия элементов часто меняется.
- может привести к нарушению инкапсуляции элементов.
*/

func main() {
	square := &Square{side: 2}
	circle := &Circle{radius: 3}
	rectangle := &Rectangle{l: 2, b: 3}

	areaCalculator := &AreaCalculator{}

	square.accept(areaCalculator)
	circle.accept(areaCalculator)
	rectangle.accept(areaCalculator)

	fmt.Println()
	middleCoordinates := &MiddleCoordinates{}
	square.accept(middleCoordinates)
	circle.accept(middleCoordinates)
	rectangle.accept(middleCoordinates)
}

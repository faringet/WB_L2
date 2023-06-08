Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
error
```
Задача очень похожа на [**listing03**](https://github.com/faringet/WB_L2/blob/master/listing/listing03.md)

error - это интерфейс

Переменная err содержит в себе:
* тип `*main.customError`

* значение `nil`

При этом переменная не считается равной `nil`. Во время сравнения `err != nil` ожидаемо получаем вывод в строку `println("error")`. (так как содержит тип).



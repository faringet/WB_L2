Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
Будет рандомный вывод от 1 до 8, а потом бесконечные нули

```

Ошибка в том что `func asChan(vs ...int) <-chan int` возвращает закрытый канал `c`--> каналы `a` и `b` также закрыты --> `case v := <-a` 
и `case v := <-b` читает из зкарытых каналов --> получаем бесконечные нули

Почему нули ?

Когда канал закрыт, значение, считанное горутиной, является нулевым значением, в зависимости от типа данных канала. Так как в нашем случае тип данных канала int, то нулевое значение будет 0.

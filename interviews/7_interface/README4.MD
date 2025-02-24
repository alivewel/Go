### nil-интерфейсы в Go

В Go очень **интересно** ведут себя `nil`-интерфейсы. Это один из [50 подводных камней](https://habr.com/ru/companies/vk/articles/314804/#54), которые неочевидны при изучении языка. 

Мы можем создать переменную пустого интерфейса, и при сравнении с `nil`, мы получаем `true`, что означает `interface == nil`. Далее мы можем создать указатель на структуру и снова сравнить его с `nil`, мы получаем `true`. Все логично. 

Теперь присвоим переменной интерфейса указатель на структуру. Теперь мы получаем `false`. Почему так происходит? После присвоения у интерфейса указателя на структуру, в интерфейсе записывается конкретный тип, это значит, что значение интерфейса уже не равно `nil`. Поэтому при сравнении мы получаем `false`.

``` go
package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct {
	Name string
}

func main() {
	var i interface{}
	// type == nil, value == nil
	// поэтому i == nil
	fmt.Println(i == nil) // true

	var d *Dog
	fmt.Println(d == nil) // true

	i = d
	// type == *Dog, value == nil
	// поскольку type != nil, то i != nil
	fmt.Println(i == nil) // false
}
```

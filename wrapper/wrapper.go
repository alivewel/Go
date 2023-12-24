package main

import "fmt"

// Определяем интерфейс для базового типа
type Base interface {
	DoSomething()
}

// Реализуем базовый тип
type ConcreteBase struct{}

func (c *ConcreteBase) DoSomething() {
	fmt.Println("Базовый тип выполняет действие")
}

// Определяем обертку
type Wrapper struct {
	Base
}

func (w *Wrapper) DoSomething() {
	fmt.Println("Обертка выполняет действие перед базовым типом")
	w.Base.DoSomething()
	fmt.Println("Обертка выполняет действие после базового типа")
}

func main() {
	// Создаем экземпляр базового типа
	base := &ConcreteBase{}

	// Создаем экземпляр обертки
	wrapper := &Wrapper{base}

	// Вызываем метод обертки, который вызывает метод базового типа
	wrapper.DoSomething()
}

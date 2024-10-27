package main

import (
	"fmt"
)

type Runner interface {
	Run() string
}

type Human struct {
	Name string
}

func (h Human) Run() string {
	return fmt.Sprintf("Человек %s бегает", h.Name)
}

func main() {
	var runner Runner
	fmt.Printf("Value %v, type %T\n", runner, runner) // Выведем значение и тип данной структуры
	if runner != nil {
		fmt.Println("runner is not nil")
	}

	// runner.Run() // паника при попытке вызвать
	// runner = "asf" // не скомпилируется, string does not implement Runner (missing method Run)

	// var unnamedRunner *Human // если создаем таким образом, то ловим сегу при попытке заполнить поле Name
	unnamedRunner := &Human{}

	runner = unnamedRunner
	runner.Run() // OK
	fmt.Printf("Type: %T Value: %#v\n", runner, runner)
	if runner != nil {
		fmt.Println("runner is not nil")
	}
	unnamedRunner.Name = "Peter"
	fmt.Printf("Type: %T Value: %#v\n", runner, runner)
}

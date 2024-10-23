### Реализовать функцию, которая возвращает ошибку, не используя импорт других библиотек

```go
type myError struct {}

func (e *myError) Error() string {
	return "my error"
}

func handle() error {
	return &myError{}
}

func main() {
	err := handle()
	if err != nil {
		println(err.Error())
	}
}
```

Для начала необходимо создать собственную структуру ошибки `myError`. Далее необходимо реализовать интерфейс `error`, с помощью метода `Error()`, который возвращает `string`.


<details>
<summary>Задача №1 (concurrency)</summary>

### Что выведет программа?

```go
func main() {
	for _, val := range []int{1, 2, 3} {
		go println(val)
	}
}
```

### Отсутствие вывода

Программа не выводит ничего на экран. Это происходит потому, что основная горутина `main` завершается раньше, чем горутины успевают выполнить `println`. В результате, программа завершает работу, не дождавшись завершения всех запущенных горутин.

### Использование `WaitGroup`

Чтобы гарантировать, что все горутины завершат выполнение, прежде чем программа выйдет, необходимо использовать примитив синхронизации `sync.WaitGroup`. Этот механизм позволяет основной горутине `main` ожидать завершения всех запущенных горутин.

### Захват переменной в замыкании

В зависимости от версии Go, необходимость добавления новой переменной для захвата значения в замыкании может изменяться. Начиная с версии 1.22, дополнительная переменная не обязательна, так как компилятор Go сам обрабатывает захват переменных в контексте горутин.

### Решение

```go
func main() {
	wg := sync.WaitGroup{}
	for _, val := range []int{1, 2, 3} {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			println(num)
		}(val)
	}
	wg.Wait()
}
```

</details>
<br>


<details>
<summary>Задача №2 (concurrency)</summary>



</details>
<br>

<details>
<summary>Задача №3 (string)</summary>



</details>
<br>

<details>
<summary>Задача №4 (string)</summary>



</details>
<br>

<details>
<summary>Задача №5 (string)</summary>



</details>
<br>

<details>
<summary>Задача №6 (slice)</summary>



</details>
<br>

<details>
<summary>Задача №7 (slice)</summary>



</details>
<br>

<details>
<summary>Задача №8 (map)</summary>



</details>
<br>

<details>
<summary>Задача №9 (map)</summary>



</details>
<br>

<details>
<summary>Задача №10 (interface)</summary>



</details>
<br>

<details>
<summary>Задача №11 (defer)</summary>



</details>
<br>

<details>
<summary>Задача №12 (error)</summary>



</details>
<br>

<details>
<summary>Задача №21 (merge-channels)</summary>



</details>
<br>

<details>
<summary>Задача №22 (add-two-numbers)</summary>



</details>
<br>

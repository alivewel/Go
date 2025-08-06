Отличная задача на использование контекста и управление временем выполнения!

---

## 🧠 Цель задачи:

* Если `process()` завершается **быстрее**, чем `processingMaxTimeout` (3 секунды), то вывести `1`.
* Если не успевает — вывести `0`.

Функция `process()` спит 2 секунды, т.е. **в данный момент она укладывается**.

---

## ✅ Правильная вставка (вместо `// === edit here ===`):

```go
    done := make(chan int, 1)

    go func() {
        result := process()
        done <- result
    }()

    select {
    case res := <-done:
        result = res
    case <-ctx.Done():
        result = 0
    }
```

---

### 🔍 Как это работает:

* `done` — буферизированный канал на 1 значение, через который `process()` вернёт результат.
* Запускаем `process()` в **отдельной горутине**.
* Через `select`:

  * Если `process()` завершился вовремя — читаем `res` и присваиваем `result`.
  * Если время вышло (`ctx.Done()`) — пишем `0`.

---

## 💻 Полный рабочий код:

```go
package main

import (
    "context"
    "fmt"
    "time"
)

const processingMaxTimeout = 3 * time.Second

func main() {
    var result int

    ctx, cancel := context.WithTimeout(context.Background(), processingMaxTimeout)
    defer cancel()

    done := make(chan int, 1)

    go func() {
        result := process()
        done <- result
    }()

    select {
    case res := <-done:
        result = res
    case <-ctx.Done():
        result = 0
    }

    fmt.Println(result)
}

func process() int {
    // long processing job
    time.Sleep(time.Second * 2)

    return 1
}
```

---

### 🧪 Проверь:

* `Sleep(2s)` — напечатает `1`
* `Sleep(4s)` — напечатает `0`

Хочешь усложнённый вариант с `context.WithCancel` или с несколькими `process()` параллельно?

1) Поиск по очень большому неотсортированному массиву с распараллеливанием на Go

```go
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	const size = 100_000_000
	const target = 42

	// Создаём большой неотсортированный массив
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(1_000_000)
	}
	// Ставим цель, чтобы точно была
	arr[size/2] = target

	numCPU := runtime.NumCPU()
	chunkSize := size / numCPU
	var wg sync.WaitGroup
	found := make(chan bool, 1) // буфер 1, чтобы не заблокироваться

	start := time.Now()

	for i := 0; i < numCPU; i++ {
		wg.Add(1)
		go func(startIndex int) {
			defer wg.Done()
			end := startIndex + chunkSize
			if end > size {
				end = size
			}
			for _, v := range arr[startIndex:end] {
				if v == target {
					select {
					case found <- true:
					default: // если уже найден — не блокируемся
					}
					return
				}
			}
		}(i * chunkSize)
	}

	go func() {
		wg.Wait()
		close(found)
	}()

	if <-found {
		fmt.Println("Число найдено!")
	} else {
		fmt.Println("Число не найдено.")
	}

	fmt.Println("Время поиска:", time.Since(start))
}
```

### ✅ Полный пример с отменой через `context`:

```go
package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	const size = 100_000_000
	const target = 42

	// 1. Создаём большой массив
	arr := make([]int, size)
	for i := range arr {
		arr[i] = rand.Intn(1_000_000)
	}
	arr[size/2] = target // точно вставим элемент

	// 2. Настройка распараллеливания
	numCPU := runtime.NumCPU()
	chunkSize := size / numCPU

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // safety

	var wg sync.WaitGroup
	found := make(chan bool, 1)

	start := time.Now()

	for i := 0; i < numCPU; i++ {
		wg.Add(1)
		go func(startIndex int) {
			defer wg.Done()
			end := startIndex + chunkSize
			if end > size {
				end = size
			}

			for _, v := range arr[startIndex:end] {
				select {
				case <-ctx.Done():
					// Контекст отменён — выходим досрочно
					return
				default:
					if v == target {
						select {
						case found <- true:
						default:
						}
						cancel() // отменить все остальные горутины
						return
					}
				}
			}
		}(i * chunkSize)
	}

	go func() {
		wg.Wait()
		close(found)
	}()

	if <-found {
		fmt.Println("✅ Число найдено")
	} else {
		fmt.Println("❌ Число не найдено")
	}

	fmt.Println("⏱ Время поиска:", time.Since(start))
}
```

2) В Go Escape analysis, в PostgreSQL EXPLAIN ANALYZE.

Для вызова escape analysis в Go используется флаг `-gcflags="-m"`.

3) Вот краткая, но точная **теория по сборщику мусора (ГЦ)** в Go, основанная на **алгоритме Mark and Sweep**, с учётом того, **как он работает в Go-runtime**.

---

## ♻️ Что такое GC (Garbage Collector)

Сборщик мусора — это компонент рантайма, который **автоматически освобождает неиспользуемую память**, чтобы предотвратить утечки и переполнение.

---

## 📌 Алгоритм **Mark and Sweep** — общая идея

Это классический двухфазный алгоритм:

1. **Mark (отметь)**:

   * GC проходит по всем **живым объектам** (достижимым из root-объектов: стеков, глобальных переменных и т.п.).
   * Помечает все объекты, до которых может дойти по ссылкам.

2. **Sweep (собери)**:

   * Проходит по всей памяти и **удаляет объекты, не отмеченные** на предыдущем шаге.
   * Освобождает память под них.

---

## 🔧 Как работает GC в Go (начиная с Go 1.5+)

Go использует **трассирующий, инкрементальный, многопоточный, параллельный GC**.
В основе — **Mark and Sweep**, но с важными улучшениями:

### 💡 Основные фазы:

| Фаза                 | Что происходит                                           |
| -------------------- | -------------------------------------------------------- |
| **STW - Start**      | Короткая пауза, запускает GC                             |
| **Mark**             | Потоками обходит объекты из root-наборов, помечает живые |
| **Concurrent Sweep** | Параллельно удаляет неиспользуемые объекты               |
| **STW - End**        | Завершает GC, наводит порядок                            |

### 🧵 Параллельно с работой приложения:

* GC **не останавливает мир (Stop The World) надолго**, как раньше.
* Только **в начале и в конце** — короткие STW-фазы (несколько миллисекунд).
* Основная часть работы происходит **в фоне (background GC)**, без полной остановки программы.

---

## ⚙️ Как GC понимает, что пора чистить?

Go GC работает **по триггеру объёма выделений**:

* Когда программа выделила **\~100% от предыдущего цикла GC**, запускается новый.
* Можно контролировать через `GOGC` (по умолчанию `GOGC=100`):

  * `GOGC=200` — реже GC, больше памяти
  * `GOGC=20` — часто GC, меньше памяти

---

## 🧠 Что считается достижимым (живым)?

* Локальные переменные функций (на стеке)
* Объекты, на которые ссылаются глобальные переменные
* Объекты, на которые ссылаются другие живые объекты

---

## 📊 Мониторинг GC

Можно использовать:

```go
var m runtime.MemStats
runtime.ReadMemStats(&m)
fmt.Println("NumGC:", m.NumGC)
```

Или `pprof`, `go tool trace`, `GODEBUG=gctrace=1`

---

## 🔬 Пример вывода `GODEBUG=gctrace=1`

```bash
GODEBUG=gctrace=1 go run main.go
```

```
gc 1 @1.202s 6%: 0.006+0.88+0.003 ms clock, 0.03+0.22/0.71/0.13+0.01 ms cpu
```

Расшифровка:

| Часть              | Значение                                         |
| ------------------ | ------------------------------------------------ |
| `gc 1 @1.202s`     | GC #1 запущен через 1.2 секунды                  |
| `6%`               | Потрачено 6% времени программы на GC             |
| `0.006+0.88+0.003` | Пауза STW + mark + final pause (в миллисекундах) |

---

## ✅ Вывод

| Особенность Go GC | Описание                                                            |
| ----------------- | ------------------------------------------------------------------- |
| Алгоритм          | Mark and Sweep                                                      |
| Поведение         | Инкрементальный, параллельный, stop-the-world только в начале/конце |
| Старт             | По достижению порога выделенной памяти (`GOGC`)                     |
| Преимущества      | Минимальные паузы, нет необходимости вручную управлять памятью      |
| Можно ли влиять   | Да — через `GOGC`, `MemStats`, `GODEBUG`                            |

---

Хочешь — могу объяснить **влияние GC на производительность**, или показать **пример программы с профилированием GC**.

4) Отслеживание уникальности

``` go
package main

import (
	"fmt"
)

func mergeUnique(a, b []int) []int {
	seen := make(map[int]bool)
	var result []int

	for _, v := range a {
		if !seen[v] {
			result = append(result, v)
			seen[v] = true
		}
	}
	for _, v := range b {
		if !seen[v] {
			result = append(result, v)
			seen[v] = true
		}
	}
	return result
}

func main() {
	a := []int{1, 2, 3, 4}
	b := []int{3, 4, 5, 6}
	merged := mergeUnique(a, b)
	fmt.Println(merged) // [1 2 3 4 5 6]
}
```

5) Решение — два указателя
``` go
func mergeSortedUnique(a, b []int) []int {
	i, j := 0, 0
	var result []int

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			if len(result) == 0 || result[len(result)-1] != a[i] {
				result = append(result, a[i])
			}
			i++
		} else if a[i] > b[j] {
			if len(result) == 0 || result[len(result)-1] != b[j] {
				result = append(result, b[j])
			}
			j++
		} else { // a[i] == b[j]
			if len(result) == 0 || result[len(result)-1] != a[i] {
				result = append(result, a[i])
			}
			i++
			j++
		}
	}

	// Добавим оставшиеся элементы (если есть)
	for i < len(a) {
		if result[len(result)-1] != a[i] {
			result = append(result, a[i])
		}
		i++
	}

	for j < len(b) {
		if result[len(result)-1] != b[j] {
			result = append(result, b[j])
		}
		j++
	}

	return result
}
```

6) Merge

``` go
package main

import (
	"fmt"
	"sync"
)

func main() {
    chA := make(chan int)
    chB := make(chan int)
    chC := make(chan int)

    go func() {
        for i := 1; i <= 3; i++ {
            chA <- i
        }
        close(chA)
    }()
    go func() {
        for i := 11; i <= 13; i++ {
            chB <- i
        }
        close(chB)
    }()
    go func() {
        for i := 21; i <= 23; i++ {
            chC <- i
        }
        close(chC)
    }()

    // Сливаем всё в один канал
    out := merge(chA, chB, chC)

    // Читаем из выходного канала
    for val := range out {
        fmt.Println(val)
    }

}

// Написать код функции, которая делает merge N каналов. 
// Весь входной поток перенаправляется в один канал.
func merge(cs ...<-chan int) <-chan int {
    out := make(chan int)
    wg := sync.WaitGroup{}
    for _, ch := range cs {
        wg.Add(1)
        go func(ch <-chan int) {
            defer wg.Done()
			for val := range ch {
				out <- val
			}
        }(ch)
    }
    go func() {
        wg.Wait()
        close(out)
    }()
    return out
}
```
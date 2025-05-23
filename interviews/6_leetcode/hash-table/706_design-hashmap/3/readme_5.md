Методы **«переместить последний элемент»** и **«ленивое удаление»** — это стратегии эффективного удаления элементов из срезов (`slice`) в языках вроде Go, где срез — это динамический массив с возможностью изменения длины, но без встроенного быстрого удаления из середины. Обычное удаление может требовать сдвига всех последующих элементов, что стоит O(n). Эти методы позволяют **избежать этого**.

---

### ✅ 1. **Метод "переместить последний элемент"** (swap with last)

**Идея**: Если порядок элементов не важен, то при удалении элемента по индексу `i` вы **перемещаете последний элемент на его место** и уменьшаете длину среза на 1.

**Пример (на Go):**

```go
func removeSwapLast(slice []int, i int) []int {
    slice[i] = slice[len(slice)-1] // перемещаем последний элемент на место i
    return slice[:len(slice)-1]    // отрезаем последний элемент
}
```

**Плюсы:**

* Быстро: O(1)
* Не нужно сдвигать элементы

**Минусы:**

* Нарушается порядок элементов

---

### ✅ 2. **Метод "ленивого удаления"** (lazy deletion)

**Идея**: Не удаляете элемент сразу, а просто **помечаете его как удалённый** (например, `nil`, `-1`, `deleted`, `valid=false` и т.д.). Фактическое удаление происходит позже (например, во время «сборки мусора» или «очистки»).

**Пример (на Go):**

```go
type Item struct {
    Key   string
    Value string
    Valid bool
}

func lazyDelete(slice []Item, i int) {
    slice[i].Valid = false
}
```

**Плюсы:**

* O(1) на удаление
* Удобно для частых удалений или итераций с фильтрацией

**Минусы:**

* Память не освобождается сразу
* Нужно потом запускать «очистку» (например, `compact()`), чтобы реально удалить

---

### 🧠 Когда использовать какой метод:

| Сценарий                           | Метод                               |
| ---------------------------------- | ----------------------------------- |
| Порядок важен                      | Стандартное удаление с копированием |
| Порядок не важен                   | `переместить последний элемент`     |
| Частые удаления, но редкие доступы | `ленивое удаление`                  |
| Частые доступы, но редкие удаления | Можно лениво + периодически чистить |

---

Хочешь примеры использования этих методов в хэш-таблице или списке задач?

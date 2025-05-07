Отлично, если **нельзя использовать встроенную `map`**, но при этом **есть идеальная хэш-функция**, то можно улучшить текущую структуру `OurSuperMap` путём замены двух срезов (`key []any`, `value []any`) на **один срез фиксированного размера**, где `индекс = хэш`. То есть делаем **прямую адресацию по хэшу**.

---

### Улучшенная структура без map и без бакетов:

```go
type pair struct {
    key   any
    value any
    valid bool
}

type OurSuperMap struct {
    data []pair
}
```

### Идея:

* `OurSuperHashFunc(key)` возвращает `int64`, который можно привести к индексу в массиве `data`.
* Если хэши в пределах, скажем, `0 <= hash < N`, то можно создать `data := make([]pair, N)`.
* `valid` нужен, чтобы отличать пустую ячейку от занятой.

---

### Реализация (предположим, что хэш возвращает `int` в диапазоне `[0, N)`):

```go
func NewOurSuperMap(size int) OurSuperMap {
    return OurSuperMap{
        data: make([]pair, size),
    }
}

func (m *OurSuperMap) OurSuperHashFunc(key any) int {
    // Предположим: ключ — строка, вернём длину строки как хэш (только как пример)
    if s, ok := key.(string); ok {
        return len(s) % len(m.data) // ограничим в пределах data
    }
    return 0
}

func (m *OurSuperMap) Add(key any, value any) {
    index := m.OurSuperHashFunc(key)
    m.data[index] = pair{key: key, value: value, valid: true}
}

func (m *OurSuperMap) Get(key any) any {
    index := m.OurSuperHashFunc(key)
    p := m.data[index]
    if p.valid {
        return p.value
    }
    return nil
}

func (m *OurSuperMap) Delete(key any) {
    index := m.OurSuperHashFunc(key)
    m.data[index] = pair{} // Обнуляем
}
```

---

### Пример использования:

```go
func main() {
    m := NewOurSuperMap(100)
    m.Add("apple", 10)
    m.Add("banana", 20)

    fmt.Println(m.Get("apple"))  // 10
    fmt.Println(m.Get("banana")) // 20

    m.Delete("apple")
    fmt.Println(m.Get("apple"))  // nil
}
```

---

### Преимущества:

* Всё ещё **без бакетов**.
* **O(1)** вставка и поиск.
* Нет `map` из стандартной библиотеки.
* Массив заранее аллоцирован.

Хочешь, чтобы я показал пример с пользовательской хэш-функцией под твои ключи (например, строки, числа и т. д.)?

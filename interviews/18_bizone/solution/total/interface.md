### Ошибка: **`Manage()` в структуре `SomeManager`**
```go
type SomeManager struct {
    Manage()
    reader    Reader
    processor []Processor
    writer    Writer
}
```
### **Почему это неправильно?**
1. **Методы не объявляются внутри `struct` в Go**  
   - В `struct` можно объявлять **только поля** (переменные), но **не методы**.
   - Запись `Manage()` внутри `SomeManager` выглядит как объявление функции, но в `struct` могут быть **только переменные**.
   - Правильный способ имплементации метода `Manage()` — **через метод `SomeManager`**, а не через поле.

2. **Для имплементации интерфейса `Manager` метод `Manage()` должен быть у `SomeManager` как метод типа**  
   - В Go интерфейс **автоматически реализуется**, если тип (`SomeManager`) **имеет метод с соответствующей сигнатурой**.
   - Указание `Manage()` внутри `struct` ничего не делает и **не добавляет метод в `SomeManager`**.

---

### **Правильный код**
```go
type SomeManager struct {
    reader    Reader
    processor []Processor
    writer    Writer
}

// Реализация метода Manage(), которая делает SomeManager соответствующим интерфейсу Manager
func (m *SomeManager) Manage() {
    for {
        dataBatch := m.reader.Read()
        if len(dataBatch) == 0 {
            continue
        }

        var processedData []*Data
        for _, data := range dataBatch {
            tempData := []*Data{data}
            for _, processor := range m.processor {
                var err error
                tempData, err = processor.Process(*data)
                if err != nil {
                    log.Println(err)
                    tempData = nil
                    break
                }
            }

            if tempData != nil {
                processedData = append(processedData, tempData...)
            }
        }

        if len(processedData) > 0 {
            m.writer.Write(processedData)
        }
    }
}
```

---

### **Как работает интерфейсная имплементация в Go**
```go
type Manager interface {
    Manage() // blocking
}

// SomeManager автоматически реализует Manager, потому что у него есть метод Manage()
var _ Manager = (*SomeManager)(nil) // Проверка соответствия интерфейсу (compile-time check)
```
Если метод `Manage()` отсутствует у `SomeManager`, компилятор выдаст ошибку.

---

### **Вывод**
❌ **Неправильно:**
```go
type SomeManager struct {
    Manage() // ❌ Это невозможно в Go
    reader    Reader
    processor []Processor
    writer    Writer
}
```
✅ **Правильно:**
```go
type SomeManager struct {
    reader    Reader
    processor []Processor
    writer    Writer
}

func (m *SomeManager) Manage() {
    // Реализация
}
```

> В Go **методы определяются снаружи `struct`, а не внутри**.
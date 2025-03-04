### **Описание задачи: 844. Backspace String Compare (Хитрое сравнение строк)**

В этой задаче нужно определить, равны ли две строки после обработки символа `#`, который действует как клавиша "Backspace". Символ `#` удаляет предыдущий символ, если он существует. Если символа для удаления нет, `#` просто игнорируется.

---

### **Пример 1:**
**Ввод:**
```
S = "ab#c", T = "ad#c"
```
**Вывод:**
```
true
```
**Объяснение:**
- После обработки строки `S`: `"ab#c"` превращается в `"ac"`.
- После обработки строки `T`: `"ad#c"` превращается в `"ac"`.
- Обе строки равны.

---

### **Пример 2:**
**Ввод:**
```
S = "ab##", T = "c#d#"
```
**Вывод:**
```
true
```
**Объяснение:**
- После обработки строки `S`: `"ab##"` превращается в `""` (оба символа удаляются).
- После обработки строки `T`: `"c#d#"` превращается в `""` (оба символа удаляются).
- Обе строки равны.

---

### **Пример 3:**
**Ввод:**
```
S = "a#c", T = "b"
```
**Вывод:**
```
false
```
**Объяснение:**
- После обработки строки `S`: `"a#c"` превращается в `"c"`.
- После обработки строки `T`: `"b"` остается `"b"`.
- Строки не равны.

---

### **Ограничения:**
1. Длины строк `S` и `T` находятся в диапазоне `[1, 200]`.
2. Строки содержат только строчные буквы английского алфавита и символ `#`.

---

### **Идея решения**

Мы используем **паттерн двух указателей** (two-pointer technique) - каждому по указателю, чтобы эффективно сравнить строки. Вместо того чтобы обрабатывать строки с начала, мы будем двигаться **с конца**, так как символ `#` удаляет символы, которые идут перед ним. Это позволяет нам избежать создания новых строк или использования дополнительных структур данных, таких как стек.

---

### **Алгоритм**

1. **Два указателя:**  
   - Заводим два указателя: `i` для строки `s` и `j` для строки `t`. Оба указателя начинаются с конца строк (`len(s) - 1` и `len(t) - 1` соответственно).

2. **Обработка символа `#`:**  
   - Если текущий символ — `#`, увеличиваем счетчик удалений (`skip`).
   - Если текущий символ — обычный символ, уменьшаем счетчик удалений. Если счетчик удалений равен нулю, оставляем указатель на этом символе.

3. **Сравнение символов:**  
   - После обработки всех `#` в обеих строках сравниваем символы, на которых остановились указатели. Если символы не равны, возвращаем `false`.

4. **Проверка окончания строк:**  
   - После завершения основного цикла оба указателя должны быть равны `-1`. Если один из указателей не дошел до конца строки, это означает, что строки не равны.

---

### **Пример решения**

Рассмотрим пример:

**Ввод:**  
`s = "ac#b#ac"`  
`t = "abc##aa#b#c"`

**Результат:**  
`true`

**Ход решения:**

1. **Инициализация:**  
   - Указатель `i` на последнем символе строки `s` (`i = 6`, `s[i] = 'c'`).
   - Указатель `j` на последнем символе строки `t` (`j = 10`, `t[j] = 'c'`).

2. **Сравнение символов:**  
   - `s[i] = 'c'` и `t[j] = 'c'`. Символы равны. Сдвигаем указатели: `i = 5`, `j = 9`.

3. **Обработка `#` в строке `s`:**  
   - `s[i] = '#'`. Увеличиваем счетчик удалений: `skipS = 1`. Сдвигаем указатель: `i = 4`.
   - `s[i] = 'a'`. Уменьшаем счетчик удалений: `skipS = 0`. Сдвигаем указатель: `i = 3`.

4. **Обработка `#` в строке `t`:**  
   - `t[j] = '#'`. Увеличиваем счетчик удалений: `skipT = 1`. Сдвигаем указатель: `j = 8`.
   - `t[j] = 'b'`. Уменьшаем счетчик удалений: `skipT = 0`. Сдвигаем указатель: `j = 7`.

5. **Сравнение символов:**  
   - `s[i] = 'b'` и `t[j] = 'b'`. Символы равны. Сдвигаем указатели: `i = 2`, `j = 6`.

6. **Обработка `#` в строке `s`:**  
   - `s[i] = '#'`. Увеличиваем счетчик удалений: `skipS = 1`. Сдвигаем указатель: `i = 1`.
   - `s[i] = 'c'`. Уменьшаем счетчик удалений: `skipS = 0`. Сдвигаем указатель: `i = 0`.

7. **Обработка `#` в строке `t`:**  
   - `t[j] = '#'`. Увеличиваем счетчик удалений: `skipT = 1`. Сдвигаем указатель: `j = 5`.
   - `t[j] = 'a'`. Уменьшаем счетчик удалений: `skipT = 0`. Сдвигаем указатель: `j = 4`.

8. **Сравнение символов:**  
   - `s[i] = 'a'` и `t[j] = 'a'`. Символы равны. Сдвигаем указатели: `i = -1`, `j = -1`.

9. **Проверка окончания строк:**  
   - Оба указателя равны `-1`. Строки равны.

**Результат:**  
`true`

---

### **Другой пример**

**Ввод:**  
`s = "a"`  
`t = "aaaa###a"`

**Результат:**  
`false`

**Ход решения:**

1. **Инициализация:**  
   - Указатель `i` на последнем символе строки `s` (`i = 0`, `s[i] = 'a'`).
   - Указатель `j` на последнем символе строки `t` (`j = 7`, `t[j] = 'a'`).

2. **Сравнение символов:**  
   - `s[i] = 'a'` и `t[j] = 'a'`. Символы равны. Сдвигаем указатели: `i = -1`, `j = 6`.

3. **Обработка `#` в строке `t`:**  
   - `t[j] = '#'`. Увеличиваем счетчик удалений: `skipT = 1`. Сдвигаем указатель: `j = 5`.
   - `t[j] = '#'`. Увеличиваем счетчик удалений: `skipT = 2`. Сдвигаем указатель: `j = 4`.
   - `t[j] = '#'`. Увеличиваем счетчик удалений: `skipT = 3`. Сдвигаем указатель: `j = 3`.
   - `t[j] = 'a'`. Уменьшаем счетчик удалений: `skipT = 2`. Сдвигаем указатель: `j = 2`.
   - `t[j] = 'a'`. Уменьшаем счетчик удалений: `skipT = 1`. Сдвигаем указатель: `j = 1`.
   - `t[j] = 'a'`. Уменьшаем счетчик удалений: `skipT = 0`. Сдвигаем указатель: `j = 0`.

4. **Проверка окончания строк:**  
   - Указатель `i = -1`, но указатель `j = 0`. Это означает, что строка `t` содержит лишние символы.  

**Результат:**  
`false`

---

### **Сложность**

1. **Временная сложность:**  
   - Мы проходим по каждой строке один раз, обрабатывая каждый символ.  
   - Итоговая сложность: **O(n + m)**, где `n` — длина строки `s`, а `m` — длина строки `t`.
    - Почему у нас не может быть O(n)? У нас же первая строка может быть примерно равна длине второй строки, а это O(2n), где мы можем округлить до O(n).
    - У нам может быть пример s = a, t = x1000'a'x1000'#'a. Мы видим, что вторая строка очень большая, в таком случае мы не можем игнорировать вторую строку.

2. **Сложность по памяти:**  
   - Мы используем только указатели и счетчики, без дополнительных структур данных.  
   - Итоговая сложность: **O(1)**.

---

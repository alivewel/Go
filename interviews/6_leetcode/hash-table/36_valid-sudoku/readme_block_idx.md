Формула `blockIdx = i / 3 * 3 + j / 3` используется для вычисления индекса 3x3 блока в судоку. Давайте разберем, зачем она нужна и как она работает.

---

### **Зачем вычислять индекс блока?**
В задаче **Valid Sudoku** необходимо проверить, что каждая из девяти 3x3 подрешеток (блоков) содержит только уникальные цифры. Для этого нужно каким-то образом идентифицировать, к какому блоку принадлежит каждая ячейка `(i, j)`.

Поскольку блоки в судоку расположены в виде сетки 3x3, мы можем пронумеровать их от 0 до 8 следующим образом:

```
0 | 1 | 2
---------
3 | 4 | 5
---------
6 | 7 | 8
```

Формула `blockIdx = i / 3 * 3 + j / 3` позволяет однозначно определить индекс блока для любой ячейки `(i, j)`.

---

### **Как работает формула?**
1. **`i / 3`:**  
   - Деление строки `i` на 3 (целочисленное деление) определяет, в какой горизонтальной группе блоков находится ячейка.  
   - Например:
     - Строки `0, 1, 2` → группа 0 (верхняя).
     - Строки `3, 4, 5` → группа 1 (средняя).
     - Строки `6, 7, 8` → группа 2 (нижняя).

2. **`j / 3`:**  
   - Деление столбца `j` на 3 определяет, в какой вертикальной группе блоков находится ячейка.  
   - Например:
     - Столбцы `0, 1, 2` → группа 0 (левая).
     - Столбцы `3, 4, 5` → группа 1 (средняя).
     - Столбцы `6, 7, 8` → группа 2 (правая).

3. **`i / 3 * 3`:**  
   - Умножение на 3 смещает индекс в зависимости от горизонтальной группы блоков.  
   - Например:
     - Для верхней группы (где `i / 3 = 0`), смещение будет `0 * 3 = 0`.
     - Для средней группы (где `i / 3 = 1`), смещение будет `1 * 3 = 3`.
     - Для нижней группы (где `i / 3 = 2`), смещение будет `2 * 3 = 6`.

4. **`i / 3 * 3 + j / 3`:**  
   - Суммируя горизонтальное смещение (`i / 3 * 3`) и вертикальный индекс (`j / 3`), мы получаем уникальный индекс блока.  
   - Например:
     - Ячейка `(0, 0)` → `0 / 3 * 3 + 0 / 3 = 0` (верхний левый блок).
     - Ячейка `(4, 4)` → `4 / 3 * 3 + 4 / 3 = 1 * 3 + 1 = 4` (центральный блок).
     - Ячейка `(7, 8)` → `7 / 3 * 3 + 8 / 3 = 2 * 3 + 2 = 8` (нижний правый блок).

---

### **Почему это удобно?**
Формула позволяет:
1. Однозначно идентифицировать блок для любой ячейки `(i, j)`.
2. Использовать массив или словарь для проверки уникальности цифр в каждом блоке. Например:
   - Можно создать массив `blocks[9]`, где каждый элемент — это набор (`set`) для хранения цифр текущего блока.
   - При обработке ячейки `(i, j)` мы просто добавляем её значение в `blocks[blockIdx]` и проверяем на повторения.

---

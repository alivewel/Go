### **Описание задачи: 283. Move Zeroes (Нули в конец массива)**

Вам дан массив целых чисел `nums`. Необходимо переместить все нули в конец массива, сохраняя **относительный порядок ненулевых элементов**. Это нужно сделать **на месте**, то есть без создания копии массива.

---

### **Условия задачи**
1. Все ненулевые элементы должны сохранять свой порядок.
2. Все нули должны быть перемещены в конец массива.
3. Операция должна быть выполнена **на месте**, без использования дополнительного массива.

---

### **Примеры**

#### **Пример 1:**
**Ввод:**  
`nums = [0, 1, 0, 3, 12]`  
**Вывод:**  
`[1, 3, 12, 0, 0]`  
**Объяснение:**  
- Нули перемещены в конец, а порядок ненулевых элементов сохранен.

#### **Пример 2:**
**Ввод:**  
`nums = [0]`  
**Вывод:**  
`[0]`  
**Объяснение:**  
- Массив уже содержит только один элемент, который равен нулю, поэтому изменений не требуется.

---

### **Алгоритм решения**

Для решения задачи можно использовать **двухуказательный подход** (two-pointer technique). Этот метод позволяет эффективно перемещать элементы без использования дополнительной памяти.

#### **Шаги алгоритма:**
1. Заводим два указателя:
   - `i` — текущий индекс, по которому мы проходим массив.
   - `lastNonZeroIndex` — индекс, на который нужно переместить следующий ненулевой элемент.
2. Проходим по массиву:
   - Если текущий элемент `nums[i]` ненулевой, перемещаем его на позицию `lastNonZeroIndex` и увеличиваем `lastNonZeroIndex`.
3. После завершения прохода по массиву все ненулевые элементы будут перемещены в начало массива. Оставшиеся элементы заполняются нулями.

---

Для решения задачи будем использовать два указателя: первый — медленный, второй — быстрый. Быстрый указатель будет указывать на первый ненулевой элемент, а медленный — на место, куда нужно переместить этот ненулевой элемент.

### Пример 1: `nums = [0, 2, 0, 0, 4, 0]`

1. **Инициализация**: Быстрый указатель указывает на первый элемент массива (nums[0] = 0). Медленный указатель также указывает на первый элемент.

2. **Проход по массиву**:
   - **nums[0] = 0**: Медленный указатель указывает на нулевой элемент. Быстрый указатель ищет первый ненулевой элемент и находит его в `nums[1] = 2`. Присваиваем `nums[0] = 2`. Теперь следующее ненулевое число мы запишем на место медленного указателя.
   - **nums[1] = 0**: Медленный указатель указывает на `nums[1]`. Быстрый указатель продолжает искать и находит `nums[4] = 4`. Присваиваем `nums[1] = 4`.
   - **nums[2] = 0**: Медленный указатель указывает на `nums[2]`. Быстрый указатель дошел до последнего элемента массива. Первая часть алгоритма завершена.

3. **Завершение**: Теперь нам нужно пройтись от элемента, на который указывает медленный указатель, и проставить каждому следующему элементу нули. В результате получаем: `nums = [2, 4, 0, 0, 0, 0]`.

### Пример 2: `nums = [1, 3, 2, 0, 4, 5]`

1. **Инициализация**: Быстрый указатель указывает на первый элемент массива (nums[0] = 1). Медленный указатель также указывает на первый элемент.

2. **Проход по массиву**:
   - **nums[0] = 1**: Медленный указатель стоит на элементе `nums[0]`, быстрый указатель тоже. Мы присваиваем `nums[0] = nums[0]` (ничего не меняется).
   - **nums[1] = 3**: Медленный указатель стоит на элементе `nums[1]`, быстрый указатель тоже. Мы присваиваем `nums[1] = nums[1]` (ничего не меняется).
   - **nums[2] = 2**: Медленный указатель стоит на элементе `nums[2]`, быстрый указатель тоже. Мы присваиваем `nums[2] = nums[2]` (ничего не меняется).
   - **nums[3] = 0**: Медленный указатель стоит на элементе `nums[3]`. Теперь быстрый указатель ищет ненулевое значение и находит его в `nums[4] = 4`. Присваиваем `nums[3] = nums[4]`.
   - **nums[4] = 0**: Медленный указатель стоит на элементе `nums[4]`. Быстрый указатель ищет ненулевое значение и находит его в `nums[5] = 5`. Присваиваем `nums[4] = nums[5]`.

3. **Завершение**: Быстрый указатель выходит за пределы массива. Теперь нам нужно пройтись от элемента, на который указывает медленный указатель, и проставить каждому следующему элементу нули. В результате получаем: `nums = [1, 3, 2, 4, 5, 0]`.

### Временная и пространственная сложность

- **Временная сложность**: O(n). В самом худшем случае мы проходим по всем элементам дважды: один раз с медленным указателем и второй раз с быстрым указателем.
- **Память**: O(1). Мы не создаем дополнительной памяти, так как работаем с массивом на месте.

### Частая ошибка

Кандидаты часто забывают занулить последние элементы массива, и результат получается неправильным, например: `nums = [1, 3, 2, 4, 5, 5]`.

### Похожие задачи

Существует похожая задача с аналогичным решением: удалить все единицы из массива. В этом случае единицы перетаскиваются в самый конец, и массив переаллокации не требуется. В Go можно вернуть срез от среза.

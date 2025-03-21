### Алгоритм решения задачи **3Sum** с использованием мапы (хэш-таблицы)

Решение задачи **3Sum** с помощью мапы (хэш-таблицы) основывается на разбиении задачи на подзадачи, аналогичные задаче **2Sum**. Вот пошаговый алгоритм:

---

### 1. **Идея**
- Задача состоит в том, чтобы найти все уникальные тройки чисел `(nums[i], nums[j], nums[k])` в массиве `nums`, такие что их сумма равна нулю: `nums[i] + nums[j] + nums[k] == 0`.
- Для этого мы фиксируем одно число из тройки, а оставшиеся два числа ищем с помощью подхода, похожего на задачу **2Sum**.

---

### 2. **Шаги алгоритма**

#### Шаг 1: Сортировка массива
- Сначала отсортируйте массив `nums`. Это упростит проверку уникальности троек и позволит избежать дубликатов.
- Пример: если `nums = [-1, 0, 1, 2, -1, -4]`, то после сортировки: `nums = [-4, -1, -1, 0, 1, 2]`.

#### Шаг 2: Фиксация первого числа
- Используйте цикл для перебора каждого элемента массива как первого числа тройки (`nums[i]`).
- Пропускайте дубликаты: если `nums[i] == nums[i-1]` (и `i > 0`), переходите к следующему элементу, чтобы избежать повторяющихся троек.

#### Шаг 3: Преобразование задачи в 2Sum
- Для каждого фиксированного числа `nums[i]` преобразуйте задачу в поиск двух чисел, сумма которых равна `-nums[i]` (чтобы в сумме с `nums[i]` получилось 0).
- Целевая сумма для двух чисел: `target = -nums[i]`.

#### Шаг 4: Использование мапы для поиска двух чисел
- Создайте пустую мапу (хэш-таблицу), чтобы отслеживать числа, которые вы уже видели.
- Пройдите по оставшимся элементам массива (начиная с индекса `i+1`):
  - Для каждого числа `nums[j]` вычислите его "комплемент": `complement = target - nums[j]`.
  - Проверьте, есть ли `complement` в мапе:
    - Если `complement` уже есть в мапе, значит, вы нашли тройку: `(nums[i], nums[j], complement)`.
    - Добавьте эту тройку в результат, избегая дубликатов.
  - Добавьте текущее число `nums[j]` в мапу, чтобы оно могло быть использовано как "комплемент" для следующих чисел.

#### Шаг 5: Уникальность троек
- Чтобы избежать дубликатов, можно:
  - Сортировать каждую найденную тройку перед добавлением в результат.
  - Использовать структуру данных, которая автоматически удаляет дубликаты (например, `set` в Python или `map` в Go).

---

### 3. **Пример работы алгоритма**

#### Входные данные:
```plaintext
nums = [-1, 0, 1, 2, -1, -4]
```

#### Шаги:
1. **Сортировка массива:**
   ```plaintext
   nums = [-4, -1, -1, 0, 1, 2]
   ```

2. **Фиксация первого числа:**
   - `i = 0`, `nums[i] = -4`, `target = 4`.

3. **Поиск двух чисел с суммой 4:**
   - Используем мапу:
     - `j = 1`, `nums[j] = -1`, `complement = 4 - (-1) = 5`. Добавляем `-1` в мапу.
     - `j = 2`, `nums[j] = -1`, `complement = 4 - (-1) = 5`. Добавляем `-1` в мапу.
     - `j = 3`, `nums[j] = 0`, `complement = 4 - 0 = 4`. Добавляем `0` в мапу.
     - `j = 4`, `nums[j] = 1`, `complement = 4 - 1 = 3`. Добавляем `1` в мапу.
     - `j = 5`, `nums[j] = 2`, `complement = 4 - 2 = 2`. Добавляем `2` в мапу.

4. **Следующее фиксированное число:**
   - `i = 1`, `nums[i] = -1`, `target = 1`.
   - Используем мапу:
     - `j = 2`, `nums[j] = -1`, `complement = 1 - (-1) = 2`. Добавляем `-1` в мапу.
     - `j = 3`, `nums[j] = 0`, `complement = 1 - 0 = 1`. Добавляем `0` в мапу.
     - `j = 4`, `nums[j] = 1`, `complement = 1 - 1 = 0`. Найдена тройка: `(-1, 1, 0)`.

5. **Продолжаем для всех элементов.**

---

### 4. **Сложность алгоритма**

- **Временная сложность:**
  - Внешний цикл по `i` выполняется `O(n)` раз.
  - Внутренний цикл с использованием мапы выполняется `O(n)` раз для каждого `i`.
  - Итоговая сложность: `O(n^2)`.

- **Пространственная сложность:**
  - Используется мапа для хранения чисел, что требует `O(n)` памяти.

---

### 5. **Преимущества и недостатки**

#### Преимущества:
- Простота реализации.
- Эффективность по сравнению с брутфорс-методом.

#### Недостатки:
- Использование дополнительной памяти для мапы.
- Может быть менее эффективным, чем метод с сортировкой и двумя указателями, который не требует дополнительной памяти.
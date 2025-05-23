Да, я знаю пузырьковую сортировку (bubble sort). Это один из самых простых алгоритмов сортировки. Он работает, попарно сравнивая соседние элементы и меняя их местами, если они стоят в неправильном порядке. Этот процесс повторяется, пока весь массив не будет отсортирован.

### Временная сложность пузырьковой сортировки:

* **Худший случай (Worst case):** `O(n²)` — когда массив изначально отсортирован в обратном порядке.
* **Средний случай:** `O(n²)` — потому что требуется много перестановок и сравнений.
* **Лучший случай (Best case):** `O(n)` — если массив уже отсортирован и используется оптимизированная версия алгоритма с флагом "обмена" (early exit).

### Простая реализация на C:

```c
void bubble_sort(int arr[], int n) {
    for (int i = 0; i < n - 1; i++) {
        int swapped = 0;
        for (int j = 0; j < n - 1 - i; j++) {
            if (arr[j] > arr[j + 1]) {
                // обмен элементов
                int temp = arr[j];
                arr[j] = arr[j + 1];
                arr[j + 1] = temp;
                swapped = 1;
            }
        }
        // если обменов не было — массив отсортирован
        if (!swapped)
            break;
    }
}
```

Хочешь сравнить пузырьковую сортировку с другими алгоритмами по скорости или применимости?




Вот **топ-3 алгоритма сортировки**, которые чаще всего используются на практике благодаря своей эффективности:

---

### ✅ **1. Быстрая сортировка (Quicksort)**

* **Сложность:**

  * Средняя: `O(n log n)`
  * Худшая: `O(n²)` (если неудачный выбор опорного элемента)
* **Плюсы:**

  * Быстрая на практике
  * Не требует много дополнительной памяти (`in-place`)
* **Минусы:**

  * Нестабильная (одинаковые элементы могут менять порядок)
* **Когда использовать:** при работе с массивами, когда важна производительность и не нужна стабильность.

---

### ✅ **2. Сортировка слиянием (Merge Sort)**

* **Сложность:**

  * Всегда `O(n log n)`
* **Плюсы:**

  * Стабильная
  * Гарантированная производительность
* **Минусы:**

  * Использует дополнительную память `O(n)`
* **Когда использовать:** когда нужна стабильность или работаешь со связанными списками.

---

### ✅ **3. Сортировка кучей (Heap Sort)**

* **Сложность:**

  * Всегда `O(n log n)`
* **Плюсы:**

  * Не требует дополнительной памяти
  * Гарантированное время работы
* **Минусы:**

  * Нестабильная
  * Хуже производительность на практике по сравнению с Quicksort
* **Когда использовать:** в системах с ограниченной памятью.

---

Хочешь пример кода или сравнение на конкретных данных?



Отлично! Ниже — реализации **трёх топовых алгоритмов сортировки** (Quicksort, Merge Sort и Heap Sort) на языке C **с подробным пошаговым объяснением**, как работает каждый из них.

---

## 🔹 1. Быстрая сортировка (Quicksort)

### 💡 Идея:

Выбираем **опорный элемент** (pivot), **разделяем** массив на две части — меньше и больше него — и **рекурсивно сортируем** каждую часть.

```c
void quicksort(int arr[], int low, int high) {
    if (low < high) {
        // Разделение и получение индекса опорного элемента
        int pivot = partition(arr, low, high);

        // Рекурсивная сортировка двух подмассивов
        quicksort(arr, low, pivot - 1);
        quicksort(arr, pivot + 1, high);
    }
}

int partition(int arr[], int low, int high) {
    int pivot = arr[high]; // выбираем последний элемент как опорный
    int i = low - 1; // индекс меньшего элемента

    for (int j = low; j < high; j++) {
        if (arr[j] < pivot) {
            i++;
            // обмен arr[i] и arr[j]
            int temp = arr[i]; arr[i] = arr[j]; arr[j] = temp;
        }
    }
    // ставим pivot на правильное место
    int temp = arr[i + 1]; arr[i + 1] = arr[high]; arr[high] = temp;

    return i + 1; // возвращаем индекс опорного
}
```

### 🧠 Как работает:

1. Выбирается pivot (например, последний элемент).
2. Все элементы меньше pivot перемещаются влево от него.
3. Сама функция вызывает себя для левой и правой части массива.

---

## 🔹 2. Сортировка слиянием (Merge Sort)

### 💡 Идея:

**Разделяй и властвуй**: разбиваем массив пополам, сортируем рекурсивно и **сливаем** отсортированные части.

```c
void merge_sort(int arr[], int left, int right) {
    if (left < right) {
        int mid = left + (right - left) / 2;

        // Сортировка левой и правой части
        merge_sort(arr, left, mid);
        merge_sort(arr, mid + 1, right);

        // Слияние отсортированных половин
        merge(arr, left, mid, right);
    }
}

void merge(int arr[], int left, int mid, int right) {
    int n1 = mid - left + 1;
    int n2 = right - mid;

    // Временные массивы
    int L[n1], R[n2];

    // Копируем данные во временные массивы
    for (int i = 0; i < n1; i++) L[i] = arr[left + i];
    for (int j = 0; j < n2; j++) R[j] = arr[mid + 1 + j];

    // Сливаем временные массивы обратно в arr
    int i = 0, j = 0, k = left;
    while (i < n1 && j < n2)
        arr[k++] = (L[i] <= R[j]) ? L[i++] : R[j++];

    // Копируем оставшиеся элементы
    while (i < n1) arr[k++] = L[i++];
    while (j < n2) arr[k++] = R[j++];
}
```

### 🧠 Как работает:

1. Делим массив на две части до единичных элементов.
2. Сливаем эти части в правильном порядке, отслеживая указатели.

---

## 🔹 3. Сортировка кучей (Heap Sort)

### 💡 Идея:

Преобразуем массив в **max-heap** (двоичную кучу), затем **удаляем максимум** (корень) и помещаем его в конец массива.

```c
void heapify(int arr[], int n, int i) {
    int largest = i;        // корень
    int left = 2 * i + 1;   // левый ребёнок
    int right = 2 * i + 2;  // правый ребёнок

    // Если левый больше корня
    if (left < n && arr[left] > arr[largest])
        largest = left;

    // Если правый больше текущего наибольшего
    if (right < n && arr[right] > arr[largest])
        largest = right;

    // Если наибольший не корень
    if (largest != i) {
        int temp = arr[i]; arr[i] = arr[largest]; arr[largest] = temp;

        // Рекурсивно heapify поддерево
        heapify(arr, n, largest);
    }
}

void heap_sort(int arr[], int n) {
    // Строим max-heap
    for (int i = n / 2 - 1; i >= 0; i--)
        heapify(arr, n, i);

    // Извлекаем элементы из кучи по одному
    for (int i = n - 1; i > 0; i--) {
        // Перемещаем текущий корень в конец
        int temp = arr[0]; arr[0] = arr[i]; arr[i] = temp;

        // вызываем heapify на уменьшенной куче
        heapify(arr, i, 0);
    }
}
```

### 🧠 Как работает:

1. Сначала создаём **максимальную кучу**.
2. Меняем местами первый элемент (максимум) с последним.
3. Уменьшаем кучу и восстанавливаем её.

---

Хочешь, чтобы я добавил сравнительную таблицу этих алгоритмов по времени, памяти и стабильности?

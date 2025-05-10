``` C
#include <stdio.h>

// Функция для обмена двух элементов
void swap(int *a, int *b) {
    int temp = *a;
    *a = *b;
    *b = temp;
}

// Основная функция быстрой сортировки
void quicksort(int arr[], int low, int high) {
    if (low < high) {
        // Разделяем массив и получаем индекс опорного элемента
        int pivotIndex = partition(arr, low, high);

        // Рекурсивно сортируем левую часть массива
        quicksort(arr, low, pivotIndex - 1);

        // Рекурсивно сортируем правую часть массива
        quicksort(arr, pivotIndex + 1, high);
    }
}

// Функция разбиения массива (partitioning)
int partition(int arr[], int low, int high) {
    // Выбираем опорный элемент (здесь — последний элемент массива)
    int pivot = arr[high];

    // i будет указывать на "границу" меньших элементов
    int i = low - 1;

    // Перебираем все элементы от low до high - 1
    for (int j = low; j < high; j++) {
        // Если текущий элемент меньше pivot,
        // перемещаем его в "левую часть"
        if (arr[j] < pivot) {
            i++; // увеличиваем границу

            // меняем местами arr[i] и arr[j]
            swap(&arr[i], &arr[j]);
        }
    }

    // Ставим опорный элемент на своё правильное место (после меньших элементов)
    swap(&arr[i + 1], &arr[high]);

    // Возвращаем индекс опорного элемента
    return i + 1;
}
```
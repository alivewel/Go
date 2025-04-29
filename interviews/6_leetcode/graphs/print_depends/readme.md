**Печать зависимостей в порядке импорта**

**Условие**: Есть мапа библиотек и их зависимостей. Нужно распечатать библиотеки в порядке правильного импорта. 
Неправильный порядок — если печатаешь библиотеку, а её зависимость ещё не напечатана. Все остальные порядки правильные.

NB: Циклических завистимостей во входных данных нет.

Входные параметры: Глобальная мапа зависимостей  

Выходные параметры: Печать в stdout

**Пример**:
```
# Мапа:
deps = {
  "tensorflow": ["nvcc", "gpu", "linux"],
  "nvcc": ["linux"],
  "linux": ["core"],
  "mylib": ["tensorflow"],
  "mylib2": ["requests"]
}

# Один из правильных порядков:
core linux nvcc gpu tensorflow mylib requests mylib2

```

time: O(N), N - число вершин в графе
mem: O(N), N - число вершин в графе

To do: Проверить у чата оценку

deps = {
  "tensorflow": ["nvcc", "gpu", "linux"],
  "nvcc": ["linux"],
  "linux": ["core"],
  "mylib": ["tensorflow"],
  "mylib2": ["requests"]
}


deps = {
  "abc": ["abc", "bca", "qwe", "wer"],
  "bca": ["abc", "bca", "qwe"],
  "qwe": ["abc", "bca", "qwe"],
  "mylib": ["abc", "bca", "qwe"],
  "mylib2": ["abc", "bca", "qwe"]
}

func traversalDepend(lib string, deps map[string][string], visited map[string]struct{}, res *[]string) {
    if _, found := visited[lib]; found {
        return
    }
    visited[lib] = struct{}{}
    for _, dep := range deps[lib] {
        traversalDepend(dep, deps, visited, res)
    }
    *res = append(*res, lib)
} 

func checkDepend(deps map[string][string] ) []string {
    res := make([]string, 0)
    visited := make(map[string]struct{}, 0)
    for lib, _ := range deps {
        traversalDepend(lib, deps, visited, &res)
    }
    return res
}


| Что оцениваем | Оценка     |
|---------------|------------|
| Время         | O(V + E)   |
| Память        | O(V)       |

где:
V — количество библиотек (вершин графа),
E — общее количество зависимостей (ребер графа).

Оценка по памяти (space complexity)
Что мы храним в памяти?

visited — для каждой вершины один ключ → занимает O(V).

res — список результата, туда добавляем все библиотеки → тоже O(V).

Рекурсивный стек вызовов:

В худшем случае глубина стека равна числу вершин V (если зависимость — длинная цепочка).

То есть стек рекурсии — O(V).

Итого по памяти: O(V) на структуру visited, O(V) на массив результата, O(V) стек вызовов.

Суммарная память:
✅ O(V).

Решение задачи сводиться к формуле `maxX + minX - x`

Эта формула используется для нахождения симметричной точки относительно вертикальной линии симметрии. Непонятно как ее найти, запомнить ее нереально. Нужно постараться вывести ее на интервью, когда попадется эта задача.

1. **Линия симметрии**:
   - Линия симметрии для множества точек определяется как вертикальная линия, которая проходит через середину между минимальной и максимальной координатами по оси X. 
   - Координата этой линии симметрии равна `(maxX + minX) / 2`.

2. **Расстояние от точки `x` до линии симметрии**:
   - Расстояние от точки `x` до линии симметрии равно `x - (maxX + minX) / 2`.

3. **Симметричная координата**:
```
Симметричная координата = Линия симметрии - Расстояние от точки `x` до линии симметрии
```

```
Симметричная координата = (maxX + minX) / 2 - (x - (maxX + minX) / 2)
                        = maxX + minX - x
```

Здесь:
- `(maxX + minX) / 2` — это координата линии симметрии (средняя точка между `maxX` и `minX`).
- `x` — координата текущей точки.
- `(x - (maxX + minX) / 2)` — это расстояние от точки `x` до линии симметрии.
- Формула вычитает это расстояние из линии симметрии, чтобы получить симметричную координату.

#### **Шаг 1: Раскрытие скобок**
Раскроем скобки во второй части формулы:

```
Симметричная координата = (maxX + minX) / 2 - x + (maxX + minX) / 2
```

Здесь:
- `(maxX + minX) / 2` из первой части остается без изменений.
- `-(x - (maxX + minX) / 2)` раскрывается как `-x + (maxX + minX) / 2`.

#### **Шаг 2: Сложение одинаковых частей**
Теперь сложим два `(maxX + minX) / 2`:

```
Симметричная координата = (maxX + minX) / 2 + (maxX + minX) / 2 - x
```

Сложение двух одинаковых дробей дает:

```
Симметричная координата = (maxX + minX) - x
```
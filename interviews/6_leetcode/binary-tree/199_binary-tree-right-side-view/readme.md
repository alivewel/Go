50:00

У нас есть бинарное дерево. Нам нужно на каждом уровне вернуть крайний правый элемент. Вернуть массив этих элементов. 

Кажется, что в решении задачи можно пройтись по всем правым веткам. Только у нас может быть ситуация, при которой крайний правый элемент будет в левом поддереве.

Для решения задачи можно использовать level order (задача №102), только сохранять последний элемент. И возвращать одномерный массив. При проходе запоминать самую последнюю вершину. Мы постоянно затираем элемент в массиве, в конце останется крайний правый.

Оценка памяти и времени:
time: O(n), где n - количество элементов в дереве
mem: O(n) - оценка сверху, в котором мы не учитываем константы
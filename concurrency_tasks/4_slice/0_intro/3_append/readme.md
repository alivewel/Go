Функция append

Если при вставке элемента мы упремся в capasity, будет создан новый базовый массив, в который будут скопирован старый массив и вставится новый элемент и в структуре SliceHeader будет изменена ссылка на базовый массив. Старый массив останется или будет удален сборщиком мусора.
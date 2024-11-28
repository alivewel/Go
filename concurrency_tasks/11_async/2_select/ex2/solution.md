-4:00:00

Внутри select за один раз выполняется только одна операция. select не опрашивает все каналы и не выполняет все. Если мы хотим делать дополнительные попытки опроса канала нам нужно добавить бесконечный цикл.

Создадим 2 канала и запишем туда некоторые значения. В цикле будет периодечески опрашивать эти каналы и если в каналах что-то есть будет выводить значение на экран. Если ничего не вернеться завершим цикл, используя брейк и метку цикла (LOOP).

При запуске программы мы видим, что порядок вывода у нас всегда разный. Какой из каналов выбрать для чтения планировщик Go выбирает случайным образом, мы это не контролируем. Они выполняются не в том порядке, в котором мы указали внутри select.
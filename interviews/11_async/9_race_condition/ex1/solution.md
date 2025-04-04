### Race_condition

В этом разделе поговорим о состоянии гонки. Оно можем быть вам знакомо, если вы занимались БД и изучали транзакции, которые нужны, чтобы атомарно обновить какое-то значение в БД.

В этом примере мы будем мапу со счетчиками из разных горутин. На первый взгляд, если мы посмотрим на код, то ошибок в нем не видны. Мы обновляем разные значения, которые не пересекаются друг с другом, но результат работы этой программы нас может расстроить. Мапа в го конкурентно не безопасный тип данных. Обращаясь к ним обращаясь из разных горутин, которые могут выполняться совершенно на разных процессах можно словить состояние гонки. 

Каким образом это может произойти? Мапа внутри себя - это ссылка на структуру данных, как хеш-таблица и она может перестраиваться, копироваться в другое место, ну или еще что-то может с ней происходить. При этом в кэше одного процессора может лежать одно значение, в кэше другого процессора может лежать другое значение и при попытке обновить значение произойдет коллапс.

При попытке запустить программу мы видим сообщение "concurrent map read and map write". Также можно заметить, что компилятор ругается на строку "counters[th*10+j]++", при попытке записи в мапу.

Тут может возникнуть вопрос Как отлавливать такие ошибки? Эти ошибки нетревиальные и не всегда воспроизводятся. Это может называться плавающим багом или "гейзенбагом". Для отлавливания таких ошибок в Go используется специальная деректива при запуске программы или компиляции под названием `-race`. При запуске программы нужно использовать следующую командку `go run -race main.go`. При использовании данной директивы мы видим другой вывод, все наши счечики вывелись на экран (так происходит не при каждом запуске) и в конце вывода можно заметить сообщение `Found 4 data race(s)`, найдено 4 состояния гонки. Все они происходят на строке `counters[th*10+j]++`.

В следующим примере рассмотрим, как бороться с состоянием гонки.

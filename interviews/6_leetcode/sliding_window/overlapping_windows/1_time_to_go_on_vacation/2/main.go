package main 

import (
    "fmt"
    "math" 
)

/*
**Пора в отпуск**

**Условие**: В жизни каждого приходит время, когда надо отправиться в отпуск. 
Но мы смотрим на календарь и понимаем, что впереди куча встреч, которые не хочется пропустить.  
Необходимо определить день X начала отпуска длиной в V дней так, чтобы отгулять весь отпуск в 
ближайшие P дней и пропустить минимум Y встреч. Считаем, что уже завтра первый возможный день отпуска (day:1).

**Пример**:
```
daysWithMeetings = [
  { day: 3, meetings: 1 },
  { day: 4, meetings: 3 },
  { day: 14, meetings: 3 },
  { day: 21, meetings: 3 },
  { day: 28, meetings: 1 },
];

periodLength = 30 # В какой срок надо отгулять весь отпуск. В данном примере в ближайшие 30 дней
vacationLength = 7

# Вывод
[A, B] # [A -- день X начала отпуска, B -- сколько встреч у пропустим]
p.s. в выводе вместо A и B должны быть целые числа, но условие не получилось полностью сохранить, 
поэтому я не могу точно сказать сейчас, какие там должны быть цифры в ответе
```
Алгоритмическая сложность
- Ожидаемая сложность по времени O(число дней, в течение которых должен пройти отпуск)
- Ожидаемая сложность по памяти O(N)
time O(periodLength)
space O(1)
*/

func bestTimeVacation(daysWithMeetings map[int]int, vacationLength, periodLength int) []int {
    // l, r := 0, -1
    r := -1
    res := make([]int, 0)
    curCountMeetings, minCountMeetings := 0, math.MaxInt
    for l := range periodLength {
        for l - r + 1 == vacationLength && r + 1 < periodLength {
            if countMeetings, found := daysWithMeetings[r + 1]; found {
                curCountMeetings += countMeetings
            }
            r++
        }
        if curCountMeetings < minCountMeetings {
            minCountMeetings = curCountMeetings
            res = res[:0]
            res = append(res, l+1, r+1)
        }
        if countMeetings, found := daysWithMeetings[l + 1]; found {
            curCountMeetings -= countMeetings
        } 
    }
    return res
}

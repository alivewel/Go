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
iteratorR = 0
iteratorL = 0
l = 14
r = 8
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

/**

period 9
vacation 5  
r < vacationLength 


l = 2
r = 6

iteratorL = 0
iteratorR = 3

sumSlitingWindow = 5 + 3 + 2 - 5 + 3

x = 4
y = 6

minMitings = search 
X = (l) 

daysWithMeetings = [
  { day: 3, meetings: 1 },
  { day: 4, meetings: 3 },
  { day: 14, meetings: 3 },
  { day: 21, meetings: 3 },
  { day: 28, meetings: 1 },
];
r - это итератор по всему периоду 
iteratorR - это итератор по массиву митингов

0 1 2 3 4 5 6 7 8 

         0  1  2  3  4
days     1  2  5  6  7
meetings 5  3  2  3  1
*/
public Result method(List<List<Integer>> meetings, int periodLength, int vacationLength){
    Result result = new Result;
    int r = 0, l = 0, iteratorL = 0, iteratorR = 0, sumSlidingWindow = 0;
    for(; r < vacationLength; r++){
        if(iteratorR < meetings.size() && meetings.get(iteratorR).get(0) == r){
            sumSlidingWindow += meetings.get(iteratorR).get(1);
            iteratorR++;
        }
    }
    result.A = l;
    result.B = sumSlidingWindow;
    while(r < periodLength){
        r++;
        l++;
        if(iteratorR < meetings.size() && meetings.get(iteratorR).get(0) == r){
            sumSlidingWindow += meetings.get(iteratorR).get(1);
            iteratorR++;
        }
        if(iteratorL < meetings.size() && meetings.get(iteratorL).get(0) == l){
            sumSlidingWindow -= meetings.get(iteratorL).get(1);
            iteratorL++;
        }
        if(result.B > sumSlidingWindow){
            result.A = l;
            result.B = sumSlidingWindow;
        }
    }
    return result;
}

class Result{
    int A;
    int B;
}

















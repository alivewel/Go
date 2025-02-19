1) Для подсчета с правого края формула неверная была
Было так:
curRes := r - l
Нужно было:
curRes := r - l + 1
``` go
func maxDistToClosest(seats []int) int {
	res := 0
	l, r := 0, 0
    for r < len(seats) {
		for r < len(seats) - 1 && seats[r] == seats[r+1] {
			r++
		}
		if l == 0 {
			res = r - l
		} else if r == len(seats) - 1 {
			// curRes := 0
			// diff := r - l + 1
			// if diff % 2 != 0 {
			// 	curRes = (diff + 1) / 2
			// } else {
			// 	curRes = diff / 2 // при diff == 2 не сходится 
			// }
			curRes := r - l + 1 / 2
			// 3 - 0 + 1 / 2 == 2
			// 2 - 0 + 1 / 2 == 1
			// 1 - 0 + 1 / 2 == 1
			res = max(curRes, res)
		} else {
			curRes := r - l  // неправильно определил при первом запуске
			res = max(curRes, res)
		}
		l = r + 1
		r = r + 1
	}
	return res
}
```

2) Для l == 0 тоже неверно считал
было:
res = r - l
стало:
if l == 0 {
    res = r - l + 1
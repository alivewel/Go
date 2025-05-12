// в плавающем окне мы определяем наиболее встречающийся символ
// двигая правую границу плав. окна мы смотрим сколько раз встретился
// символ в окне на котором стоит правый указатель
// если этот символ встретился больше чем прошлый максимальный символ обновляем переменную maxCount
// (right - left + 1) - maxCount > k  — кол-во символов, которые НЕ равны самому частому символу их и нужно заменить, чтобы сделать все символы одинаковыми.

func characterReplacement(s string, k int) int {
    freqChars := make(map[byte]int, 0)
    maxCount, res, l := 0, 0, 0
    for r := 0; r < len(s); r++ {
        freqChars[s[r]]++
        if freqChars[s[r]] > maxCount {
            maxCount = freqChars[s[r]]
        }
        if r - l + 1 - maxCount > k {
            freqChars[s[l]]--
            l++
        }
        if r - l + 1 > res {
            res = r - l + 1 
        }
    } 
    return res
}

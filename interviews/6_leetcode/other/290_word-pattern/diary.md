1) Тестовый случай сломал мое решение

pattern = "abba" s = "dog dog dog dog"
Output true
Expected false

```
func wordPattern(pattern string, s string) bool {
    strings := strings.Split(s, " ")
    if len(pattern) != len(strings) {
        return false
    }
    matching := make(map[byte]string)
    for i, str := range strings {
        if val, found := matching[pattern[i]]; found && val != str {
            return false
        }
        matching[pattern[i]] = str
    }
    return true
}
```

2) Добавление второй мапы помогло.

```
func wordPattern(pattern string, s string) bool {
    strings := strings.Split(s, " ")
    if len(pattern) != len(strings) {
        return false
    }
    matching := make(map[byte]string)
    matchingReverse := make(map[string]byte)
    for i, str := range strings {
        if val, found := matching[pattern[i]]; found && val != str {
            return false
        }
        if val, found := matchingReverse[str]; found && val != pattern[i] {
            return false
        }
        matching[pattern[i]] = str
        matchingReverse[str] = pattern[i]
    }
    return true
}
```

Время решения 10:47
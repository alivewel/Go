package main

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

func main() {
	pattern := "abba"
	s := "dog cat cat dog"
}

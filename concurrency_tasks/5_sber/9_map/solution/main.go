package main

// что выведет программа?

func main() {
	var m map[string]struct{}
	for k := range m {
		println(k)
	}
	println("end!")
}

// программа выведет end!
// мапа пустая, мы ничего туда не положили

// что такое мапа в го?

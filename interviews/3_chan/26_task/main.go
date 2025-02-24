package main

func joinChannels(in ...<-chan int) {
	out := make(chan int, len(in))
	
	for _, ch := range in {
		go func (ch <-chan int) {

		}(ch)
	}
	
	// for i := 0; i < len(in); 
}

func main() {
	var a, b, c = make(chan int), make(chan int), make(chan int)

	// ...

	for num := range joinChannels(a, b, c) {
		println(num)
	}

}

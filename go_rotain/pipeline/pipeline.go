package main

func Generator(num ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, data := range num {
			out <- data
		}
		close(out)
	}()
	return out
}

func square(val <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for data := range val {
			out <- data * data
		}
		close(out)
	}()
	return out
}

func main() {

	ch := Generator(1, 2, 2, 5, 12, 90, 23, 19)
	out := square(ch)

	for data := range out {
		println(data)
	}
}

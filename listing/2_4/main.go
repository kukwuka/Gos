package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i += 2 {
			ch <- i
		}
		// close(ch)
	}()

	for n := range ch {
		println(n)
	}
}

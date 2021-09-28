package main

func main() {
	ch := make(chan int, 1)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	//дудлок, т.к. мы никогда ниче не получим
	for n := range ch {
		println(n)
	}
}

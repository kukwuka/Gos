package main

import (
	"fmt"
	"time"
)

var or func(channels ...<-chan interface{}) <-chan interface{}

func main() {
	// Пример использования функции:
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	or = func(channels ...<-chan interface{}) <-chan interface{} {
		for {
			for i := range channels {
				select {
				case <-channels[i]:
					return channels[i]
				default:
					continue
				}
			}
		}
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("fone after %v\n", time.Since(start))

}

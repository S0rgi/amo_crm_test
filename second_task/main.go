package main

import (
	"fmt"
)

func Merge(ch1, ch2 <-chan int) <-chan int {
	result := make(chan int)

	go func() {
		defer close(result)
		ch1open, ch2open := true, true

		for ch1open || ch2open {
			select {
			case val, ok := <-ch1:
				if ok {
					result <- val
				} else {
					ch1open = false
				}
			case val, ok := <-ch2:
				if ok {
					result <- val
				} else {
					ch2open = false
				}
			}
		}
	}()

	return result
}

func main() {
	a := make(chan int)
	b := make(chan int)

	go func() {
		defer close(a)
		a <- 4
		a <- 1
	}()

	go func() {
		defer close(b)
		b <- 2
		b <- 4
	}()

	for v := range Merge(a, b) {
		fmt.Println(v)
	}
}

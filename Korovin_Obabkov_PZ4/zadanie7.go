package main

import "fmt"

func main() {
	set, get := make(chan int), make(chan int)
	go func() {
		var state int
		for {
			select {
			case state = <-set:
			case get <- state:
			}
		}
	}()
	set <- 77
	fmt.Println(<-get)
}

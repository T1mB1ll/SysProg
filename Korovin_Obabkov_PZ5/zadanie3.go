package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan string, 1)
	ch <- "Важное данное"
	test := func() {
		defer wg.Done()
		select {
		case data := <-ch:
			fmt.Printf(data)
		default:
			fmt.Println("\nДанного нет важного тем")
		}
	}
	go test()
	go test()
	wg.Wait()
}

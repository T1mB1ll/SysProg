package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	done := make(chan bool)
	go byevirus(&wg, done)
	time.Sleep(7 * time.Second)
	done <- true
	wg.Wait()
}

func byevirus(wg *sync.WaitGroup, done <-chan bool) {
	defer wg.Done()
	for {
		select {
		case <-done:
			fmt.Println("Пока пока вирусняк поганый")
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("Я вирус я вирус")
		}
	}
}

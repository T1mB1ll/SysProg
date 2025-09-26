package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan string, 1)
	for i := 1; i < 8; i++ {
		wg.Add(1)
		ch <- "Горутина"
		go infoutput(&wg, ch, i)
	}
	wg.Wait()
	fmt.Println("Все горутины завершили работу")
}

func infoutput(wg *sync.WaitGroup, ch <-chan string, id int) {
	defer wg.Done()
	if len(ch) == 0 {
		return
	}
	fmt.Println(<-ch, id)
}

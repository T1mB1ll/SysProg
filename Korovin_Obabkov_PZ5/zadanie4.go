package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	visoko := make(chan string, 2)
	obichno := make(chan string, 2)
	obichno <- "Обычный приоритет 1"
	visoko <- "Высокий приоритет 1"
	visoko <- "Высокий приоритет 2"
	obichno <- "Обычный приоритет 2"
	proverik := func() {
		defer wg.Done()
		for len(visoko) > 0 {
			fmt.Println(<-visoko)
		}
		for len(obichno) > 0 {
			fmt.Println(<-obichno)
		}
	}
	go proverik()
	wg.Wait()
}

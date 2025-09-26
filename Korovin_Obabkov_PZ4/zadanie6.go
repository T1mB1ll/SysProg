package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	sources := []string{"Database1", "Database2", "API"}
	for _, src := range sources {
		go func(s string) {
			select {
			case ch <- "Первый успешный результат пришёл из " + s:
			}
		}(src)
	}
	fmt.Println(<-ch)
}

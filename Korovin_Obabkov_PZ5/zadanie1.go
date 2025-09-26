package main

import (
	"fmt"
	"time"
)

func longwait(ch1 chan<- string) {
	time.Sleep(3 * time.Second)
	ch1 <- "Хельо"
	close(ch1)
}
func main() {

	ch1 := make(chan string)
	go longwait(ch1)
	select {
	case res := <-ch1:
		fmt.Printf("Хельо тут - %v", res)
	case <-time.After(2 * time.Second):
		fmt.Print("Хельо не будет")
	}

}

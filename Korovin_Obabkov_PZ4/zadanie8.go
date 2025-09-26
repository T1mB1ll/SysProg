package main

import "fmt"

func main() {
	a := []string{"1", "2", "3"}
	zaprosi := make(chan string)
	go func() {
		for i := 0; ; i++ {
			zaprosi <- a[i%len(a)]
		}
	}()
	for i := 1; i < 16; i++ {
		fmt.Println("Запрос обработал", <-zaprosi)
	}
}

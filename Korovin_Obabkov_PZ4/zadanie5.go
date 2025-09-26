package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"sync"
)

func main() {
	files := []string{"Korovin.txt", "Obabkov.txt"}
	var wg sync.WaitGroup
	for _, f := range files {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			data, _ := os.ReadFile(file)
			fmt.Println(file, md5.Sum(data))
		}(f)
	}
	wg.Wait()
}

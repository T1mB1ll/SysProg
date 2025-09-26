package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	urls := []string{"https://www.roblox.com/login", "https://market.yandex.ru/kolesoprizov?_redirectCount=1", "https://github.com"}
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			resp, _ := http.Get(u)
			fmt.Println(u, resp.StatusCode)
		}(url)
	}
	wg.Wait()
}

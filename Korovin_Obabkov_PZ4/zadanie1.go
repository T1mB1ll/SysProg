package main
import (
	"fmt"
	"sync"
	"time"
)
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	work := func(){
		for i:=1;i<6;i++{
		fmt.Println(i)
		time.Sleep(1)
		}
		wg.Done()
	}
	go work()
	wg.Wait()
}
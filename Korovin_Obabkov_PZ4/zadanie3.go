package main
import (
	"fmt"
	"sync"
	"time"
)
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	tick := time.Tick(200 * time.Millisecond)
	work :=func(){
		for i := 1; i<16; i++ {
			<-tick
			fmt.Println(i)
		}
		wg.Done()
	}
	go work()
	wg.Wait()
}

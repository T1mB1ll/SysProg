package main
import (
	"fmt"
	"sync"
)
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	jobs := make(chan int)
	results := make(chan int)
	job:=func(){
		for i:=1;i<11;i++{
			jobs <- i
		}
		close(jobs)
	}
	result:=func(){
		for true{
			val, ok := <- jobs
			if (!ok){
				break
			}
			results <- val*val
		}
		close(results)
	}
	getResult:= func(){
	for true{
		val, ok := <- results
			if (!ok){
				break
				}
			fmt.Println(val)
			}
		wg.Done()
	}
	go job()
	go result()
	go getResult()
	wg.Wait()
}
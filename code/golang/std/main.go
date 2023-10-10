package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	for i := 1; i < 10; i++ {
		wg.Add(1)
		go read(ch, fmt.Sprintf("Iam:%d", i), wg)
	}

	for i := 1; i < 10; i++ {

		ch <- i
	}
	wg.Wait()

}

func read(ch chan int, who string, wg *sync.WaitGroup) {

	abc := <-ch
	fmt.Println(who, "read", abc)
	wg.Done()
}

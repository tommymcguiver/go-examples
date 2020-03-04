package main

import (
	"fmt"
	"sync"
)

func main() {
	values := []string{"a", "b", "c"}
	var wg sync.WaitGroup
	for _, v := range values {
		wg.Add(1)
		go func() {
			fmt.Println(v)
			wg.Done()
		}()
	}
	wg.Wait()
}

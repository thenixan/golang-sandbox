package main

import (
	"fmt"
	"sync"
	"time"
	"userdatabase"
)

func main() {
	start := time.Now()
	var wg = sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		go func() {
			wg.Add(1)
			defer wg.Done()
			err := userdatabase.CheckAuth("dial","ESXbgvqY")
			if err != nil {
				fmt.Println(err)
			}
		}()
	}
	wg.Wait()
	fmt.Printf("Took %s\n", time.Since(start))
}

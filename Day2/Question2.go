package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	sum := 0.0
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(i%50) * time.Millisecond)
			mu.Lock()
			sum += float64(rand.Intn(100))
			mu.Unlock()

		}()
	}
	wg.Wait()
	fmt.Println("answer", sum/200.0)
}

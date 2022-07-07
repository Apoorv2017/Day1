package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func deposit(x int, amount *int, wg *sync.WaitGroup, mu *sync.Mutex) {

	defer wg.Done()

	mu.Lock()
	fmt.Println("current Balance: ", *amount, " amount to be deposited: ", x)
	(*amount) += x
	fmt.Println("Balance debited to account, current Balance: ", *amount)
	mu.Unlock()
	fmt.Println("-----")
}

func withdrawal(x int, amount *int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	mu.Lock()
	fmt.Println("current Balance: ", *amount, " amount to be withddrawn: ", x)
	if (*amount) < x {
		fmt.Println("Balance not sufficient")
	} else {
		(*amount) -= x
		fmt.Println("Balance credited form account, current Balance: ", *amount)
	}
	mu.Unlock()
	fmt.Println("-----")

}

func main() {
	amount := 500
	temp := &amount
	var wg sync.WaitGroup
	var mu sync.Mutex
	for _, i := range rand.Perm(5) {
		x := rand.Intn(100000)
		if i%2 == 0 {
			wg.Add(1)
			go deposit(x, temp, &wg, &mu)
		} else {
			wg.Add(1)
			go withdrawal(x, temp, &wg, &mu)
		}
	}
	wg.Wait()
}

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const N = 1000000
	add := 0           // N
	sub := 0           // -N
	mu := sync.Mutex{} // Mutual Exclusion
	for i := 0; i < N; i++ {
		add++
		go func() {
			defer func() {
				mu.Lock()   // Lock
				sub--       // critical section
				mu.Unlock() // unLock
			}()
		}()
	}
	for {
		h, m, s := time.Now().Clock()
		fmt.Printf("%d-%d-%d : add %d, sub %d\r", h, m, s, add, sub) // \r reset cursor
		if add == N && sub == -N {
			fmt.Println("Success")
			break
		}
	}

	fmt.Println("Main Done", add, sub)
}

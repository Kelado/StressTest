package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

var parallelWorkers = 5

func isPrime(num int) bool {
	var count int
	count = 0

	for i := 2; i < num/2; i++ {
		if num%i == 0 {
			count++
			break
		}
	}

	if count == 0 && num != 1 {
		// fmt.Printf("%d ", num)
		return true
	} else {
		return false
	}
}

func findPrimeNumbersUpTo(n int) int {
	primes := 0
	for i := 0; i < n; i++ {
		if isPrime(i) {
			primes++
		}
	}
	return primes
}

func main() {

	limit := flag.Int("l", 1000, "Limit of search")
	workers := flag.Int("w", 1, "Number of workers")
	flag.Parse()

	fmt.Println("Start Stress test with")
	fmt.Println("Limit: ", *limit)
	fmt.Println("Workers: ", *workers)
	fmt.Println()

	var wg sync.WaitGroup

	wg.Add(*workers)

	start := time.Now()

	for i := 0; i < *workers; i++ {
		go func() {
			defer wg.Done()
			findPrimeNumbersUpTo(*limit)
		}()
	}

	wg.Wait()
	fmt.Println("Everything finished in ", time.Since(start))
}

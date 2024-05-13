package main

import (
	"encoding/hex"
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/crypto"
)

const (
	desiredOutput = "371636862"
	fixedA        = "371636862"
	fixedB        = "17355924313246092532519660475566608967507701394993171841024585092939205358430"
)

func main() {
	var wg sync.WaitGroup
	resultChan := make(chan []byte, 1)

	// Start multiple goroutines to perform the brute-force attack
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go bruteForce(uint64(i), resultChan, &wg)
	}

	// Wait for a result from any of the goroutines
	result := <-resultChan
	if result != nil {
		fmt.Printf("Found C: %x\n", result)
	} else {
		fmt.Println("No solution found")
	}

	// Wait for all goroutines to finish
	wg.Wait()
}

func bruteForce(start uint64, resultChan chan<- []byte, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := start; ; i += 100 {
		c := make([]byte, 32)
		for j := 0; j < 32; j++ {
			c[j] = byte(i >> uint(8*j) & 0xff)
		}

		hash := crypto.Keccak256(append(append([]byte(fixedA), []byte(fixedB)...), c...))
		output := hex.EncodeToString(hash)

		if output == desiredOutput {
			resultChan <- c
			return
		}
	}
}

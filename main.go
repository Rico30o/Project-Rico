package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

// checkOnline checks if a server is online by attempting to establish a TCP connection.
func checkOnline(server string, wg *sync.WaitGroup) bool {
	defer wg.Done()

	conn, err := net.DialTimeout("tcp", server, 2*time.Second)
	if err != nil {
		fmt.Printf("%s is offline\n", server)
		return false
	}
	defer conn.Close()

	fmt.Printf("%s is online\n", server)
	return true
}

func main() {
	// List of servers to check
	servers := []string{"example.com:80", "google.com:80", "nonexistent.com:80"}

	// WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Channel to receive results from goroutines
	resultCh := make(chan bool)

	for _, server := range servers {
		// Increment the WaitGroup counter for each goroutine
		wg.Add(1)

		// Launch a goroutine to check the online status of each server
		go func(server string) {
			resultCh <- checkOnline(server, &wg)
		}(server)
	}

	// Close the result channel when all goroutines are done
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// Collect and process the results
	for result := range resultCh {
		// Process the result (true for online, false for offline)
		if result {
			// Do something when online
		} else {
			// Do something when offline
		}
	}
}

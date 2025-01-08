package main

import (
	"flag"
	"log"
	"sync"

	memory "github.com/ssbostan/reserved-capacity-manager/consumer/internal"
)

func main() {
	var wg sync.WaitGroup

	numberOfWorkers := flag.Int("workers", 2, "Number of goroutines to spawn up.")
	memoryReserveMegabytes := flag.Int("memory", 512, "Amount of Memory (MB) to reserve.")
	flag.Parse()

	memoryPerWorker := *memoryReserveMegabytes / *numberOfWorkers

	log.Printf("Number of workers: %d\n", *numberOfWorkers)
	log.Printf("Total memory reservation (MB): %d\n", *memoryReserveMegabytes)
	log.Printf("Memory reservation per worker (MB): %d\n", memoryPerWorker)

	wg.Add(*numberOfWorkers)
	for i := 0; i < *numberOfWorkers; i++ {
		go memory.ReserveMemory(&wg, memoryPerWorker)
	}

	wg.Wait()
}

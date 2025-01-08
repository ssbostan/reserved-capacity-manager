package memory

import (
	"math/rand"
	"sync"
	"time"
)

func ReserveMemory(wg *sync.WaitGroup, reserveMegabytes int) {
	defer wg.Done()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	reserveBytes := reserveMegabytes * 1024 * 1024

	memory := make([]byte, reserveBytes)

	for i := range memory {
		memory[i] = byte(r.Intn(256))
	}

	for {
		memory[0] = byte(r.Intn(256))
		time.Sleep(10 * time.Second)
	}
}

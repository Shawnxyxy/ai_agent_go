package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func RunTasks(tasks []string) {
	var wg sync.WaitGroup

	for _, t := range tasks {
		wg.Add(1)
		go func(task string) {
			defer wg.Done()
			fmt.Println("Processing:", task)
			time.Sleep(1 * time.Second)
		}(t)
	}
	wg.Wait()
	fmt.Println("All tasks done")
}
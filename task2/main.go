package main //exr2

import (
	"fmt"
	"sync"
)

type LogBuffer struct {
	buffer []string
	mutex  sync.Mutex
}

func (lb *LogBuffer) WriteLog(message string) {
	lb.mutex.Lock()
	defer lb.mutex.Unlock()

	lb.buffer = append(lb.buffer, message)
}

func main() {
	var wg sync.WaitGroup
	logBuffer := &LogBuffer{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			logBuffer.WriteLog(fmt.Sprintf("Log message %d", i))
		}(i)
	}

	wg.Wait()

	fmt.Println("Log Buffer Contents:")
	for _, logMessage := range logBuffer.buffer {
		fmt.Println(logMessage)
	}
}

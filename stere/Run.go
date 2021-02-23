package stere

import (
	
	"os"
	"sync"
)

func Run(file *os.File) {
	wg := sync.WaitGroup{}
	lines := make(chan string)

	wg.Add(2)

	go ReadFile(file, lines)
    go Load(lines)

	wg.Wait()
}
package stere

import (

	"os"
	"sync"
)

func Run(file *os.File) {
	wg := &sync.WaitGroup{}
	lines := make(chan string)

	wg.Add(2)

	go ReadFile(file, lines, wg)
    go Load(lines, wg)

	wg.Wait()
}
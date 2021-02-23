package stere

import (
	"fmt"
	"sync"
	"bufio"
)

func ReadFile(file *os.File, wg *sync.WaitGroup, chLine chan string) {
	defer wg.Done()
	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		line := scanner.Text()

		chLine <- line
	}

	if err := scanner.Err(); err != nil {
		chLine <- fmt.Sprint(err)
	}

	close(chLine)
}
package stere

import (

	"os"
	"fmt"
	"sync"
	"bufio"
)

func ReadFile(file *os.File, chLine chan string, wg *sync.WaitGroup) {
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
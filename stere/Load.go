package stere

import (
	"fmt"
	"sync"
)

func Load(lines chan string, wg *sync.WaitGroup){
	defer wg.Done()

	for line := range lines {
    	fmt.Println(line)
    }
}
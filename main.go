package main

import (

	"bufio"
	"fmt"
	"sync"
	"os"
	"github.com/spf13/viper"
)

var wg sync.WaitGroup

func configLoad() (*viper.Viper, error) {
	conf := viper.GetViper()
	conf.AddConfigPath(".")
	conf.SetConfigFile("conf")
	conf.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return conf, nil
}

func read(file *os.File, chLine chan string) {
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

func load(lines chan string){
	defer wg.Done()

	for line := range lines {
    	fmt.Println(line)
    }
}

func main() {

	conf, err := configLoad()
	if err != nil {
		panic(err)
	}

	file, err := os.Open(conf.GetString("data"))
	if err != nil {
		panic(err)
	}
	
	lines := make(chan string)

	wg.Add(2)

	go read(file, lines)
    go load(lines)

	wg.Wait()
}
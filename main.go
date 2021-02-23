package main

import (

	"os"
	"github.com/gothello/short/config"
)

func main() {

	conf, err := config.Load()
	if err != nil {
		panic(err)
	}

	file, err := os.Open(conf.GetString("data"))
	if err != nil {
		panic(err)
	}

	proxys, err := os.Open(conf.GetString("proxys"))
	if err != nil {
		panic(err)
	}
	
	Run(file)
	Run(proxys)
}

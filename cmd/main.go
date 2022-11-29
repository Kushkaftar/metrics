package main

import (
	"flag"
	"log"
	"metrics/internal/run"
	"metrics/pkg/config"
)

const pathToConfig = "./configs"

func main() {

	var fileName string

	flag.StringVar(&fileName, "env", "", "desc")
	flag.Parse()

	c, err := config.NewConfig(fileName, pathToConfig)
	if err != nil {
		log.Fatalf("Error init config - %s", err)
	}

	if err = run.Run(c); err != nil {
		log.Fatalf("fatal mistake - %s", err)
	}
}

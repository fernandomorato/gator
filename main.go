package main

import (
	"fmt"
	"log"

	"github.com/fernandomorato/gator/internal/config"
)

func main() {
	config, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	err = config.SetUser("morato")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(config)
}

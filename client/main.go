package main

import (
	"log"

	"github.com/alekstet/linux_service_checker/client/cmd"
)

func main() {
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

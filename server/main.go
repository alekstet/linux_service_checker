package main

import (
	"log"

	"github.com/alekstet/linux_service_checker/server/cmd"
)

func main() {
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

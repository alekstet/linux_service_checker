package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alekstet/linux_service_checker/cmd"
)

func main() {
	f, err := os.Open("C:/Users/atete/.ssh/VDS.ppk")
	fmt.Println("file:", f, err)
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

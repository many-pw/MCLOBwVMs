package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("mclob v0.0.1")
		fmt.Println("")
		fmt.Println("  --add-service")

		return
	}

	if os.Args[1] == "--add-service" {

		service := os.Args[2]
		mysqlPassword := os.Args[3]

		addService(service, mysqlPassword)
	}

}

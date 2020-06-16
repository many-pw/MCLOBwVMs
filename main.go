package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("mclob v0.0.1")
		fmt.Println("")
		fmt.Println("  --list")

		return
	}

	if os.Args[1] == "--list" {

		fmt.Println("")
		fmt.Println("1. DigitalOcean")
		fmt.Println("2. Vultr")
		fmt.Println("")
	}

}

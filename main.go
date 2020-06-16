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

		// https://www.terraform.io/docs/providers/do/r/droplet.html
		// https://www.terraform.io/docs/providers/do/r/loadbalancer.html

		// https://www.terraform.io/docs/providers/vultr/r/server.html
		// https://www.terraform.io/docs/providers/vultr/r/load_balancer.html

		// https://www.terraform.io/docs/providers/aws/r/instance.html
		// https://www.terraform.io/docs/providers/aws/r/lb.html

		// https://www.terraform.io/docs/providers/google/r/compute_instance.html
		// https://www.terraform.io/docs/providers/google/r/compute_target_pool.html

		// https://www.terraform.io/docs/providers/linode/r/instance.html
		// https://www.terraform.io/docs/providers/linode/r/nodebalancer.html

		fmt.Println("")
		fmt.Println("1. DigitalOcean")
		fmt.Println("2. Vultr")
		fmt.Println("")
	}

}

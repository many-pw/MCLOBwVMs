package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
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
		doId := os.Args[4]
		doSecret := os.Args[5]

		addService(service, mysqlPassword, doId, doSecret)
	} else if os.Args[1] == "--terraform" {
		exec.Command("openssl", "genrsa", "-out", "priv_dkim.key", "1024").Run()
		exec.Command("openssl", "rsa", "-in", "priv_dkim.key",
			"-pubout", "-out", "pub_dkim.key").Run()
		b, _ := ioutil.ReadFile("pub_dkim.key")
		lines := []string{}
		for _, line := range strings.Split(string(b), "\n") {
			if len(line) == 0 || strings.HasPrefix(line, "---") {
				continue
			}
			lines = append(lines, line)
		}
		content := fmt.Sprintf("\"v=DKIM1; k=rsa; p=%s\"", []byte(strings.Join(lines, "")))
		ioutil.WriteFile("pub_dkim.key", []byte(content), 0644)
	} else if os.Args[1] == "--dkim" {

	}

}

package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func addService(name, mysqlPassword string) {
	service := fmt.Sprintf(SERVICE, name, name, name, mysqlPassword)
	ioutil.WriteFile("/etc/systemd/system/"+name+".service", []byte(service), 0644)
	exec.Command("systemctl", "daemon-reload").Run()
	exec.Command("systemctl", "enable", name).Run()
	exec.Command("systemctl", "start", name).Run()
}

const SERVICE = `#/usr/lib/systemd/system/%s.service
[Unit]
Description=%s
After=network-online.target

[Service]
User=root
WorkingDirectory=/
ExecStart=/%s 3000 3000
Type=simple
Restart=always
RestartSec=0
LimitNOFILE=65536
Environment=GIN_MODE=release
Environment=DB_USER=root
Environment=DB_NAME=jjaa_me
Environment=DB_HOST=127.0.0.1
Environment=DB_PASSWORD=%s
Environment=DB_PORT=3306

[Install]
WantedBy=multi-user.target`

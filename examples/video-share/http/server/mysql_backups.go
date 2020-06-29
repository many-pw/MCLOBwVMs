package server

import (
	"bytes"
	"os/exec"
	"time"

	"jjaa.me/http/controllers"
)

func DoMysqlBackups() {
	for {
		bu, _ := exec.Command("mysqldump", "jjaa_me").Output()

		controllers.UploadToBucket("mysql_backup.sql", bytes.NewReader(bu))
		time.Sleep(time.Second * 86400)
	}
}

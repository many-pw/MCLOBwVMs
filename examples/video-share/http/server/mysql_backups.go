package server

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func DoMysqlBackups() {
	for {
		bu, _ := exec.Command("mysqldump", "jjaa_me").Output()

		endpoint := "sfo2.digitaloceanspaces.com"
		region := "sfo2"
		sess := session.Must(session.NewSession(&aws.Config{
			Endpoint: &endpoint,
			Credentials: credentials.NewStaticCredentials(os.Getenv("DO_ID"),
				os.Getenv("DO_SECRET"), ""),
			Region: &region,
		}))

		uploader := s3manager.NewUploader(sess)

		myBucket := "jjaa.me.cloud"
		myString := "mysql_backup.sql"
		_, err := uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(myBucket),
			Key:    aws.String(myString),
			Body:   bytes.NewReader(bu),
		})
		if err != nil {
			fmt.Println(err)
		}

		time.Sleep(time.Second * 86400)
	}
}

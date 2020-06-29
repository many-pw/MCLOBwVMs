package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func DownloadFromBucket(filename string) {
	endpoint := "sfo2.digitaloceanspaces.com"
	region := "sfo2"
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: &endpoint,
		Credentials: credentials.NewStaticCredentials(os.Getenv("DO_ID"),
			os.Getenv("DO_SECRET"), ""),
		Region: &region,
	}))

	downloader := s3manager.NewDownloader(sess)

	myBucket := "jjaa.me.cloud"
	file, _ := os.Create(filename)
	defer file.Close()

	numBytes, _ := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(myBucket),
			Key:    aws.String(filename),
		})

	fmt.Println(numBytes)
}

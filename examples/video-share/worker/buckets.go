package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var bucket = "jjaa.me.cloud"

func Session() client.ConfigProvider {
	endpoint := "sfo2.digitaloceanspaces.com"
	region := "sfo2"
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: &endpoint,
		Credentials: credentials.NewStaticCredentials(os.Getenv("DO_ID"),
			os.Getenv("DO_SECRET"), ""),
		Region: &region,
	}))
	return sess
}
func DownloadFromBucket(filename string) {

	downloader := s3manager.NewDownloader(Session())

	file, _ := os.Create("orig_" + filename)
	defer file.Close()

	numBytes, _ := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(filename),
		})

	fmt.Println(numBytes)
}
func UploadToPublicBucket(filename string) {

	uploader := s3manager.NewUploader(Session())

	f, _ := os.Open(filename)

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String("public/" + filename),
		Body:   f,
	})
	f.Close()
	os.Remove(filename)
	if err != nil {
		fmt.Println(err)
	}
}

func DeleteFileFromBucket(filename string) {
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
	}

	svc := s3.New(Session())
	_, err := svc.DeleteObject(input)
	if err != nil {
		fmt.Println(err)
	}
}

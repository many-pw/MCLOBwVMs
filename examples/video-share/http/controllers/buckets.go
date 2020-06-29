package controllers

import (
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadToBucket(filename string, r io.Reader) {
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
	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(myBucket),
		Key:    aws.String(filename),
		Body:   r,
	})
	if err != nil {
		fmt.Println(err)
	}
}

package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3Client *s3.S3
	s3Bucket string
)

func init() {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String("us-east-1"),
			Credentials: credentials.NewStaticCredentials(
				"ASDEXAMPLE",
				"was/Ssdadaw/ASDEXAMPLEKEY",
				"",
			),
		},
	)
	if err != nil {
		panic(err)
	}
	s3Client = s3.New(sess)
	s3Bucket = "go-lexlab-example-bucket"
}

func main() {
	dir, err := os.Open(".tmp")
	if err != nil {
		panic(err)
	}
	defer dir.Close()
}

func uploadFile(fileName string) {
	completeFileName := fmt.Sprintf("./tmp/%s", fileName)
	f, err := os.Open(completeFileName)
	if err != nil {
		fmt.Printf("Error opening file %s\n", completeFileName)
		return
	}
	defer f.Close()

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(fileName),
		Body:   f,
	})
	if err != nil {
		fmt.Printf("Error uploading file %s\n", completeFileName)
	}
	fmt.Printf("File %s uploaded successfully\n", completeFileName)

}

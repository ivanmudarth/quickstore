package handlers

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type s3config struct {
	endpoint   string
	client     *s3.S3
	uploader   *s3manager.Uploader
	downloader *s3manager.Downloader
	bucketName *string
}

// Export config to be used in handlers
var AWSConfig s3config

func CreateAWSConfig() {
	// Initialize AWS Session
	localstackEndpoint := "http://localhost:4566"
	creds := credentials.NewStaticCredentials("test", "test", "")
	config := &aws.Config{
		Region:           aws.String("us-east-1"),
		Endpoint:         aws.String(localstackEndpoint),
		Credentials:      creds,
		S3ForcePathStyle: aws.Bool(true), // Necessary for LocalStack S3
	}
	sess := session.Must(session.NewSession(config))
	s3Uploader := s3manager.NewUploader(sess)
	s3Downloader := s3manager.NewDownloader(sess)
	s3Client := s3.New(sess)

	// Create new bucket if it doesn't exist
	bucketName := "my-bucket"
	_, err := s3Client.CreateBucket(&s3.CreateBucketInput{
		Bucket: &bucketName,
	})
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	} else {
		log.Printf("Bucket '%s' created or already exists.", bucketName)
	}

	// Set values for struct representing AWS config
	AWSConfig = s3config{endpoint: localstackEndpoint, client: s3Client, uploader: s3Uploader, downloader: s3Downloader, bucketName: aws.String(bucketName)}
}

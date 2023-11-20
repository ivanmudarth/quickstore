package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	// endpoint := "https://0.0.0.0:4566"

	// The session the S3 Uploader will use
	localstackEndpoint := "http://localhost:4566"
	creds := credentials.NewStaticCredentials("test", "test", "")
	config := &aws.Config{
		Region:           aws.String("us-east-1"),
		Endpoint:         aws.String(localstackEndpoint),
		Credentials:      creds,
		S3ForcePathStyle: aws.Bool(true), // Necessary for LocalStack S3
	}

	// Create session
	sess := session.Must(session.NewSession(config))
	s3Client := s3.New(sess)

	// f, err := os.Open("cat.jpeg")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Create S3 Bucket
	bucketName := "my-bucket"
	_, err := s3Client.CreateBucket(&s3.CreateBucketInput{
		Bucket: &bucketName,
	})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Bucket '%s' created or already exists.", bucketName)
	}

	// Upload object to S3.
	// objectKey := "test-object2"
	// content := "Hello, again LocalStack!"

	// result, err := s3Client.PutObject(&s3.PutObjectInput{
	// 	Bucket: aws.String(bucketName),
	// 	Key:    aws.String(objectKey),
	// 	Body:   strings.NewReader(content),
	// })

	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	log.Println(result)
	// }

	// List objects
	result2, err := s3Client.ListObjects(&s3.ListObjectsInput{Bucket: aws.String(bucketName)})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(result2)
	}
}

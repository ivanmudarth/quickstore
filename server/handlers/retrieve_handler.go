package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func DisplayHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all file URLs...")

	// Download all files from S3
	result, err := getURLsFromS3()
	if err != nil {
		http.Error(w, "Error downloading from S3", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func getAllKeys() []string {
	// TODO: finish once upload metadata works
	return []string{"cat.jpeg", "dog.jpeg"}
}

func getURLsFromS3() ([]string, error) {
	keys := getAllKeys()
	file_urls := []string{}

	// Check that all keys exist in S3 bucket
	for _, k := range keys {
		// Input parameters for HeadObject operation
		input := &s3.HeadObjectInput{
			Bucket: AWSConfig.bucketName,
			Key:    aws.String(k),
		}

		// Check if the object (key) exists
		_, err := AWSConfig.client.HeadObject(input)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		url := AWSConfig.endpoint + "/" + *AWSConfig.bucketName + "/" + k
		file_urls = append(file_urls, url)
	}

	fmt.Println("Got all file URLs\n")
	return file_urls, nil
}

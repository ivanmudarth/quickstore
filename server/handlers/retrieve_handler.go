package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"../database"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func DisplayHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all file URLs...")

	// Download all files from S3
	result, err := getURLsFromS3()
	if err != nil {
		http.Error(w, "Error getting files from S3", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func getAllKeys() ([]string, error) {
	// get all S3Keys from File table
	rows, err := database.DB.Query(`
		SELECT S3Key FROM File 
		`)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var result []string
	for rows.Next() {
		var key string
		if err := rows.Scan(&key); err != nil {
			log.Fatal(err)
			return nil, err
		}
		result = append(result, key)
	}

	return result, nil
}

func getURLsFromS3() ([]string, error) {
	keys, err := getAllKeys()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
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

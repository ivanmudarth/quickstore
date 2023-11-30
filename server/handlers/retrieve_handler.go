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

type FileInfo struct {
	Key      string
	URL      string
	Name     string
	Size     string
	UserTags []string
}

func DisplayHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all file URLs...")

	// Download all files from S3
	result, err := getURLsFromS3()
	if err != nil {
		http.Error(w, "Error getting files from S3", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Error getting files from S3", http.StatusBadRequest)
		return
	}
	w.Write(data)
}

func getUserTagsForFile(fileID int) ([]string, error) {
	// get all tags for current file from File table
	rows, err := database.DB.Query(`
		SELECT Name FROM Tag
		WHERE FileID = ?
		`, fileID)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var userTags []string
	for rows.Next() {
		var tag string
		if err := rows.Scan(&tag); err != nil {
			log.Fatal(err)
			return nil, err
		}
		userTags = append(userTags, tag)
	}

	return userTags, nil
}

func getAllFileInfo() ([]FileInfo, error) {
	// get all file info from File table
	rows, err := database.DB.Query(`
		SELECT FileID, S3Key, Name, Size FROM File 
		`)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// process results from queries
	var result []FileInfo
	for rows.Next() {
		var fileID int
		var key, fileName, fileSize string
		if err := rows.Scan(&fileID, &key, &fileName, &fileSize); err != nil {
			log.Fatal(err)
			return nil, err
		}

		userTags, err := getUserTagsForFile(fileID)
		if err != nil {
			return nil, err
		}

		result = append(result, FileInfo{key, "", fileName, fileSize, userTags})
	}
	return result, nil
}

func getURLsFromS3() ([]FileInfo, error) {
	info, err := getAllFileInfo()
	if err != nil {
		return nil, err
	}

	// Check that all keys exist in S3 bucket
	for idx, i := range info {
		key := i.Key
		// Input parameters for HeadObject operation
		input := &s3.HeadObjectInput{
			Bucket: AWSConfig.bucketName,
			Key:    aws.String(key),
		}

		// Check if the object (key) exists
		_, err := AWSConfig.client.HeadObject(input)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		// Modify file info to include AWS url of file
		url := AWSConfig.endpoint + "/" + *AWSConfig.bucketName + "/" + key
		info[idx].URL = url
	}

	fmt.Println("Got all file URLs\n")
	return info, nil
}

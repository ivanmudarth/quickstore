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
	AutoTags []string
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

func getTagsByType(fileID int, tagType string) ([]string, error) {
	// get all tags for current file from File table
	rows, err := database.DB.Query(`
		SELECT Name FROM Tag
		WHERE FileID = ? AND Type = ?
		`, fileID, tagType)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var tags []string
	for rows.Next() {
		var tag string
		if err := rows.Scan(&tag); err != nil {
			log.Fatal(err)
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

func getAllFileInfo() ([]FileInfo, error) {
	// get necessary file info from File table
	rows, err := database.DB.Query(`
		SELECT FileID, S3Key, Name, Size FROM File 
		`)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// get file's associated user and auto tags from Tag table
	var result []FileInfo
	for rows.Next() {
		var fileID int
		var key, fileName, fileSize string
		if err := rows.Scan(&fileID, &key, &fileName, &fileSize); err != nil {
			log.Fatal(err)
			return nil, err
		}

		userTags, err := getTagsByType(fileID, "User")
		if err != nil {
			return nil, err
		}

		autoTags, err := getTagsByType(fileID, "Auto")
		if err != nil {
			return nil, err
		}

		result = append(result, FileInfo{key, "", fileName, fileSize, userTags, autoTags})
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
		s3Key := i.Key
		// Input parameters for HeadObject operation
		input := &s3.HeadObjectInput{
			Bucket: AWSConfig.bucketName,
			Key:    aws.String(s3Key),
		}

		// Check if the object (key) exists
		_, err := AWSConfig.client.HeadObject(input)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		// Modify file info to include AWS url of file
		url := AWSConfig.endpoint + "/" + *AWSConfig.bucketName + "/" + s3Key
		info[idx].URL = url
	}

	fmt.Println("Got all file URLs\n")
	return info, nil
}

package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"../database"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jmoiron/sqlx"
)

type FileInfo struct {
	ID       int
	Key      string
	URL      string
	Name     string
	Size     string
	Type     string
	UserTags []string
	AutoTags []string
}

func DisplayHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all file URLs...")

	// Get appropriate file metadata given tag input
	tags := r.URL.Query()["tags[]"]
	var fileInfo []FileInfo
	var err error
	if len(tags) == 0 {
		fileInfo, err = getAllFileInfo()
	} else {
		fileInfo, err = getSearchFileInfo(tags)
	}
	if err != nil {
		return
	}

	// verify files to be displayed are in S3
	result, err := getURLsFromS3(fileInfo)
	if err != nil {
		http.Error(w, "Error getting files from S3", http.StatusBadRequest)
		return
	}

	// send file info to be displayed back to client
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Error getting files from S3", http.StatusBadRequest)
		return
	}
	w.Write(data)
}

func processQueryResults(rows *sql.Rows) ([]FileInfo, error) {
	// add query results to list of FileInfo objects
	fileList := []FileInfo{}
	for rows.Next() {
		var fileID int
		var key, fileName, fileSize, fileType, tagName, tagType string
		if err := rows.Scan(&fileID, &key, &fileName,
			&fileSize, &fileType, &tagName, &tagType); err != nil {
			log.Fatal(err)
			return nil, err
		}

		// if FileInfo entry exists, update its tags. Otherwise create entry
		lastIdx := len(fileList) - 1
		if lastIdx >= 0 && fileList[lastIdx].ID == fileID {
			if tagType == "User" {
				fileList[lastIdx].UserTags = append(fileList[lastIdx].UserTags, tagName)
			} else {
				fileList[lastIdx].AutoTags = append(fileList[lastIdx].AutoTags, tagName)
			}
		} else {
			if tagType == "User" {
				fileList = append(fileList, FileInfo{fileID, key, "", fileName, fileSize, fileType, []string{tagName}, []string{}})
			} else {
				fileList = append(fileList, FileInfo{fileID, key, "", fileName, fileSize, fileType, []string{}, []string{tagName}})
			}
		}
	}

	return fileList, nil
}

func getSearchFileInfo(tags []string) ([]FileInfo, error) {
	// make query string
	q, args, err := sqlx.In(`
		SELECT s.FileID, s.S3Key, s.Name, s.Size, s.Type, t2.Name, t2.Type   
		FROM (SELECT DISTINCT f.FileID, f.S3Key, f.Name, f.Size, f.Type
			FROM File f INNER JOIN Tag t ON f.FileID = t.FileID
			WHERE LOWER(t.Name) IN (?)) s
		INNER JOIN Tag t2 ON s.FileID = t2.FileID
		`, tags)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// make actual query
	rows, err := database.DB.Query(q, args...)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// proess results
	fileList, err := processQueryResults(rows)
	if err != nil {
		return nil, err
	}

	return fileList, nil
}

// TODO: bug - only files with tags will get returned
func getAllFileInfo() ([]FileInfo, error) {
	// get all file info by joining File and Tag table
	rows, err := database.DB.Query(`
		SELECT f.FileID, f.S3Key, f.Name, f.Size, f.Type, t.Name, t.Type
		FROM File f INNER JOIN Tag t ON f.FileID = t.FileID
		ORDER BY UploadTime
		`)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// process results
	fileList, err := processQueryResults(rows)
	if err != nil {
		return nil, err
	}

	return fileList, nil
}

func getURLsFromS3(info []FileInfo) ([]FileInfo, error) {
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

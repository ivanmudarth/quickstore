package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"time"

	"../database"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jmoiron/sqlx"
)

type FileInfo struct {
	ID         int
	S3Key      string
	FileURL    string
	Name       string
	Size       string
	Type       string
	UploadTime string
	UserTags   []string
	AutoTags   []string
}

type UrlInfo struct {
	ID         int
	ImageURL   string
	Title      string
	URL        string
	UploadTime string
	UserTags   []string
	AutoTags   []string
}

type ItemInfo struct {
	FileInfo   *FileInfo
	UrlInfo    *UrlInfo
	UploadTime time.Time
}

// handles both regular display and search
func DisplayHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Getting all item display info...")

	// Get appropriate item metadata given tag input
	tags := r.URL.Query()["tags[]"]
	var itemInfo []ItemInfo
	var err error
	if len(tags) == 0 {
		itemInfo, err = getAllItemInfo()
	} else {
		itemInfo, err = getSearchItemInfo(tags)
	}
	if err != nil {
		return
	}

	// verify files to be displayed are in S3
	result, err := getURLsFromS3(itemInfo)
	if err != nil {
		http.Error(w, "Error getting files from S3", http.StatusBadRequest)
		return
	}

	// send file info to be displayed back to client
	w.Header().Set("Content-Type", "application/json")
	data, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Error marshaling data", http.StatusBadRequest)
		return
	}
	w.Write(data)
}

func processFileQueryResults(rows *sql.Rows) ([]FileInfo, error) {
	// add query results to list of FileInfo objects
	fileList := []FileInfo{}
	for rows.Next() {
		var fileID int
		var key, fileName, fileSize, fileType, fileUploadTime, tagName, tagType string
		if err := rows.Scan(&fileID, &key, &fileName,
			&fileSize, &fileType, &fileUploadTime, &tagName, &tagType); err != nil {
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
				fileList = append(fileList, FileInfo{fileID, key, "", fileName, fileSize, fileType, fileUploadTime, []string{tagName}, []string{}})
			} else {
				fileList = append(fileList, FileInfo{fileID, key, "", fileName, fileSize, fileType, fileUploadTime, []string{}, []string{tagName}})
			}
		}
	}

	return fileList, nil
}

func processUrlQueryResults(rows *sql.Rows) ([]UrlInfo, error) {
	// add query results to list of UrlInfo objects
	urlList := []UrlInfo{}
	for rows.Next() {
		var urlID int
		var urlImageURL, urlTitle, url, urlUploadTime, tagName, tagType string
		if err := rows.Scan(&urlID, &urlImageURL, &urlTitle,
			&url, &urlUploadTime, &tagName, &tagType); err != nil {
			log.Fatal(err)
			return nil, err
		}

		// if FileInfo entry exists, update its tags. Otherwise create entry
		lastIdx := len(urlList) - 1
		if lastIdx >= 0 && urlList[lastIdx].ID == urlID {
			if tagType == "User" {
				urlList[lastIdx].UserTags = append(urlList[lastIdx].UserTags, tagName)
			} else {
				urlList[lastIdx].AutoTags = append(urlList[lastIdx].AutoTags, tagName)
			}
		} else {
			if tagType == "User" {
				urlList = append(urlList, UrlInfo{urlID, urlImageURL, urlTitle, url, urlUploadTime, []string{tagName}, []string{}})
			} else {
				urlList = append(urlList, UrlInfo{urlID, urlImageURL, urlTitle, url, urlUploadTime, []string{}, []string{tagName}})
			}
		}
	}

	return urlList, nil
}

func joinAndSortItems(fileList []FileInfo, urlList []UrlInfo) []ItemInfo {
	// join FileInfo list and UrlInfo list to become ItemInfo list
	var itemList []ItemInfo

	for _, f := range fileList {
		file := f
		time, err := time.Parse("2006-01-02 15:04:05 -0700 MST", f.UploadTime+" +0000 UTC")
		if err != nil {
			log.Fatal(err)
		}

		item := ItemInfo{
			FileInfo:   &file,
			UrlInfo:    nil,
			UploadTime: time,
		}
		itemList = append(itemList, item)
	}

	for _, u := range urlList {
		url := u
		time, err := time.Parse("2006-01-02 15:04:05 -0700 MST", u.UploadTime+" +0000 UTC")
		if err != nil {
			log.Fatal(err)
		}

		item := ItemInfo{
			FileInfo:   nil,
			UrlInfo:    &url,
			UploadTime: time,
		}
		itemList = append(itemList, item)
	}

	// sort itemList by uploadTime
	sort.Slice(itemList, func(i, j int) bool {
		return itemList[i].UploadTime.Before(itemList[j].UploadTime)
	})

	return itemList
}

// TODO: bug - only items with tags will get returned (add row in tag table with empty str if list of tags is len 0)
func getAllItemInfo() ([]ItemInfo, error) {
	// get all file info by joining File and Tag table
	fileRows, err := database.DB.Query(`
		SELECT f.FileID, f.S3Key, f.Name, f.Size, f.Type, f.UploadTime, t.Name, t.Type
		FROM File f INNER JOIN Tag t ON f.FileID = t.FileID
		`)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// get all url info by joining Url and Tag table
	urlRows, err := database.DB.Query(`
		SELECT u.UrlID, u.ImageURL, u.Title, u.URL, u.UploadTime, t.Name, t.Type
		FROM Url u INNER JOIN Tag t ON u.UrlID = t.UrlID
		`)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// process results
	fileList, err := processFileQueryResults(fileRows)
	if err != nil {
		return nil, err
	}
	urlList, err := processUrlQueryResults(urlRows)
	if err != nil {
		return nil, err
	}

	// join results together and sort by upload time
	itemList := joinAndSortItems(fileList, urlList)

	return itemList, nil
}

// TODO: CURR PROBLEMS:
// uploading subtenant pdf gives error on autotagging (textrank)

func getSearchItemInfo(tags []string) ([]ItemInfo, error) {
	// make query strings
	qFile, args, err := sqlx.In(`
		SELECT s.FileID, s.S3Key, s.Name, s.Size, s.Type, s.UploadTime, t2.Name, t2.Type
		FROM (SELECT DISTINCT f.FileID, f.S3Key, f.Name, f.Size, f.Type, f.UploadTime
			FROM File f INNER JOIN Tag t ON f.FileID = t.FileID
			WHERE LOWER(t.Name) IN (?)) s
		INNER JOIN Tag t2 ON s.FileID = t2.FileID
		`, tags)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	uFile, args, err := sqlx.In(`
		SELECT s.UrlID, s.ImageURL, s.Title, s.URL, s.UploadTime, t2.Name, t2.Type
		FROM (SELECT DISTINCT u.UrlID, u.ImageURL, u.Title, u.URL, u.UploadTime
			FROM Url u INNER JOIN Tag t ON u.UrlID = t.UrlID
			WHERE LOWER(t.Name) IN (?)) s
		INNER JOIN Tag t2 ON s.UrlID = t2.UrlID
		`, tags)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// make actual queries
	fileRows, err := database.DB.Query(qFile, args...)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	urlRows, err := database.DB.Query(uFile, args...)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// proess results
	fileList, err := processFileQueryResults(fileRows)
	if err != nil {
		return nil, err
	}
	urlList, err := processUrlQueryResults(urlRows)
	if err != nil {
		return nil, err
	}

	// join results together and sort by upload time
	itemList := joinAndSortItems(fileList, urlList)

	return itemList, nil
}

func getURLsFromS3(info []ItemInfo) ([]ItemInfo, error) {
	// Check that all keys exist in S3 bucket
	for idx, i := range info {
		if i.FileInfo == nil {
			continue
		}

		s3Key := i.FileInfo.S3Key
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
		info[idx].FileInfo.FileURL = url
	}

	fmt.Println("Got all file URLs\n")
	return info, nil
}

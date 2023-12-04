package handlers

import (
	"fmt"
	"log"
	"mime/multipart"
	"net/http"

	"../database"
	"../tags"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
)

const oneMB = 1000000.0

func createS3Key() string {
	return uuid.New().String()
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: add check to ensure only valid file type is received
	// edgecase: s3 upload fails but metadata does not (and vice versa)
	fmt.Println("Uploading a file...")

	file, header, err := r.FormFile("file")

	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Enforce file size limit
	if header.Size > int64(10*oneMB) {
		http.Error(w, "Size limit of 10 MB reached", http.StatusBadRequest)
		return
	}

	// Generate key to represent file in S3
	key := createS3Key()

	// Upload file to S3
	err = uploadToS3(header, key)
	if err != nil {
		http.Error(w, "Error uploading to S3", http.StatusBadRequest)
		return
	}

	// Upload file's metadata and user tags to MySQL
	tags := r.PostForm["tags[]"]
	fileEntryID, err := uploadMetaData(header, key, tags)
	if err != nil {
		http.Error(w, "Error uploading file metadata", http.StatusBadRequest)
		return
	}

	// Upload file's auto tags to MySQL
	err = uploadAutoTags(header, fileEntryID)
	if err != nil {
		http.Error(w, "Error uploading file autotags", http.StatusBadRequest)
		return
	}

	w.Write([]byte("File uploaded successfully"))
}

func uploadToS3(fileHeader *multipart.FileHeader, key string) (err error) {
	// Open the file from HTTP request
	reqFile, err := fileHeader.Open()
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer reqFile.Close()

	// Upload file
	result, err := AWSConfig.uploader.Upload(&s3manager.UploadInput{
		Bucket: AWSConfig.bucketName,
		Key:    aws.String(key),
		Body:   reqFile,
	})
	if err != nil {
		log.Fatal(err)
		return err
	} else {
		log.Println(result)
	}

	fmt.Printf("File '%s' uploaded successully\n", fileHeader.Filename)
	return nil
}

func uploadMetaData(fileHeader *multipart.FileHeader, s3Key string, userTags []string) (int64, error) {
	fileName := fileHeader.Filename
	fileSize := float64(fileHeader.Size) / oneMB

	// create new entry in File table
	res, err := database.DB.Exec(`
		INSERT INTO File (S3Key, Name, Size) 
		VALUES (?, ?, ?);
		`, s3Key, fileName, fileSize)
	if err != nil {
		log.Fatal(err)
		return -1, err
	}

	// create new entry in Tag table for each user tag
	id, _ := res.LastInsertId()
	for _, t := range userTags {
		_, err := database.DB.Exec(`
			INSERT INTO Tag (FileID, Name, Type) 
			VALUES (?, ?, ?);
			`, id, t, "User")
		if err != nil {
			log.Fatal(err)
			return -1, err
		}
	}

	fmt.Println("File metadata uploaded successfully\n")
	return id, nil
}

func uploadAutoTags(fileHeader *multipart.FileHeader, fileEntryID int64) (err error) {
	// get auto tags from Imagga
	autoTags, err := tags.AutoTagImage(fileHeader)
	if err != nil {
		return err
	}

	// create new entry in Tag table for each auto tag
	for _, t := range autoTags {
		_, err := database.DB.Exec(`
			INSERT INTO Tag (FileID, Name, Type) 
			VALUES (?, ?, ?);
			`, fileEntryID, t, "Auto")
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	fmt.Println("File's auto tags uploaded successfully\n")
	return nil
}

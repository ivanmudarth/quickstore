package handlers

import (
	"fmt"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: add check to ensure only valid file type is received
	fmt.Println("Uploading a file...")

	file, header, err := r.FormFile("file")

	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Enforce file size limit
	one_MB := 1000000
	if header.Size > int64(10*one_MB) {
		http.Error(w, "Size limit of 10 MB reached", http.StatusBadRequest)
		return
	}

	// Upload file to S3
	err = uploadToS3(header)
	if err != nil {
		http.Error(w, "Error uploading to S3", http.StatusBadRequest)
		return
	}

	// Upload file's metadata to Postgres
	err = uploadMetaData()
	if err != nil {
		http.Error(w, "Error uploading file metadata", http.StatusBadRequest)
	}
	w.Write([]byte("File uploaded successfully"))
}

// TODO: name object key better (consider same file name diff contents, use UUID)
func uploadToS3(fileHeader *multipart.FileHeader) (err error) {
	// Open the file from HTTP request
	reqFile, err := fileHeader.Open()
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer reqFile.Close()

	// Upload file
	key := fileHeader.Filename
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

	fmt.Printf("File '%s' uploaded successully\n\n", fileHeader.Filename)
	return nil
}

func uploadMetaData() (err error) {
	// TODO:

	return nil
}

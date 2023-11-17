package handlers

import (
	"fmt"
	"net/http"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Uploading a file...")

	_, handler, _ := r.FormFile("file")
	fileName := handler.Filename
	fmt.Fprintf(w, "filename: "+fileName)

	// http.Error(w, "Unable to parse form", http.StatusBadRequest)

	// error if size too large, type not supported, unable to parse
	// store metadata in postgres
	// upload to S3
	// error if either metadata/file storage fails
}

func UploadMetaData() {
	// TODO
}

func UploadToS3() {
	// TODO
}

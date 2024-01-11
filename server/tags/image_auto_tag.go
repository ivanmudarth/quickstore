package tags

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

type imaggaResponse struct {
	Result struct {
		Tags []struct {
			Tag struct {
				En string
			}
		}
	}
}

// get upload ID which represents file uploaded to Imagga
func getImaggaUploadId(fileHeader *multipart.FileHeader) (string, error) {
	// get raw file from upload request header
	file, err := fileHeader.Open()
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	// read file contents but in bytes
	fileContents, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	file.Close()

	// create payload for new request
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("image", fileHeader.Filename)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	part.Write(fileContents)

	// request upload id represented file in payload
	client := &http.Client{}
	req, _ := http.NewRequest("POST", "https://api.imagga.com/v2/uploads", body)
	req.SetBasicAuth(os.Getenv("IMAGGA_API_KEY"), os.Getenv("IMAGGA_API_SECRET"))
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	defer resp.Body.Close()

	// process response
	resp_body, _ := io.ReadAll(resp.Body)

	m := make(map[string]map[string]string)
	_ = json.Unmarshal(resp_body, &m)
	upload_id := m["result"]["upload_id"]
	fmt.Println("Created Imagga upload ID successfully")

	return upload_id, nil
}

func AutoTagImage(fileHeader *multipart.FileHeader) ([]string, error) {
	// get Imagga ID
	uploadId, err := getImaggaUploadId(fileHeader)
	if err != nil {
		return nil, err
	}

	// request tags for recently uploaded image
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.imagga.com/v2/tags?image_upload_id="+uploadId+"&limit=5", nil)
	req.SetBasicAuth(os.Getenv("IMAGGA_API_KEY"), os.Getenv("IMAGGA_API_SECRET"))

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// process response
	defer resp.Body.Close()

	var output imaggaResponse
	resp_body, err := io.ReadAll(resp.Body)
	json.Unmarshal(resp_body, &output)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	var res []string
	for _, tagInfo := range output.Result.Tags {
		tag := tagInfo.Tag.En
		res = append(res, tag)
	}

	return res, err
}

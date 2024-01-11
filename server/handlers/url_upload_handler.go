package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"../database"
	"../tags"
)

// :
// change get all files query, change search files query,
// make sure response handles new objects properly
// rename FILE to ITEM where appropriate
// use website ss instead of image
func UrlUploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	url := r.FormValue("url")

	// fetch website's title and image
	title, imageUrl, err := getUrlOpenGraph(url)
	if err != nil {
		http.Error(w, "Error uploading to S3", http.StatusBadRequest)
		return
	}

	// upload website's metadata and user tags to mysql
	tags := r.PostForm["tags[]"]
	urlEntryID, err := uploadUrlMetadata(title, imageUrl, url, tags)
	if err != nil {
		http.Error(w, "Error uploading url metadata", http.StatusBadRequest)
		return
	}

	// upload url autotags to mysql
	err = uploadUrlAutoTags(url, urlEntryID)
	if err != nil {
		http.Error(w, "Error uploading url autotags", http.StatusBadRequest)
		return
	}

	w.Write([]byte("Url uploaded successfully"))
}

func getUrlOpenGraph(websiteURL string) (string, string, error) {
	// make a request to opengraph.io api to get websites image and title
	encodedURL := url.PathEscape(websiteURL)
	requestURL := "https://opengraph.io/api/1.1/site/" + encodedURL + "?app_id=" + os.Getenv("OPENGRAPHIO_API_KEY")
	client := &http.Client{}
	req, _ := http.NewRequest("GET", requestURL, nil)

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
		return "", "", err
	}

	// process response
	defer resp.Body.Close()

	var openGraphResponse map[string]map[string]interface{}
	resp_body, err := io.ReadAll(resp.Body)
	json.Unmarshal(resp_body, &openGraphResponse)
	if err != nil {
		log.Fatal(err)
		return "", "", err
	}

	title := openGraphResponse["hybridGraph"]["title"].(string)
	imageURL := openGraphResponse["hybridGraph"]["image"].(string)

	return title, imageURL, nil
}

func uploadUrlMetadata(title string, imageUrl string, url string, userTags []string) (int64, error) {
	// create new entry in Url table
	res, err := database.DB.Exec(`
		INSERT INTO Url (ImageURL, Title, URL, Type)
		VALUES (?, ?, ?, ?);
		`, imageUrl, title, url, "Url")
	if err != nil {
		log.Fatal(err)
		return -1, err
	}

	// create new entry in Tag table for each user tag
	id, _ := res.LastInsertId()
	for _, t := range userTags {
		_, err := database.DB.Exec(`
			INSERT INTO Tag (UrlID, Name, Type) 
			VALUES (?, ?, ?);
			`, id, t, "User")
		if err != nil {
			log.Fatal(err)
			return -1, err
		}
	}

	fmt.Println("Url metadata uploaded successfully\n")
	return id, nil
}

func uploadUrlAutoTags(url string, urlEntryID int64) (err error) {
	// auto tag websites content
	autoTags, err := tags.AutoTagUrl(url)

	// create new entry in Tag table for each auto tag
	for _, t := range autoTags {
		_, err := database.DB.Exec(`
			INSERT INTO Tag (UrlID, Name, Type) 
			VALUES (?, ?, ?);
			`, urlEntryID, t, "Auto")
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	fmt.Println("Url's auto tags uploaded successfully\n")
	return nil
}

package tags

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func scrapeTextFromSite(websiteUrl string) (string, error) {
	// make a request to opengraph.io api to get websites image and title
	encodedURL := url.PathEscape(websiteUrl)
	requestURL := "https://opengraph.io/api/1.1/extract/" + encodedURL + "?app_id=" + os.Getenv("OPENGRAPHIO_API_KEY")
	client := &http.Client{}
	req, _ := http.NewRequest("GET", requestURL, nil)

	resp, err := client.Do(req)

	if err != nil {
		log.Fatal(err)
		return "", err
	}

	// process response
	defer resp.Body.Close()

	var openGraphResponse map[string]interface{}
	resp_body, err := io.ReadAll(resp.Body)
	json.Unmarshal(resp_body, &openGraphResponse)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	siteText := openGraphResponse["concatenatedText"].(string)

	return siteText, nil
}

func AutoTagUrl(url string) ([]string, error) {
	urlText, err := scrapeTextFromSite(url)
	if err != nil {
		return nil, err
	}

	res := GetTopPhrasesFromText(urlText)

	return res, nil
}

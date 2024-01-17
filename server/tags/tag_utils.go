package tags

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func extractKeywordsFromOutput(output string) []string {
	// Extract the keyword with ** around them
	pattern := "\\*\\*(.*?)\\*\\*"
	re := regexp.MustCompile(pattern)

	matches := re.FindAllStringSubmatch(output, -1)

	var keywords []string
	for _, match := range matches {
		if len(match) < 50 {
			keywords = append(keywords, strings.ToLower(match[1]))
		}
	}

	return keywords
}

func GetKeywordsFromText(text string) ([]string, error) {
	// generate list of key phrases using Cohere's Generate API

	url := "https://api.cohere.ai/v1/generate"

	prompt := "Generate only 5 keywords (where each keyword is surrounded by **) that summarize the following text : " + text
	payload := strings.NewReader("{\"truncate\":\"END\",\"return_likelihoods\":\"NONE\",\"prompt\":\"" + prompt + "\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", "Bearer "+os.Getenv("COHERE_API_KEY")) // TODO: add to env file

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return []string{}, err
	}

	// process response
	defer res.Body.Close()

	var cohereResponse map[string][]map[string]interface{}
	resp_body, err := io.ReadAll(res.Body)
	json.Unmarshal(resp_body, &cohereResponse)
	if err != nil {
		log.Fatal(err)
		return []string{}, err
	}
	fmt.Println(cohereResponse)
	output := cohereResponse["generations"][0]["text"].(string)
	keywords := extractKeywordsFromOutput(output)
	fmt.Println(keywords)

	return keywords, nil
}

package tags

import (
	"bytes"
	"log"
	"mime/multipart"

	textrank "github.com/DavidBelicza/TextRank"
	"github.com/ledongthuc/pdf"
)

func extractTextFromPdf(fileHeader *multipart.FileHeader) (string, error) {
	// get raw file from upload request header
	f, err := fileHeader.Open()
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	r, err := pdf.NewReader(f, fileHeader.Size)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}
	buf.ReadFrom(b)

	return buf.String(), nil
}

func AutoTagPdf(fileHeader *multipart.FileHeader) ([]string, error) {
	pdfText, err := extractTextFromPdf(fileHeader)
	if err != nil {
		return nil, err
	}

	// generate list of key phrases using TextRank algorithm
	tr := textrank.NewTextRank()
	rule := textrank.NewDefaultRule()
	language := textrank.NewDefaultLanguage()
	algorithmDef := textrank.NewDefaultAlgorithm()

	tr.Populate(pdfText, language, rule)
	tr.Ranking(algorithmDef)

	// Get top n phrases with weight over 0.5
	rankedPhrases := textrank.FindPhrases(tr)
	res := []string{}
	for i := 0; i <= 5; i++ {
		rp := rankedPhrases[i]
		if rp.Weight >= 0.5 {
			phrase := rp.Left + " " + rp.Right
			res = append(res, phrase)
		}
	}

	return res, nil
}

package tags

import (
	"bytes"
	"log"
	"mime/multipart"

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

	res := GetTopPhrasesFromText(pdfText)

	return res, nil
}
